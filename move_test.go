package chester_test

import (
	"testing"

	"github.com/bluescreen10/chester"
)

func TestParseMove(t *testing.T) {
	pos, _ := chester.ParseFEN(chester.DefaultFEN)

	tests := []struct {
		name     string
		move     string
		expected string
		wantErr  bool
	}{
		{
			name:     "Standard move e2e4",
			move:     "e2e4",
			expected: "e2e4",
			wantErr:  false,
		},
		{
			name:     "Knight move g1f3",
			move:     "g1f3",
			expected: "g1f3",
			wantErr:  false,
		},
		{
			name:     "Promotion move e7e8q",
			move:     "e7e8q",
			expected: "e7e8q",
			wantErr:  false,
		},
		{
			name:    "Invalid move",
			move:    "invalid",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := chester.ParseMove(test.move, pos)
			if (err != nil) != test.wantErr {
				t.Errorf("ParseMove(%s) error = %v, wantErr %v", test.move, err, test.wantErr)
				return
			}
			if !test.wantErr {
				if got := m.String(); got != test.expected {
					t.Errorf("ParseMove(%s) = %s, want %s", test.move, got, test.expected)
				}
			}
		})
	}
}
