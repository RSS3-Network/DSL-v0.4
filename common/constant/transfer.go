//go:generate stringer -type=TransferTagPriority -linecomment
package constant

// tag priority
type TransferTagPriority int

// int -> *string
func (i TransferTagPriority) AsInt() *int {
	str := int(i)
	return &str
}

const (
	TransferTagPriorityToken TransferTagPriority = 1 // token
	TransferTagSwap          TransferTagPriority = 2 // swap
)
