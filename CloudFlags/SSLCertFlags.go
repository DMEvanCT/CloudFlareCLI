package CloudFlags

import (
	"github.com/spf13/cobra"
	"log"
)

var BaseSSL = &cobra.Command{
	Use: "ssl",
	Short: "CloudFlare DNS options",
	Long: "This reaches out to the cloudlare api to create a dns record",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("You must choose an action")

	},
}

var buyssl = &cobra.Command{
	Use: "purchase",
	Short: "CloudFlare SSL Purchase",
	Long: "This reaches out to the cloudlare api to buy an ssl cert",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Yep bought this ssl cert here")
	},
}


func SSLCertFlags() {
	buyssl.PersistentFlags().StringP("dnstype", "t", "", "DNS type Ex: A record")
	buyssl.PersistentFlags().StringP("name", "n", "", "DNS type Ex: dark.darkmatterit.io")
	buyssl.PersistentFlags().StringP("ip", "i", "", "IP address type Ex: 8.8.8.8")
	buyssl.AddCommand(AddDNS)
	buyssl.Execute()

}
