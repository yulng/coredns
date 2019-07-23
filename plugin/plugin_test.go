package plugin

import (
	"testing"

	"github.com/caddyserver/caddy"
)

func TestOrigins(t *testing.T) {
	c := caddy.NewTestController("dns", "")
	c.ServerBlockKeys = []string{"example.net", "example.org"}

	c.ServerBlockKeyIndex = 0
	if err := Origins(c, []string{"example.net"}); err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}
	c.ServerBlockKeyIndex = 0
	if err := Origins(c, []string{"sub.example.net"}); err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}
	c.ServerBlockKeyIndex = 1
	if err := Origins(c, []string{"example.net"}); err == nil {
		t.Errorf("Expected ErrOrigin, got nil")
	}
}
