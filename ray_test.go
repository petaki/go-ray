package ray

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name         string
		settings     *Settings
		expectedHost string
		expectedPort int
	}{
		{"Default Settings", nil, "localhost", 23517},
		{"Custom Settings", &Settings{Host: "192.168.100.1", Port: 23518}, "192.168.100.1", 23518},
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
				}
			}
		})
	}
}
