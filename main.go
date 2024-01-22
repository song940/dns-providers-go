package main

import (
	"log"

	"github.com/song940/dns-providers-go/providers"
	"github.com/song940/dns-providers-go/types"
)

func main() {
	config := types.Config{
		"email":  "",
		"apiKey": "",
	}
	dns, err := providers.NewDNSProviderByName("cloudflare", config)
	if err != nil {
		panic(err)
	}
	zones, err := dns.ListZones()
	if err != nil {
		panic(err)
	}
	for _, zone := range zones {
		log.Println(zone.ID, zone.Name)
	}

	records, err := dns.ListRecords("a23330422d79eedb6821579eb313cf68")
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		log.Println(record.ID, record.Name, record.Type, record.Value)
	}
}
