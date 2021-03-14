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

// DefaultRay variable.
var DefaultRay = New(nil)

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

// NewScreen function.
func (r *Ray) NewScreen(name string) *Ray {
	return r.SendRequest([]*payload{newNewScreenPayload(name)}, nil)
}

// ClearAll function.
func (r *Ray) ClearAll() *Ray {
	return r.SendRequest([]*payload{newClearAllPayload()}, nil)
}

// ClearScreen function.
func (r *Ray) ClearScreen() *Ray {
	return r.NewScreen("")
}

// Color function.
func (r *Ray) Color(color Color) *Ray {
	return r.SendRequest([]*payload{newColorPayload(color)}, nil)
}

// Size function.
func (r *Ray) Size(size Size) *Ray {
	return r.SendRequest([]*payload{newSizePayload(size)}, nil)
}

// Remove function.
func (r *Ray) Remove() *Ray {
	return r.SendRequest([]*payload{newRemovePayload()}, nil)
}

// Hide function.
func (r *Ray) Hide() *Ray {
	return r.SendRequest([]*payload{newHidePayload()}, nil)
}

// Notify function.
func (r *Ray) Notify(text string) *Ray {
	return r.SendRequest([]*payload{newNotifyPayload(text)}, nil)
}

// ToJSON function.
func (r *Ray) ToJSON(arguments ...interface{}) *Ray {
	if len(arguments) == 0 {
		return r
	}

	payloads := make([]*payload, len(arguments))

	for key, argument := range arguments {
		payloads[key] = newJSONStringPayload(argument)
	}

	return r.SendRequest(payloads, nil)
}

// Time function.
func (r *Ray) Time(t time.Time, format string) *Ray {
	return r.SendRequest([]*payload{newTimePayload(t, format)}, nil)
}

// Ban function.
func (r *Ray) Ban() *Ray {
	return r.Send("🕶")
}

// Charles function.
func (r *Ray) Charles() *Ray {
	return r.Send("🎶 🎹 🎷 🕺")
}

// Table function.
func (r *Ray) Table(values []interface{}, label string) *Ray {
	return r.SendRequest([]*payload{newTableType(values, label)}, nil)
}

// HTML function.
func (r *Ray) HTML(html string) *Ray {
	return r.SendRequest([]*payload{newHTMLPayload(html)}, nil)
}

// Raw function.
func (r *Ray) Raw(arguments ...interface{}) *Ray {
	if len(arguments) == 0 {
		return r
	}

	return r.SendRequest([]*payload{
		newLogPayload(arguments...),
	}, nil)
}

// Send wrapper around DefaultRay.Send.
func Send(arguments ...interface{}) *Ray {
	return DefaultRay.Send(arguments...)
}

// Send function.
func (r *Ray) Send(arguments ...interface{}) *Ray {
	if len(arguments) == 0 {
		return r
	}

	if r.settings.AlwaysSendRawValues {
		return r.Raw(arguments...)
	}

	return r.SendRequest(
		createPayloadsForValues(arguments...),
		nil,
	)
}

// Pass function.
func (r *Ray) Pass(argument interface{}) interface{} {
	r.Send(argument)

	return argument
}

// ShowApp function.
func (r *Ray) ShowApp() *Ray {
	return r.SendRequest([]*payload{newShowAppPayload()}, nil)
}

// HideApp function.
func (r *Ray) HideApp() *Ray {
	return r.SendRequest([]*payload{newHideAppPayload()}, nil)
}

// SendCustom function.
func (r *Ray) SendCustom(content, label string) *Ray {
	return r.SendRequest([]*payload{newCustomPayload(content, label)}, nil)
}

// SendRequest function.
func (r *Ray) SendRequest(payloads []*payload, meta map[string]interface{}) *Ray {
	if payloads == nil {
		return r
	}

	if r.Disabled() {
		return r
	}

	for key := range payloads {
		payloads[key].Origin = newOrigin(3)
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
