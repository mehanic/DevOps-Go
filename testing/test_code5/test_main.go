package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		lang     string
		expected string
	}{
		{
			name:     "says hello to people in English",
			input:    "Alice",
			lang:     "",
			expected: "Hello, Alice",
		},
		{
			name:     "says hello to people in Spanish",
			input:    "Carlos",
			lang:     "Spanish",
			expected: "Hola, Carlos",
		},
		{
			name:     "defaults to 'World' when empty name",
			input:    "",
			lang:     "",
			expected: "Hello, World",
		},
		{
			name:     "defaults to 'World' in Spanish when empty name",
			input:    "",
			lang:     "Spanish",
			expected: "Hola, World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.input, tt.lang)
			if got != tt.expected {
				t.Errorf("Hello(%q, %q) = %q; want %q", tt.input, tt.lang, got, tt.expected)
			}
		})
	}
}
