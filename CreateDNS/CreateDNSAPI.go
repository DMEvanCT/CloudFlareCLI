package CreateDNSCloudflare

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

func CreateDNSRequest(zone, AuthEmail, AuthToken string, dns *DNSCreate) {
	data, _ := json.Marshal(dns)
	log.Println(zone)
	dnsurl := "https://api.cloudflare.com/client/v4/zones/" + zone + "/dns_records"
	log.Println(dnsurl)
	log.Println("I got here")
	req, err := http.NewRequest("POST", dnsurl, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("X-Auth-Email", AuthEmail)
	req.Header.Set("X-Auth-Key", AuthToken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// sets up the client and executes post
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println("I got here....")

	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		log.Println("DNS name " + dns.name + " DNS type " + dns.dnstype + " DNS IP " + dns.content )
	} else {
		log.Println(resp.Status)
		log.Println("There seems to be an issue with the request")

	}
}
