package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/song940/dns-providers-go/types"
)

type CloudflareDNSProvider struct {
	client *cloudflare.API
}

func NewDNSProvider(config types.Config) (provider *CloudflareDNSProvider, err error) {
	provider = &CloudflareDNSProvider{}
	key := config["apiKey"]
	email := config["email"]
	token := config["token"]
	if token != "" {
		provider.client, err = cloudflare.NewWithAPIToken(token)
	} else {
		provider.client, err = cloudflare.New(key, email)
	}
	return
}

// ListZones implements types.IDNSProvider.
func (c *CloudflareDNSProvider) ListZones() (zones []types.Zone, err error) {
	ctx := context.Background()
	arr, err := c.client.ListZones(ctx)
	for _, z := range arr {
		zones = append(zones, types.Zone{
			ID:   z.ID,
			Name: z.Name,
		})
	}
	return
}

// ListRecords implements types.IDNSProvider.
func (c *CloudflareDNSProvider) ListRecords(zoneID string) (records []types.Record, err error) {
	ctx := context.Background()
	zone := cloudflare.ZoneIdentifier(zoneID)
	arr, _, err := c.client.ListDNSRecords(ctx, zone, cloudflare.ListDNSRecordsParams{})
	for _, record := range arr {
		records = append(records, types.Record{
			ID:    record.ID,
			Name:  record.Name,
			Type:  record.Type,
			Value: record.Content,
		})
	}
	return
}

// AddRecord implements types.IDNSProvider.
func (c *CloudflareDNSProvider) AddRecord(zoneID string, record *types.Record) error {
	ctx := context.Background()
	zone := cloudflare.ZoneIdentifier(zoneID)
	_, err := c.client.CreateDNSRecord(ctx, zone, cloudflare.CreateDNSRecordParams{
		Type:    record.Type,
		Name:    record.Name,
		Content: record.Value,
		TTL:     record.TTL,
		Comment: record.Comment,
	})
	return err
}

// DeleteRecord implements types.IDNSProvider.
func (c *CloudflareDNSProvider) DeleteRecord(zoneID string, recordID string) error {
	ctx := context.Background()
	zone := cloudflare.ZoneIdentifier(zoneID)
	err := c.client.DeleteDNSRecord(ctx, zone, recordID)
	return err
}

// UpdateRecord implements types.IDNSProvider.
func (c *CloudflareDNSProvider) UpdateRecord(zoneID string, record *types.Record) error {
	ctx := context.Background()
	zone := cloudflare.ZoneIdentifier(zoneID)
	_, err := c.client.UpdateDNSRecord(ctx, zone, cloudflare.UpdateDNSRecordParams{
		ID:      record.ID,
		Type:    record.Type,
		Name:    record.Name,
		Content: record.Value,
		TTL:     record.TTL,
		Comment: &record.Comment,
	})
	return err
}
