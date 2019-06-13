package CloudFlags

import (
	CreateDNSCloudflare "github.com/DMEvanCT/cloudflare/CreateDNS"
	"github.com/spf13/cobra"
	"log"
)

var dnstype string
var name string
var ip string
var ttl int
var zone string



var BaseCMD = &cobra.Command{
	Use: "cloudflair",
	Short: "CloudFlare DNS options",
	Long: "This reaches out to the cloudlare api to create a dns record",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("You must choose an action")

	},
}

var AddDNS = &cobra.Command{
	Use: "add",
	Short: "Create Cloudflare DNS record",
	Long: "This reaches out to the cloudlare api to create a dns record",

	Run: func(cmd *cobra.Command, args []string) {
		CreateDNSCloudflare.CreateDNS(dnstype, name, ip, zone , ttl)

	},
}






func CloudFlareFlags() {
	AddDNS.PersistentFlags().StringVarP(&dnstype, "dnstype",  "t", "", "DNS type EX: A, CNAME, TXT")
	AddDNS.PersistentFlags().StringVarP(&name, "name", "n", "","DNS name Ex: dark.darkmatterit.io")
	AddDNS.PersistentFlags().StringVarP(&ip, "ip", "i", "", "IP address type Ex: 8.8.8.8")
	AddDNS.PersistentFlags().StringVarP(&zone, "zone", "z", "", "zone1 zone2 zone3 zone4")
	AddDNS.PersistentFlags().IntVarP(&ttl, "ttl", "s", 120, "Time to live in seconds above 120")


	BaseCMD.AddCommand(AddDNS)


	BaseCMD.Execute()




}

