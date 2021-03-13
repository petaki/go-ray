package ray

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name            string
		settings        *Settings
		expectedEnabled bool
		expectedHost    string
		expectedPort    int
	}{
		{"Default Settings", nil, true, "localhost", 23517},
		{"Custom Settings", &Settings{
			Enable: false,
			Host:   "192.168.100.1",
			Port:   23518,
		}, false, "192.168.100.1", 23518},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(tt.settings)

			if r == nil {
				t.Errorf("expected: not nil, got: nil")
			} else {
				if r.settings == nil {
					t.Errorf("expected: not nil, got: nil")
				} else {
					if r.settings.Host != tt.expectedHost {
						t.Errorf("expected: %s, got: %s", tt.expectedHost, r.settings.Host)
					}

					if r.settings.Port != tt.expectedPort {
						t.Errorf("expected: %d, got: %d", tt.expectedPort, r.settings.Port)
					}

					if r.enabled != tt.expectedEnabled {
						t.Errorf("expected: %v, got: %v", tt.expectedEnabled, r.enabled)
					}
				}
			}
		})
	}
}

func TestEnable(t *testing.T) {
	r := New(&Settings{
		Enable: false,
		Host:   "192.168.100.1",
		Port:   23518,
	})

	if r.Enabled() {
		t.Errorf("expected: false, got: %v", r.Enabled())
	}

	r.Enable()

	if r.Disabled() {
		t.Errorf("expected: false, got: %v", r.Disabled())
	}
}

func TestDisable(t *testing.T) {
	r := New(nil)

	if r.Disabled() {
		t.Errorf("expected: false, got: %v", r.Disabled())
	}

	r.Disable()

	if r.Enabled() {
		t.Errorf("expected: false, got: %v", r.Enabled())
	}
}
