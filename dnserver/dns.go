package main

import (
	"errors"
	"net"

	"golang.org/x/net/dns/dnsmessage"
)

type DNSServer interface {
	Listen()
	Query(Packet)
}

type DNSService struct {
	conn       *net.UDPConn
	book       store
	memo       addrBag
	forwarders []net.UDPAddr
}

type Packet struct {
	addr    net.UDPAddr
	message dnsmessage.Message
}

const (
	udpPort   int = 8080
	packetLen int = 512
)

var (
	errTypeNotSupport = errors.New("Type not supported")
	errIPInvalid      = errors.New("Invalid IP address")
)
