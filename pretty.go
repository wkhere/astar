package astar

import "fmt"

// pretty-print:
func (x *qitem) String() string {
	return fmt.Sprintf("%v", *x)
}
