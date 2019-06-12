package CloudFlags

import (
	CreateDNSCloudflare "github.com/DMEvanCT/cloudflare/CreateDNS"
	"github.com/spf13/cobra"
	"log"
)

var dnstype string
var name string
var ip string
var ttl string
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
		CreateDNSCloudflare.CreateDNS(dnstype, name, ip, ttl, zone)

	},
}







func CloudFlareFlags() {
	AddDNS.PersistentFlags().StringVar(&dnstype, "t", "", "DNS type Ex: A record")
	AddDNS.PersistentFlags().StringVar(&name, "n", "", "DNS type Ex: dark.darkmatterit.io")
	AddDNS.PersistentFlags().StringVar(&ip, "i", "", "IP address type Ex: 8.8.8.8")
	AddDNS.PersistentFlags().StringVar(&ttl, "s", "", "ttl between 120 and 2,147,483,647")
	AddDNS.PersistentFlags().StringVar(&zone, "z", "zone1", "zone1 zone2 zone3 zone4")


	BaseCMD.AddCommand(AddDNS)


	BaseCMD.Execute()




}

