package ray

import "github.com/google/uuid"

// Ray type.
type Ray struct {
	uuid     string
	settings *Settings
}

// New function.
func New(settings *Settings) *Ray {
	r := new(Ray)
	r.uuid = uuid.NewString()

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
