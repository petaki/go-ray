package ray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Ray type.
type Ray struct {
	uuid     string
	settings *Settings
	enabled  bool
}

// New function.
func New(settings *Settings) *Ray {
	r := new(Ray)
	r.uuid = uuid.NewString()

	if settings != nil {
		r.settings = settings
	} else {
		r.settings = &Settings{
			Enable: true,
			Host:   "localhost",
			Port:   23517,
		}
	}

	r.enabled = r.settings.Enable

	return r
}

// Enable function.
func (r *Ray) Enable() *Ray {
	r.enabled = true

	return r
}

// Disable function.
func (r *Ray) Disable() *Ray {
	r.enabled = false

	return r
}

// Enabled function.
func (r *Ray) Enabled() bool {
	return r.enabled
}

// Disabled function.
func (r *Ray) Disabled() bool {
	return !r.enabled
}

// Color function.
func (r *Ray) Color(color Color) *Ray {
	return r.sendRequest(newColorPayload(color))
}

// Ban function.
func (r *Ray) Ban() *Ray {
	return r.Send("🕶")
}

// Charles function.
func (r *Ray) Charles() *Ray {
	return r.Send("🎶 🎹 🎷 🕺")
}

// Send function.
func (r *Ray) Send(arguments ...interface{}) *Ray {
	return r.sendRequest(newLogPayload(arguments))
}

func (r *Ray) sendRequest(payloads ...*payload) *Ray {
	if r.Disabled() {
		return r
	}

	data, _ := json.Marshal(map[string]interface{}{
		"uuid":     r.uuid,
		"payloads": payloads,
		"meta":     []string{},
	})

	resp, _ := http.Post(
		fmt.Sprintf("http://%s:%d/", r.settings.Host, r.settings.Port),
		"application/json",
		bytes.NewBuffer(data),
	)

	defer resp.Body.Close()

	return r
}
