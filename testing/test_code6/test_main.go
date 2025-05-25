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
			name:     "says Hello in English to a named person",
			input:    "Alice",
			lang:     "",
			expected: "Hello, Alice",
		},
		{
			name:     "defaults to World in English",
			input:    "",
			lang:     "",
			expected: "Hello, World",
		},
		{
			name:     "says Hola in Spanish to a named person",
			input:    "Carlos",
			lang:     "Spanish",
			expected: "Hola, Carlos",
		},
		{
			name:     "defaults to World in Spanish",
			input:    "",
			lang:     "Spanish",
			expected: "Hola, World",
		},
		{
			name:     "says Bonjour in French to a named person",
			input:    "Claire",
			lang:     "French",
			expected: "BonjourClaire",
		},
		{
			name:     "defaults to World in French",
			input:    "",
			lang:     "French",
			expected: "BonjourWorld",
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
