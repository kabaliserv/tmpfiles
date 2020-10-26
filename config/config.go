package config

import (
	"fmt"
)

// Conf struct
type Conf struct {
	AppName     string
	Description string
	Web         WebConf
	Store       StoreConf
}

// WebConf struct
type WebConf struct {
	Address string
	Port    int
	Domain  string
	Path    string
}

// StoreConf struct
type StoreConf struct {
	Path string
}

// Valid :
func (state *StoreConf) Valid() error {
	errString := "Config: [store] property"
	if state.Path == "" {
		return fmt.Errorf("%v %v", errString, "path cannot be found or is empty")
	}
	return nil
}
