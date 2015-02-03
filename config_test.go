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

func TestConfigBool(t *testing.T) {
	c := LoadConfigString(`{"istrue":true,"isfalse":false}`)
	if c == nil {
		t.Fatalf("expected a config object")
	}
	if c.GetBool("istrue") == false {
		t.Errorf("expected true")
	}
	if c.GetBool("isfalse") == true {
		t.Errorf("expected false")
	}
}

func TestConfigMerge(t *testing.T) {
	c := LoadConfigString(`{"one":1,"two":"zwei","three":["a","b","c"]}`)
	if c == nil {
		t.Fatalf("expected a config object")
	}
	c.StringMerge(`{"two":2,"four":"vier"}`)
	if c.GetInt("one") != 1 {
		t.Errorf("expected 1, got %d:", c.GetInt("one"))
	}
	if len(c.GetArray("three")) != 3 {
		t.Errorf("expected 3 elt array")
	}
	if c.GetInt("two") != 2 {
		t.Errorf("expected 2, got %d:", c.GetInt("two"))
	}
	if c.GetString("four") != "vier" {
		t.Errorf("expected vier, got %s:", c.GetString("four"))
	}
}

func TestGetConfig(t *testing.T) {
	c := LoadConfigString(`{"one":{"two":1,"three":"zwei"}}`)
	if c == nil {
		t.Fatalf("expected a config object")
	}
	subc := c.GetConfig("one")
	if subc == nil {
		t.Fatalf("expected sub-config object")
	}
	if subc.GetInt("two") != 1 {
		t.Fatalf("expected int in sub-config")
	}
	if subc.GetString("three") != "zwei" {
		t.Fatalf("expected string in sub-config")
	}
}
