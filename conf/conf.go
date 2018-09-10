package conf

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

var (
	// Conf holds the global app config.
	Conf              config
	defaultConfigFile = "conf/conf.toml"
	// DBConfig hmm
	DBConfig string
)

type config struct {
	Mode string
	DB   database `toml:"database"`
}

type database struct {
	DBName string `toml:"db_name"`
	DBUser string `toml:"db_user"`
	DBPwd  string `toml:"db_pwd"`
	DBHost string `toml:"db_host"`
	DBPort string `toml:"db_port"`
}

// InitConfig initializes the app configuration by first setting defaults,
// then overriding settings from the app config file, then overriding
// It returns an error if any.
func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	_, err := toml.DecodeFile(configFile, &Conf)

	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}

	DBConfig = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", Conf.DB.DBUser, Conf.DB.DBPwd, Conf.DB.DBHost, Conf.DB.DBPort, Conf.DB.DBName)
	return nil
}

func init() {
	InitConfig("")
}