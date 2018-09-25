package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mfrw/gophercises/tls/utils"
)

func main() {
	client := getClient()
	resp, err := client.Get("https://some-service:8080")
	must(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	must(err)

	fmt.Printf("Status: %s Body: %s\n", resp.Status, string(body))
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../ca/minica.pem")
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		RootCAs:              cp,
		GetClientCertificate: utils.ClientCertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertifcate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func must(err error) {
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
}
