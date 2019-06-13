package CreateDNSCloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type DNSCreate struct {
	Types string `json:"type"  bson:"type"`
	Name string `json:"name"  bson:"name"`
	Content string `json:"content"  bson:"content"`
	Ttl int `json:"ttl"  bson:"ttl"`
	Proxied bool `json:"proxied"  bson:"proxied"`
	Priority int `json:"priority"  bson:"priority"`

}

func CreateDNSRequest(zone, AuthEmail, AuthToken string, dns *DNSCreate) {
	data, _ := json.Marshal(&dns)
	//data, _ := json.Marshal(DNSCreate{Type: dns.Type, name: dns.name, content: dns.content, ttl: dns.ttl, proxied: dns.proxied, priority: dns.priority})
	fmt.Println("The Data!!!!")
	fmt.Println(string(data))
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


	// sets up the client and executes post
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		log.Println("DNS name " + dns.Name + " DNS type " )
	} else {
		log.Println(resp.Status)
		log.Println("There seems to be an issue with the request")


		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)


	}
}
