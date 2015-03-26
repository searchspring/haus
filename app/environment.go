package haus

// Environment represents a single environment from the haus config file.
type Environment struct {
	Name string
	Requirements []string
	Variables map[string]string
}