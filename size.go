package ray

// Size type.
type Size string

const (
	// Small size.
	Small Size = "sm"

	// Large size.
	Large Size = "lg"
)

// Small function.
func (r *Ray) Small() *Ray {
	return r.Size(Small)
}

// Large function.
func (r *Ray) Large() *Ray {
	return r.Size(Large)
}
