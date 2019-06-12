package CreateDNSCloudflare

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type DNSCreate struct {
	dnstype string `json:"type"  bson:"type"`
	name string `json:"name"  bson:"name"`
	content string `json:"content"  bson:"content"`
	ttl int `json:"ttl"  bson:"ttl"`
}




func CreateDNS(dnstype, name,  ipaddres, ttl, zn string) {
	jsondata := make(map[string]string)

	jsondata["type"] = dnstype
	jsondata["name"] = name
	jsondata["content"] = ipaddres
	jsondata["ttl"] = ttl
	jsondata["priority"] = "10"
	jsondata["proxied"] = "false"

	data, _ := json.Marshal(jsondata)

	viper.AddConfigPath("/etc/cloudflare/")
	viper.SetConfigName("cloudcreds")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Looks like there was a problem reading /etc/cloudflare/clouldcreds.yaml. Please make sure it exists.")
	}

	AuthEmail := viper.GetString("cloudflare.authemail")
	Authtoken := viper.GetString("cloudflare.authtoken")

	zone1 := viper.GetString("cloudflare.zone1")
	zone2 := viper.GetString("cloudflare.zone2")
	zone3 := viper.GetString("cloudflare.zone3")
	zone4 := viper.GetString("cloudflare.zone3")
	log.Println("The zone is")
	log.Println(zn)
	log.Println(name)
	log.Println(ipaddres)

	switch zn {

	case "zone1":
		log.Println(zone1)
		dnsurl := "https://api.cloudflare.com/client/v4/zones/" + zone1 + "/dns_records"
		log.Println(dnsurl)
		log.Println("I got here")
		req, err := http.NewRequest("POST", dnsurl, bytes.NewBuffer(data))

		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Println(err)
		}
		req.Header.Set("X-Auth-Email", AuthEmail)
		req.Header.Set("X-Auth-Key", Authtoken)
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
			log.Println("We sent the slack message!")
		} else {
			fmt.Print(&resp.Body)
			fmt.Println(resp.Status)
			fmt.Println("There seems to be an issue. Check the logs.")
		}

	case "zone2":
		dnsurl := "https://api.cloudflare.com/client/v4/zones/" + zone2 + "/dns_records"
		req, err := http.NewRequest("POST", dnsurl, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Println(err)
		}
		req.Header.Set("X-Auth-Email", AuthEmail)
		req.Header.Set("X-Auth-Key", Authtoken)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		// sets up the client and executes post
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		if resp.Status == "200 OK" {
			log.Println("We sent the slack message!")
		} else {
			fmt.Println(resp.Status)
			fmt.Println("There seems to be an issue. Check the logs api pods!")
		}
		log.Println(resp.Body)

	case "zone3":
		dnsurl := "https://api.cloudflare.com/client/v4/zones/" + zone3 + "/dns_records"
		req, err := http.NewRequest("POST", dnsurl, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Println(err)
		}
		req.Header.Set("X-Auth-Email", AuthEmail)
		req.Header.Set("X-Auth-Key", Authtoken)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		// sets up the client and executes post
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		if resp.Status == "200 OK" {
			log.Println("We sent the slack message!")
		} else {
			fmt.Println(resp.Status)
			fmt.Println("There seems to be an issue. Check the logs")
		}
		log.Println(resp.Body)

	case "zone4":
		dnsurl := "https://api.cloudflare.com/client/v4/zones/" + zone4 + "/dns_records"
		req, err := http.NewRequest("POST", dnsurl, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Println(err)
		}
		req.Header.Set("X-Auth-Email", AuthEmail)
		req.Header.Set("X-Auth-Key", Authtoken)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		// sets up the client and executes post
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		if resp.Status == "200 OK" {
			log.Println("We sent the slack message!")
		} else {
			fmt.Println(resp.Status)
			fmt.Println("There seems to be an issue. Check logs!")
		}
		log.Println(resp.Body)

	}
}
