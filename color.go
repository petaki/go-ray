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

// Green function.
func (r *Ray) Green() *Ray {
	return r.Color(Green)
}

// Orange function.
func (r *Ray) Orange() *Ray {
	return r.Color(Orange)
}

// Red function.
func (r *Ray) Red() *Ray {
	return r.Color(Red)
}

// Purple function.
func (r *Ray) Purple() *Ray {
	return r.Color(Purple)
}

// Blue function.
func (r *Ray) Blue() *Ray {
	return r.Color(Blue)
}

// Gray function.
func (r *Ray) Gray() *Ray {
	return r.Color(Gray)
}
