# CloudFlareCLI

####Configuration /etc/cloudflare/cloudcreds.yaml
The current configuration supports up to 4 zones
```yaml
cloudflare:
  authemail: (cloudflare email)
  authtoken: (cloudflare api key)
  # example.com
  zone1: (zone identifier 1)
  #example2.com
  zone2: (zone identifier 2)
  #example3.com
  zone3: (zone identifier 3)
  #example4.com
  zone4: (zone identifier 4)
```