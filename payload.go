package ray

import (
	"fmt"
	"time"
)

type payloadType string

const (
	colorType  payloadType = "color"
	customType payloadType = "custom"
	logType    payloadType = "log"
	timeType   payloadType = "carbon"
	sizeType   payloadType = "size"
)

type payload struct {
	Type    payloadType            `json:"type"`
	Content map[string]interface{} `json:"content"`
	Origin  *origin                `json:"origin"`
}

func createPayloadsForValues(values ...interface{}) []*payload {
	var payloads []*payload

	for _, value := range values {
		switch v := value.(type) {
		case bool:
			payloads = append(payloads, newBoolPayload(v))
			continue
		case time.Time:
			payloads = append(payloads, newTimePayload(v, ""))
			continue
		default:
			if v == nil {
				payloads = append(payloads, newNilPayload())
			} else {
				payloads = append(payloads, newLogPayload(v))
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
		Origin: newOrigin(4),
	}
}

func newColorPayload(color Color) *payload {
	return &payload{
		Type: colorType,
		Content: map[string]interface{}{
			"color": color,
		},
		Origin: newOrigin(4),
	}
}

func newLogPayload(arguments ...interface{}) *payload {
	values := make([]interface{}, len(arguments))

	for key, argument := range arguments {
		values[key] = fmt.Sprintf("%+v", argument)
	}

	return &payload{
		Type: logType,
		Content: map[string]interface{}{
			"values": values,
		},
		Origin: newOrigin(4),
	}
}

func newNilPayload() *payload {
	return &payload{
		Type: customType,
		Content: map[string]interface{}{
			"content": nil,
			"label":   "Nil",
		},
		Origin: newOrigin(4),
	}
}

func newSizePayload(size Size) *payload {
	return &payload{
		Type: sizeType,
		Content: map[string]interface{}{
			"size": size,
		},
		Origin: newOrigin(4),
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
		Origin: newOrigin(4),
	}
}
