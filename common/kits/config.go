package kits

import (
	"github.com/ducmeit1/kafka-client/common/drivers/viper"
	"github.com/ducmeit1/kafka-client/common/utils/io"
	"io/ioutil"
	"strings"
)

func LoadConfig() error {
	paths := strings.Split(*ConfigPath, ",")
	if len(paths) > 0 {
		isMerge := true
		for i, p := range paths {
			if i == 0 {
				isMerge = false
			}
			fp, err := io.GetAbsolutelyFilePath(p)
			if err != nil {
				return err
			}
			f, err := ioutil.ReadFile(fp)
			if err != nil {
				return err
			}
			err = viper.LoadConfigFileByViper(*ConfigType, isMerge, f)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
