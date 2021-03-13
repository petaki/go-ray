package ray

// Size type.
type Size string

const (
	// Small size.
	Small Size = "sm"

	// Large size.
	Large Size = "lg"
)

// Size function.
func (r *Ray) Size(size Size) *Ray {
	return r.sendRequest(newSizePayload(size))
}

// Small function.
func (r *Ray) Small() *Ray {
	return r.Size(Small)
}

// Large function.
func (r *Ray) Large() *Ray {
	return r.Size(Large)
}
