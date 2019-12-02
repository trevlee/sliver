package logger

import (
	"fmt"
	"strings"
)

// ScriptLogger can be used to store script outputs
type ScriptLogger struct {
	stdout strings.Builder
	stderr strings.Builder
}

// Info writes a message to the buffer
func (s *ScriptLogger) Info(message string) {
	s.stdout.WriteString(fmt.Sprintf("[*] %s\n", message))
}

// Raw writes unaltered message to stdout
func (s *ScriptLogger) Raw(message string) {
	s.stdout.WriteString(message)
}

// Error writes an error message
func (s *ScriptLogger) Error(message string) {
	s.stderr.WriteString(fmt.Sprintf("[!] %s\n", message))
}

// String returns the concatenation of stdout and stderr
func (s *ScriptLogger) String() string {
	return s.stdout.String() + s.stderr.String()
}
