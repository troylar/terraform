package statefile

import (
	"github.com/hashicorp/terraform/config/configschema"
)

type ProviderLoader interface {
	ResourceTypeConfig(providerType string, resourceType string) *configschema.Block
	DataSourceConfig(providerType string, dataSource string) *configschema.Block
}
