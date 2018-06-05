package states

import (
	"github.com/zclconf/go-cty/cty"
)

// ResourceInstanceObject is the local representation of a specific remote
// object associated with a resource instance. In practice not all remote
// objects are actually remote in the sense of being accessed over the network,
// but this is the most common case.
type ResourceInstanceObject struct {
	// Value is the value (of the object type implied by the associated resource
	// type schema) that represents this remote object in Terraform Language
	// expressions and is compared with configuration when producing a diff.
	Value cty.Value

	// Internal is an opaque value set by the provider when this object was
	// last created or updated. Terraform Core does not use this value in
	// any way and it is not exposed anywhere in the user interface, so
	// a provider can use it for retaining any necessary private state.
	Private cty.Value

	// Status represents the "readiness" of the object as of the last time
	// it was updated.
	Status ObjectStatus
}

// ObjectStatus represents the status of a RemoteObject.
type ObjectStatus rune

//go:generate stringer -type ObjectStatus

const (
	// ObjectReady is an object status for an object that is ready to use.
	ObjectReady ObjectStatus = 'R'

	// ObjectTainted is an object status representing an object that is in
	// an unrecoverable bad state due to a partial failure during a create,
	// update, or delete operation. Since it cannot be moved into the
	// ObjectRead state, a tainted object must be replaced.
	ObjectTainted ObjectStatus = 'T'
)
