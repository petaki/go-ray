package ray

import "testing"

func TestNew(t *testing.T) {
	r := New(nil)

	if r == nil {
		t.Errorf("expected: not nil, got: nil")
	}

	r = New(&Client{
		Host: "192.168.100.1",
		Port: 23517,
	})

	if r == nil {
		t.Errorf("expected: not nil, got: nil")
	}
}
