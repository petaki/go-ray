package ray

// Ray type.
type Ray struct {
	settings *Settings
}

// New function.
func New(settings *Settings) *Ray {
	r := new(Ray)

	if settings != nil {
		r.settings = settings
	} else {
		r.settings = &Settings{
			Host: "localhost",
			Port: 23517,
		}
	}

	return r
}
