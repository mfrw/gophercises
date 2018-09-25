package utils

import (
	"crypto/tls"
	"fmt"
	"log"
)

func getCert(certfile, keyfile string) (c tls.Certificate, err error) {
	if certfile != "" && keyfile != "" {
		c, err := tls.LoadX509KeyPair(certfile, keyfile)
		if err != nil {
			log.Printf("Error lodaing key pair: %v\n", err)
		}
	} else {
		err = fmt.Errorf("I have no certificate")
	}
	return
}

func ClientCertReqFunc(certfile, keyfile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c, err := getCert(certfile, keyfile)
	return func(certReq *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		fmt.Println("Recieved certificate request: sending certificate")
		if err != nil {
			fmt.Println("I have no certificate")
		} else {
			err := OutputPEMFile(certificate)
			if err != nil {
				log.Printf("%v\n", err)
			}
		}
		Wait()
		return &c, nil
	}
}
