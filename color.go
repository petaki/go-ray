package ray

// Color type.
type Color string

const (
	// Green color.
	Green Color = "green"

	// Orange color.
	Orange Color = "orange"

	// Red color.
	Red Color = "red"

	// Purple color.
	Purple Color = "purple"

	// Blue color.
	Blue Color = "blue"

	// Gray color.
	Gray Color = "gray"
)

// Color function.
func (r *Ray) Color(color Color) *Ray {
	return r.sendRequest(newColorPayload(color))
}

// Green function.
func (r *Ray) Green() *Ray {
	return r.sendRequest(newColorPayload(Green))
}

// Orange function.
func (r *Ray) Orange() *Ray {
	return r.sendRequest(newColorPayload(Orange))
}

// Red function.
func (r *Ray) Red() *Ray {
	return r.sendRequest(newColorPayload(Red))
}

// Purple function.
func (r *Ray) Purple() *Ray {
	return r.sendRequest(newColorPayload(Purple))
}

// Blue function.
func (r *Ray) Blue() *Ray {
	return r.sendRequest(newColorPayload(Blue))
}

// Gray function.
func (r *Ray) Gray() *Ray {
	return r.sendRequest(newColorPayload(Gray))
}
