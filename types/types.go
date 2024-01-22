package types

type Zone struct {
	ID   string
	Name string
}

type Record struct {
	ID      string
	Type    string
	Name    string
	Value   string
	TTL     int
	Comment string
}

type IDNSProvider interface {
	ListZones() (zones []Zone, err error)
	ListRecords(zoneID string) (records []Record, err error)
	AddRecord(zoneID string, record *Record) error
	UpdateRecord(zoneID string, record *Record) error
	DeleteRecord(zoneID string, recordID string) error
}

type Config = map[string]string
