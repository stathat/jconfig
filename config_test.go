package jconfig

import (
	"testing"
)

func TestConfigFromString(t *testing.T) {
	c := LoadConfigString(`{"one":1,"two":"zwei","three":["a","b","c"]}`)
	if c == nil {
		t.Fatalf("expected a config object")
	}
	if c.data == nil {
		t.Fatalf("expected config object to have data")
	}
	if c.GetInt("one") != 1 {
		t.Errorf("expected 1, got %d:", c.GetInt("one"))
	}
	if c.GetString("two") != "zwei" {
		t.Errorf("expected zwei, got %s:", c.GetString("two"))
	}
	if len(c.GetArray("three")) != 3 {
		t.Errorf("expected 3 elt array")
	}
}
