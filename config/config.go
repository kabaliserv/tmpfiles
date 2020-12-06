package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
)

// Conf struct
type Conf struct {
	AppName        string
	Description    string
	Web            WebConf
	Store          StoreConf
	ConfigPathFile string
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

var server Conf

// Valid : Get if Store is indicate and is existe on filesystem
func (state *StoreConf) Valid() error {

	errString := "Config: [store] property"

	// Get if path is indicate
	if state.Path == "" {
		return fmt.Errorf("%v %v", errString, "path cannot be found or is empty")
	}

	// Get if path exist
	if _, err := os.Stat(state.Path); os.IsNotExist(err) {
		return fmt.Errorf("%v %v", errString, "property path is not exist")
	}

	absPath, err := filepath.Abs(state.Path)
	if err != nil {
		log.Print(err)
	}

	// Get if path is directory
	if fileInfo, _ := os.Stat(absPath); !fileInfo.IsDir() {
		return fmt.Errorf("%v %v", errString, "property path is not directory")
	}

	state.Path = absPath

	return nil

}

// SetConfigOptions add config server to singleton
/* func SetConfigOptions(conf Conf) {
	config = conf
} */

// ParseConfigFile : Parse Config file to Struct
func ParseConfigFile(path string) error {
	// Get Config file to array bytes
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Parse Config file
	server = Conf{}
	err = toml.Unmarshal(src, &server)
	if err != nil {
		return err
	}

	return nil
}

// ValidConf : Valide Conf and set default option if not found on this
func (state *Conf) ValidConf() error {

	// Valid Name APP
	if state.AppName == "" {
		state.AppName = "TMPFiles"
	}

	// Valid Address
	if ok, _ := regexp.MatchString(`^((25[0-5]|2[0-4][0-9]|1?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1?[0-9]?[0-9])$`, state.Web.Address); !ok {
		state.Web.Address = "0.0.0.0"
	}

	// Valid Port
	if state.Web.Port == 0 {
		state.Web.Port = 3000
	}
	if state.Web.Port < 1 || state.Web.Port > 65535 {
		return fmt.Errorf("Config: [webconf] property port is invalid")
	}

	// Valid Web Root
	if ok, _ := regexp.MatchString(`^/([-_/\.\w\W])*$`, state.Web.Path); !ok {
		state.Web.Path = "/"
	}

	// Valid Store Path
	if err := state.Store.Valid(); err != nil {
		return err
	}

	return nil
}

// GetConfigOptions is finction to get config server
func GetConfigOptions() *Conf {
	return &server
}

// GetStorePath : return store path directory
func GetStorePath() string {
	return server.Store.Path
}

// GetWebAddr : return listen address
func GetWebAddr() string {
	return server.Web.Address
}

// GetWebPort : return listen Port
func GetWebPort() string {
	return fmt.Sprint(server.Web.Port)
}

// GetWebPath : return web root path
func GetWebPath() string {
	return server.Web.Path
}
