package ray

type payloadType string

const (
	colorType payloadType = "color"
	logType   payloadType = "log"
)

type payload struct {
	Type    payloadType            `json:"type"`
	Content map[string]interface{} `json:"content"`
	Origin  *origin                `json:"origin"`
}

func newColorPayload(color Color) *payload {
	return &payload{
		Type: colorType,
		Content: map[string]interface{}{
			"color": color,
		},
		Origin: newOrigin(),
	}
}

func newLogPayload(values ...interface{}) *payload {
	return &payload{
		Type: logType,
		Content: map[string]interface{}{
			"values": values,
		},
		Origin: newOrigin(),
	}
}
