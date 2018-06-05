package statefile

import (
	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/states"
	"github.com/zclconf/go-cty/cty"
)

// Raw represents the structure of a state file, ready to be decoded into a
// states.State object.
type Raw struct {

	// TerraformVersion is the version of Terraform that created the state file.
	TerraformVersion string

	// Serial is incremented on any operation that modifies
	// the State file. It is used to detect potentially conflicting
	// updates.
	Serial int64

	// Lineage is set when a new, blank state is created and then
	// never updated. This allows us to determine whether the serials
	// of two states can be meaningfully compared.
	// Apart from the guarantee that collisions between two lineages
	// are unlikely, this value is opaque and external callers should only
	// compare lineage strings byte-for-byte for equality.
	Lineage string

	OutputValues map[string]*states.OutputValue

	ResourceInstanceObjects []*RawResourceInstanceObject
}

// RawResourceInstanceObject is a flattened representation of the heirarchy
// of states.Resource, states.ResourceInstance and states.ResourceInstanceObject,
// giving the state for an instance object and enough information to construct
// the heirarchical structure around it.
//
// The value has not yet been decoded because this requires access to the
// provider that manages the object in order to run any required upgrade steps.
type RawResourceInstanceObject struct {
	Addr       addrs.AbsResourceInstance
	Generation states.Generation
	Status     states.ObjectStatus

	ProviderConfig addrs.AbsProviderConfig

	Dependencies []addrs.Referenceable

	Value RawValue

	Private cty.Value
}

// RawValue represents a value that may need further processing (e.g. schema
// upgrade or flatmap inflating) before it can be returned.
type RawValue interface {
}
