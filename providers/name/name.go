package name

import (
	"fmt"
	"strconv"

	"github.com/namedotcom/go/namecom"
	"github.com/song940/dns-providers-go/types"
)

type DNSProvider struct {
	client *namecom.NameCom
}

// https://pkg.go.dev/github.com/namedotcom/go
func NewDNSProvider(config types.Config) (provider *DNSProvider, err error) {
	provider = &DNSProvider{}
	user := config["user"]
	token := config["token"]
	provider.client = namecom.New(user, token)
	return
}

// AddRecord implements types.IDNSProvider.
func (c *DNSProvider) AddRecord(zoneID string, record *types.Record) error {
	_, err := c.client.CreateRecord(&namecom.Record{
		DomainName: zoneID,
		Type:       record.Type,
		Host:       record.Name,
		Answer:     record.Value,
	})
	return err
}

// DeleteRecord implements types.IDNSProvider.
func (c *DNSProvider) DeleteRecord(zoneID string, recordID string) error {
	id, err := strconv.Atoi(recordID)
	if err != nil {
		return err
	}
	_, err = c.client.DeleteRecord(&namecom.DeleteRecordRequest{
		DomainName: zoneID,
		ID:         int32(id),
	})
	return err
}

// ListRecords implements types.IDNSProvider.
func (c *DNSProvider) ListRecords(zoneID string) (records []types.Record, err error) {
	res, err := c.client.ListRecords(&namecom.ListRecordsRequest{})
	for _, record := range res.Records {
		records = append(records, types.Record{
			ID:    fmt.Sprintf("%d", record.ID),
			Name:  record.Host,
			Type:  record.Type,
			Value: record.Answer,
			TTL:   int(record.TTL),
		})
	}
	return
}

// ListZones implements types.IDNSProvider.
func (c *DNSProvider) ListZones() (zones []types.Zone, err error) {
	domains, err := c.client.ListDomains(&namecom.ListDomainsRequest{})
	if err != nil {
		return
	}
	for _, domain := range domains.Domains {
		zones = append(zones, types.Zone{
			ID:        domain.DomainName,
			Name:      domain.DomainName,
			CreatedAt: domain.CreateDate,
		})
	}
	return
}

// UpdateRecord implements types.IDNSProvider.
func (c *DNSProvider) UpdateRecord(zoneID string, record *types.Record) error {
	_, err := c.client.UpdateRecord(&namecom.Record{
		DomainName: zoneID,
		Host:       record.Name,
		Type:       record.Type,
		Answer:     record.Value,
		TTL:        uint32(record.TTL),
	})
	return err
}
