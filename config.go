// Copyright 2011 Numerotron Inc.
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.
//
// Developed at www.stathat.com by Patrick Crosby
// Contact us on twitter with any questions:  twitter.com/stat_hat

// The jconfig package provides a simple, basic configuration file parser using JSON.
package jconfig

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Data     map[string]interface{}
	filename string
}

func newConfig() *Config {
	result := new(Config)
	result.Data = make(map[string]interface{})
	return result
}

// Loads config information from a JSON file
func LoadConfig(filename string) *Config {
	result := newConfig()
	result.filename = filename
	err := result.parse()
	if err != nil {
		log.Fatalf("error loading config file %s: %s", filename, err)
	}
	return result
}

// Loads config information from a JSON string
func LoadConfigString(s string) *Config {
	result := newConfig()
	err := json.Unmarshal([]byte(s), &result.Data)
	if err != nil {
		log.Fatalf("error parsing config string %s: %s", s, err)
	}
	return result
}

func (c *Config) StringMerge(s string) {
	next := LoadConfigString(s)
	c.Merge(next.Data)
}

func (c *Config) LoadMerge(filename string) {
	next := LoadConfig(filename)
	c.Merge(next.Data)
}

func (c *Config) Merge(nData map[string]interface{}) {
	for k, v := range nData {
		c.Data[k] = v
	}
}

func (c *Config) parse() error {
	f, err := os.Open(c.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	b := new(bytes.Buffer)
	_, err = b.ReadFrom(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b.Bytes(), &c.Data)
	if err != nil {
		return err
	}

	return nil
}

// Returns a string for the config variable key
func (c *Config) GetString(key string) string {
	result, present := c.Data[key]
	if !present {
		return ""
	}
	return result.(string)
}

// Returns an int for the config variable key
func (c *Config) GetInt(key string) int {
	x, ok := c.Data[key]
	if !ok {
		return -1
	}
	return int(x.(float64))
}

// Returns a float for the config variable key
func (c *Config) GetFloat(key string) float64 {
	x, ok := c.Data[key]
	if !ok {
		return -1
	}
	return x.(float64)
}

// Returns a bool for the config variable key
func (c *Config) GetBool(key string) bool {
	x, ok := c.Data[key]
	if !ok {
		return false
	}
	return x.(bool)
}

// Returns an array for the config variable key
func (c *Config) GetArray(key string) []interface{} {
	result, present := c.Data[key]
	if !present {
		return []interface{}(nil)
	}
	return result.([]interface{})
}
