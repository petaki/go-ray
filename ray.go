package ray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Ray type.
type Ray struct {
	uuid     string
	settings *Settings
	client   *http.Client
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
			Enable:              true,
			Host:                "localhost",
			Port:                23517,
			AlwaysSendRawValues: false,
		}
	}

	r.client = &http.Client{
		Timeout: 2 * time.Second,
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

// Ban function.
func (r *Ray) Ban() *Ray {
	return r.Send("🕶")
}

// Charles function.
func (r *Ray) Charles() *Ray {
	return r.Send("🎶 🎹 🎷 🕺")
}

// Raw function.
func (r *Ray) Raw(arguments ...interface{}) *Ray {
	if len(arguments) == 0 {
		return r
	}

	return r.sendRequest([]*payload{
		newLogPayload(arguments...),
	}, nil)
}

// Send function.
func (r *Ray) Send(arguments ...interface{}) *Ray {
	if len(arguments) == 0 {
		return r
	}

	if r.settings.AlwaysSendRawValues {
		return r.Raw(arguments...)
	}

	return r.sendRequest(
		createPayloadsForValues(arguments...),
		nil,
	)
}

func (r *Ray) sendRequest(payloads []*payload, meta map[string]interface{}) *Ray {
	if r.Disabled() {
		return r
	}

	data, _ := json.Marshal(map[string]interface{}{
		"uuid":     r.uuid,
		"payloads": payloads,
		"meta":     meta,
	})

	resp, _ := r.client.Post(
		fmt.Sprintf("http://%s:%d/", r.settings.Host, r.settings.Port),
		"application/json",
		bytes.NewBuffer(data),
	)

	defer resp.Body.Close()

	return r
}
