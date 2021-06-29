package setting

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Mode         string
	Port         string
	ReadTimeout  int `mapstructure:"read-timeout"`
	WriteTimeout int `mapstructure:"write-timeout"`
}

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Name     string
}

type Setting struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

// Load is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Load() Setting {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("application")
	v.AddConfigPath(".")

	bindEnv(v)

	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("error on loading configuration file")
	}

	var setting Setting
	err = v.Unmarshal(&setting)
	if err != nil {
		log.Fatal("error on unmarshalling configuration file")
	}

	settingJson, _ := json.Marshal(setting)
	log.Printf("running app with setting: %s\n", settingJson)

	return setting
}

func bindEnv(v *viper.Viper) {
	v.BindEnv("database.host", "DBHOST")
	v.BindEnv("database.username", "DBUSERNAME")
	v.BindEnv("database.password", "DBPASSWORD")

	v.BindEnv("server.mode", "SERVER_MODE")
	v.BindEnv("server.port", "SERVER_PORT")
	v.BindEnv("server.read-timeout", "SERVER_READTIMEOUT")
	v.BindEnv("server.write-timeout", "SERVER_WRITETIMEOUT")
}