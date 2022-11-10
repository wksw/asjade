package logger

import "testing"

func TestLogger(t *testing.T) {
	Info("a", "b", 1, 2)
	Infof("a %s", "abc")
	Debug("d", "e", "b", "u", "g")
}
