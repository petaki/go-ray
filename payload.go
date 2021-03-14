package ray

import (
	"encoding/json"
	"fmt"
	"time"
)

type payloadType string

const (
	clearAllType   payloadType = "clear_all"
	colorType      payloadType = "color"
	customType     payloadType = "custom"
	hideType       payloadType = "hide"
	hideAppType    payloadType = "hide_app"
	jsonStringType payloadType = "json_string"
	logType        payloadType = "log"
	newScreenType  payloadType = "new_screen"
	notifyType     payloadType = "notify"
	removeType     payloadType = "remove"
	showAppType    payloadType = "show_app"
	sizeType       payloadType = "size"
	tableType      payloadType = "table"
	timeType       payloadType = "carbon"
)

type payload struct {
	Type    payloadType            `json:"type"`
	Content map[string]interface{} `json:"content"`
	Origin  *origin                `json:"origin"`
}

func createPayloadsForValues(values ...interface{}) []*payload {
	payloads := make([]*payload, len(values))

	for key, value := range values {
		payloads[key] = createPayloadForValue(value)
	}

	return payloads
}

func createPayloadForValue(value interface{}) *payload {
	switch v := value.(type) {
	case bool:
		return newBoolPayload(v)
	case time.Time:
		return newTimePayload(v, "")
	case map[string]interface{}:
		return newJSONStringPayload(v)
	default:
		if v == nil {
			return newNilPayload()
		}

		return newLogPayload(v)
	}
}

func convertToPrimitive(value interface{}) interface{} {
	return fmt.Sprintf("%+v", value)
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

func newClearAllPayload() *payload {
	return &payload{
		Type:    clearAllType,
		Content: map[string]interface{}{},
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

func newCustomPayload(content, label string) *payload {
	return &payload{
		Type: customType,
		Content: map[string]interface{}{
			"content": content,
			"label":   label,
		},
	}
}

func newHidePayload() *payload {
	return &payload{
		Type:    hideType,
		Content: map[string]interface{}{},
	}
}

func newHideAppPayload() *payload {
	return &payload{
		Type:    hideAppType,
		Content: map[string]interface{}{},
	}
}

func newHTMLPayload(html string) *payload {
	return &payload{
		Type: customType,
		Content: map[string]interface{}{
			"content": html,
			"label":   "HTML",
		},
	}
}

func newJSONStringPayload(value interface{}) *payload {
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
		values[key] = convertToPrimitive(value)
	}

	return &payload{
		Type: logType,
		Content: map[string]interface{}{
			"values": values,
		},
	}
}

func newNewScreenPayload(name string) *payload {
	return &payload{
		Type: newScreenType,
		Content: map[string]interface{}{
			"name": name,
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

func newNotifyPayload(text string) *payload {
	return &payload{
		Type: notifyType,
		Content: map[string]interface{}{
			"value": text,
		},
	}
}

func newRemovePayload() *payload {
	return &payload{
		Type:    removeType,
		Content: map[string]interface{}{},
	}
}

func newShowAppPayload() *payload {
	return &payload{
		Type:    showAppType,
		Content: map[string]interface{}{},
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

func newTableType(values []interface{}, label string) *payload {
	for key, value := range values {
		values[key] = convertToPrimitive(value)
	}

	if label == "" {
		label = "Table"
	}

	return &payload{
		Type: tableType,
		Content: map[string]interface{}{
			"values": values,
			"label":  label,
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
