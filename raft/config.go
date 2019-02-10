package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	multierror "github.com/hashicorp/go-multierror"
	template "github.com/hashicorp/go-sockaddr/template"
	flag "github.com/ogier/pflag"
)

type RawConfig struct {
	BindAddress string
	JoinAddress string
	DataDir     string
	RaftPort    int
	HTTPPort    int
	Bootstrap   bool
}

type Config struct {
	RaftAddress net.Addr
	HTTPAddress net.Addr
	JoinAddress string
	DataDir     string
	Bootstrap   bool
}

type ConfigError struct {
	ConfigurationPoint string
	Err                error
}

func (err *ConfigError) Error() string {
	return fmt.Sprintf("%s: %s", err.ConfigurationPoint, err.Err.Error())
}

func resolveConfig(rawConfig *RawConfig) (*Config, error) {
	var errors *multierror.Error

	var bindAddr net.IP
	resolvedBindAddr, err := template.Parse(rawConfig.BindAddress)
	if err != nil {
		configErr := &ConfigError{
			ConfigurationPoint: "bind-address",
			Err:                err,
		}
		errors = multierror.Append(errors, configErr)
	} else {
		bindAddr = net.ParseIP(resolvedBindAddr)
		if bindAddr == nil {
			err := fmt.Errorf("cannot parse IP address: %s", resolvedBindAddr)
			configErr := &ConfigError{
				ConfigurationPoint: "bind-address",
				Err:                err,
			}
			errors = multierror.Append(errors, configErr)
		}
	}

	// Raft Port
	if rawConfig.RaftPort < 1 || rawConfig.RaftPort > 65535 {
		configErr := &ConfigError{
			ConfigurationPoint: "bind-address",
			Err:                err,
		}
		errors = multierror.Append(errors, configErr)
	}

	// Raft Address
	raftAddr := &net.TCPAddr{
		IP:   bindAddr,
		Port: rawConfig.RaftPort,
	}

	// HTTP port
	if rawConfig.HTTPPort < 1 || rawConfig.HTTPPort > 65536 {
		configErr := &ConfigError{
			ConfigurationPoint: "http-port",
			Err:                fmt.Errorf("port numbers must be 1 < port < 65536"),
		}
		errors = multierror.Append(errors, configErr)
	}

	// Construct HTTP Address
	httpAddr := &net.TCPAddr{
		IP:   bindAddr,
		Port: rawConfig.HTTPPort,
	}

	// Data Dir
	dataDir, err := filepath.Abs(rawConfig.DataDir)
	if err != nil {
		configErr := &ConfigError{
			ConfigurationPoint: "data-dir",
			Err:                err,
		}
		errors = multierror.Append(errors, configErr)

	}

	if err := errors.ErrorOrNil(); err != nil {
		return nil, err
	}
	return &Config{
		DataDir:     dataDir,
		JoinAddress: rawConfig.JoinAddress,
		RaftAddress: raftAddr,
		HTTPAddress: httpAddr,
		Bootstrap:   rawConfig.Bootstrap,
	}, nil
}

func readRawConfig() *RawConfig {
	var config RawConfig

	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}

	defaultDataPath := filepath.Join(pwd, "raft")
	flag.StringVarP(&config.DataDir, "data-dir", "d", defaultDataPath, "Path in which to store Raft Data")
	flag.StringVarP(&config.BindAddress, "bind-address", "a", "127.0.0.1", "IP address on which to bind")
	flag.IntVarP(&config.RaftPort, "raft-port", "r", 7000, "Port on which to bind Raft")
	flag.IntVarP(&config.HTTPPort, "http-port", "h", 8000, "Port on which to bind HTTP")
	flag.StringVarP(&config.JoinAddress, "join", "j", "", "Address of another node to join")
	flag.BoolVar(&config.Bootstrap, "bootstrap", false, "Bootstrap the cluster with this node")

	flag.Parse()
	return &config
}
