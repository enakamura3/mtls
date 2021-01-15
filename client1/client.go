package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	URL := "https://some.company.wow:8443/hello"

	// Read the key pair to create certificate using X509KeyPair
	cf, err := ioutil.ReadFile("client.crt")
	if err != nil {
		log.Fatal(err)
	}
	kf, err := ioutil.ReadFile("client.key")
	if err != nil {
		log.Fatal(err)
	}
	cert, err := tls.X509KeyPair(cf, kf)
	if err != nil {
		log.Fatal(err)
	}

	// Read the key pair to create certificate using LoadX509KeyPair
	//cert, err := tls.LoadX509KeyPair("client.crt", "client.key")

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("../server1/server.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
				//InsecureSkipVerify: true,
			},
		},
	}

	// Request /hello via the created HTTPS client over port 8443 via GET
	r, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}
