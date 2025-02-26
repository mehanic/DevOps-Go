package main

import "fmt"

// Config struct to hold configuration
type Config struct {
	username string
	age      int
}

// Option is a function type that modifies the Config
type Option func(*Config)

// WithUsername sets the username in the Config
func WithUsername(name string) Option {
	return func(c *Config) {
		c.username = name
	}
}

// NewConfig creates a new Config object with given options
func NewConfig(opts ...Option) *Config {
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func main() {
	// Initialize the configuration with the WithUsername option
	cfg := NewConfig(WithUsername("Jane"))
	fmt.Println("Username:", cfg.username) // Output: Username: Jane
}
