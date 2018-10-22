package conf

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

type config struct {
	Mode   string
	DB     database `toml:"database"`
	Assets assets   `toml:"assets"`
}

type assets struct {
	Avatar string `toml:"avatar"`
}

type database struct {
	DBName string `toml:"db_name"`
	DBUser string `toml:"db_user"`
	DBPwd  string `toml:"db_pwd"`
	DBHost string `toml:"db_host"`
	DBPort string `toml:"db_port"`
}

var (
	// Conf holds the global app config.
	Conf              config
	defaultConfigFile = "conf/conf.toml"
	// DBConfig hmm
	DBConfig string
)

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

	DBConfig = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", Conf.DB.DBUser, Conf.DB.DBPwd, Conf.DB.DBHost, Conf.DB.DBPort, Conf.DB.DBName)
	return nil
}

func init() {
	InitConfig("")
}
