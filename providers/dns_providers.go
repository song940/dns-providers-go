package providers

import (
	"fmt"

	"github.com/song940/dns-providers-go/providers/cloudflare"
	"github.com/song940/dns-providers-go/providers/dnspod"
	"github.com/song940/dns-providers-go/providers/name"
	"github.com/song940/dns-providers-go/types"
)

func NewDNSProviderByName(key string, config types.Config) (types.IDNSProvider, error) {
	switch key {
	case "cloudflare":
		return cloudflare.NewDNSProvider(config)
	case "dnspod":
		return dnspod.NewDNSProvider(config)
	case "name":
		return name.NewDNSProvider(config)
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", key)
	}
}
