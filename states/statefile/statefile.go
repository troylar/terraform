package statefile

import (
	"io"

	"github.com/hashicorp/terraform/states"
)

// WriteState writes the given state to the given writer in the current state
// serialization format.
func WriteState(s *states.State, w io.Writer) error {

}

// ReadState reads a state from the given reader, in the current state
// serialization format.
func ReadState(r io.Reader) (*states.State, error) {

}
