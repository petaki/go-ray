package ray

import (
	"encoding/json"
	"fmt"
	"time"
)

type payloadType string

const (
	colorType      payloadType = "color"
	customType     payloadType = "custom"
	jsonStringType payloadType = "json_string"
	logType        payloadType = "log"
	timeType       payloadType = "carbon"
	sizeType       payloadType = "size"
)

type payload struct {
	Type    payloadType            `json:"type"`
	Content map[string]interface{} `json:"content"`
	Origin  *origin                `json:"origin"`
}

func createPayloadsForValues(values ...interface{}) []*payload {
	payloads := make([]*payload, len(values))

	for key, value := range values {
		switch v := value.(type) {
		case bool:
			payloads[key] = newBoolPayload(v)
			continue
		case time.Time:
			payloads[key] = newTimePayload(v, "")
			continue
		case map[string]interface{}:
			payloads[key] = newJsonStringPayload(v)
			continue
		default:
			if v == nil {
				payloads[key] = newNilPayload()
			} else {
				payloads[key] = newLogPayload(v)
			}
		}
	}

	return payloads
}

func newBoolPayload(value bool) *payload {
	return &payload{
		Type: customType,
		Content: map[string]interface{}{
			"content": value,
			"label":   "Boolean",
		},
	}
}

func newColorPayload(color Color) *payload {
	return &payload{
		Type: colorType,
		Content: map[string]interface{}{
			"color": color,
		},
	}
}

func newJsonStringPayload(value interface{}) *payload {
	v, _ := json.Marshal(value)

	return &payload{
		Type: jsonStringType,
		Content: map[string]interface{}{
			"value": string(v),
		},
	}
}

func newLogPayload(values ...interface{}) *payload {
	for key, value := range values {
		values[key] = fmt.Sprintf("%+v", value)
	}

	return &payload{
		Type: logType,
		Content: map[string]interface{}{
			"values": values,
		},
	}
}

func newNilPayload() *payload {
	return &payload{
		Type: customType,
		Content: map[string]interface{}{
			"content": nil,
			"label":   "Nil",
		},
	}
}

func newSizePayload(size Size) *payload {
	return &payload{
		Type: sizeType,
		Content: map[string]interface{}{
			"size": size,
		},
	}
}

func newTimePayload(t time.Time, format string) *payload {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}

	timezone, _ := t.Zone()

	return &payload{
		Type: timeType,
		Content: map[string]interface{}{
			"formatted": t.Format(format),
			"timestamp": t.Unix(),
			"timezone":  timezone,
		},
	}
}
