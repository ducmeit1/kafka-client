package viper

import (
	"bytes"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
)

const DefaultConfigType = "toml"

func LoadConfigFileByViper(configType string, isMerge bool, values []byte) error {
	viper.SetConfigType(configType)
	r := bytes.NewReader(values)
	defer io.Copy(ioutil.Discard, r)

	if isMerge {
		return viper.MergeConfig(r)
	}

	return viper.ReadConfig(r)
}
