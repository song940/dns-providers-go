package dnspod

import (
	"strconv"

	"github.com/nrdcg/dnspod-go"
	"github.com/song940/dns-providers-go/types"
)

type DNSPodDNSProvider struct {
	client *dnspod.Client
}

// https://github.com/nrdcg/dnspod-go
func NewDNSProvider(config types.Config) (provider *DNSPodDNSProvider, err error) {
	provider = &DNSPodDNSProvider{}
	token := config["token"]
	params := dnspod.CommonParams{
		LoginToken: token,
		Format:     "json",
	}
	provider.client = dnspod.NewClient(params)
	return
}

// AddRecord implements types.IDNSProvider.
func (c *DNSPodDNSProvider) AddRecord(zoneID string, record *types.Record) error {
	id, err := strconv.Atoi(zoneID)
	if err != nil {
		return err
	}
	domain, _, err := c.client.Domains.Get(id)
	if err != nil {
		return err
	}
	_, _, err = c.client.Records.Create(domain.Name, dnspod.Record{
		Name:  record.Name,
		Type:  record.Type,
		Value: record.Value,
	})
	return err
}

// DeleteRecord implements types.IDNSProvider.
func (c *DNSPodDNSProvider) DeleteRecord(zoneID string, recordID string) error {
	id, err := strconv.Atoi(zoneID)
	if err != nil {
		return err
	}
	domain, _, err := c.client.Domains.Get(id)
	if err != nil {
		return err
	}
	_, err = c.client.Records.Delete(domain.Name, recordID)
	return err
}

// ListRecords implements types.IDNSProvider.
func (c *DNSPodDNSProvider) ListRecords(zoneID string) (records []types.Record, err error) {
	arr, _, err := c.client.Records.List(zoneID, "")
	for _, record := range arr {
		records = append(records, types.Record{
			ID:      record.ID,
			Name:    record.Name,
			Type:    record.Type,
			Value:   record.Value,
			Comment: record.Remark,
			// TTL:     record.TTL,
		})
	}
	return
}

// ListZones implements types.IDNSProvider.
func (c *DNSPodDNSProvider) ListZones() (zones []types.Zone, err error) {
	domains, _, _ := c.client.Domains.List()
	for _, domain := range domains {
		zones = append(zones, types.Zone{
			ID:   domain.ID.String(),
			Name: domain.Name,
		})
	}
	return
}

// UpdateRecord implements types.IDNSProvider.
func (c *DNSPodDNSProvider) UpdateRecord(zoneID string, record *types.Record) error {
	id, err := strconv.Atoi(zoneID)
	if err != nil {
		return err
	}
	domain, _, err := c.client.Domains.Get(id)
	if err != nil {
		return err
	}
	_, _, err = c.client.Records.Update(domain.Name, record.ID, dnspod.Record{
		Name:   record.Name,
		Type:   record.Type,
		Value:  record.Value,
		Remark: record.Comment,
	})
	return err
}
