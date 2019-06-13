package CreateDNSCloudflare

import (
	"github.com/spf13/viper"
	"log"
)









func CreateDNS(dnstype, name,  ipaddres, zones string, ttl int) {



	viper.AddConfigPath("/etc/cloudflare/")
	viper.SetConfigName("cloudcreds")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Looks like there was a problem reading /etc/cloudflare/clouldcreds.yaml. Please make sure it exists.")
	}

	AuthEmail := viper.GetString("cloudflare.authemail")
	AuthToken := viper.GetString("cloudflare.authtoken")

	zone1 := viper.GetString("cloudflare.zone1")
	zone2 := viper.GetString("cloudflare.zone2")
	zone3 := viper.GetString("cloudflare.zone3")
	zone4 := viper.GetString("cloudflare.zone3")
	log.Println("The zone is")
	log.Println(zones)
	log.Println(name)
	log.Println(ipaddres)
	log.Println("Data is ")

	switch zones {

	case "zone1":
		request := &DNSCreate{
			Types: dnstype,
			Name: name,
			Content: ipaddres,
			Ttl: ttl,
			Proxied: false,
			Priority: 10,


		}



		CreateDNSRequest(zone1, AuthEmail, AuthToken, request)







	case "zone2":
		request := DNSCreate{
			dnstype,
			name,
			ipaddres,
			ttl,
			false,
			10,

		}

		CreateDNSRequest(zone2, AuthEmail, AuthToken, &request)


	case "zone3":
		request := DNSCreate{
			dnstype,
			name,
			ipaddres,
			ttl,
			false,
			10,

		}
		CreateDNSRequest(zone3, AuthEmail, AuthToken, &request)



	case "zone4":
		request := DNSCreate{
			dnstype,
			name,
			ipaddres,
			ttl,
			false,
			10,

		}
		CreateDNSRequest(zone4, AuthEmail, AuthToken, &request)


}
	}