package config

import (
	"path/filepath"
	"sync"

	toml "github.com/BurntSushi/toml"
)

type http struct {
	Scheme string `toml:"scheme"`
	Host   string `toml:"host"`
	Header struct {
		UA string `toml:"ua"`
	} `toml:"header"`
}

type html struct {
	SharedDataRegex string `toml:"shared_data_regex"`
}

type download struct {
	Folder string `toml:"folder"`
}

type postgres struct {
	PWD  string `toml:"pwd"`
	User string `toml:"user"`
	DB   string `toml:"db"`
}

type worker struct {
	Concurrency int `toml:"concurrency"`
}

// TomlCfg the toml configuration file
type TomlCfg struct {
	HTTP     http     `toml:"http"`
	HTML     html     `toml:"html"`
	POSTGRES postgres `toml:"postgres"`
	Download download `toml:"download"`
	Worker   worker   `toml:"concurrency"`
}

var once sync.Once
var tomlCfg *TomlCfg

// Load load toml configuration
func Load() TomlCfg {
	once.Do(func() {
		filePath, err := filepath.Abs("./config/default.toml")
		if err != nil {
			panic(err)
		}

		_, err = toml.DecodeFile(filePath, &tomlCfg)

		if err != nil {
			panic(err)
		}
	})

	return *tomlCfg
}
