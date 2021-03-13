package ray

type payloadType string

const (
	colorType payloadType = "color"
	logType   payloadType = "log"
	sizeType  payloadType = "size"
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
		Origin: newOrigin(4),
	}
}

func newLogPayload(values ...interface{}) *payload {
	return &payload{
		Type: logType,
		Content: map[string]interface{}{
			"values": values,
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
