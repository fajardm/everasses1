package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

// LoadFile .
func LoadFile(path string) (config *Config, err error) {
	defer func() {
		if re := recover(); re != nil {
			err = re.(error)
		}
		return
	}()

	var (
		dir      = filepath.Dir(path)
		filename = filepath.Base(path)
		ext      = filepath.Ext(path)[1:]
	)

	viper.AddConfigPath(dir)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return
}
