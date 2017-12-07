/**
 * 設定檔操作類別，固定讀取執行檔路徑下的 project_path/config
 *
 *  Usage:
 * 	c := config.Instance()
 *	key := c.Setting.GetString(key_name)
 */
package config

import (
	"github.com/spf13/viper"

	"log"
	"errors"
)

var _instance map[string] *viper.Viper

func Instance(configName string) *viper.Viper {
	if _instance[configName] == nil {
		if _instance == nil {
			_instance = make(map[string] *viper.Viper)
		}
		_instance[configName] = getConfig(configName)
	}

	return _instance[configName]
}

func getConfig(configName string) *viper.Viper {
	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	config_err := viper.ReadInConfig()

	if config_err != nil {
		config_err = errors.New("No config file")
		log.Fatal(config_err)
	}

	return viper.GetViper()
}