package ray

// Ray type.
type Ray struct {
	client *Client
}

// New function.
func New(client *Client) *Ray {
	r := new(Ray)

	if client != nil {
		r.client = client
	} else {
		r.client = &Client{
			Host: "localhost",
			Port: 23517,
		}
	}

	return r
}
