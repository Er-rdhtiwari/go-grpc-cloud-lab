package stringsx

import "testing"

func TestNormalizeSpace(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{"already clean", "hello world", "hello world"},
		{"multiple spaces", "hello    world", "hello world"},
		{"tabs/newlines", "hello\t\nworld", "hello world"},
		{"leading/trailing", "   hello world   ", "hello world"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		tt := tt // capture range var
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NormalizeSpace(tt.in)
			if got != tt.want {
				t.Fatalf("NormalizeSpace(%q)=%q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"empty", "", true},
		{"spaces", "   ", true},
		{"tabs/newlines", "\t\n", true},
		{"non-empty", "a", false},
		{"trimmed non-empty", "  a  ", false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := IsBlank(tt.in)
			if got != tt.want {
				t.Fatalf("IsBlank(%q)=%v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
