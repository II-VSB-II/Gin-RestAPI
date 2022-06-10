package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Application Application `json:"application"`
	Logging     Logging     `json:"logging"`
	Database    Database    `json:"database"`
}

type Application struct {
	PageSize             int64    `json:"PageSize"`
	JwtSecret            int64    `json:"JwtSecret"`
	ImageSavePath        string   `json:"ImageSavePath"`
	ImageMaxSize         int64    `json:"ImageMaxSize"`
	ImageAllowExtensions []string `json:"ImageAllowExtensions"`
	RunMode              string   `json:"RunMode"`
	HTTPPort             int64    `json:"HttpPort"`
	ReadTimeout          int64    `json:"ReadTimeout"`
	WriteTimeout         int64    `json:"WriteTimeout"`
}

type Database struct {
	Type     string `json:"Type"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Host     string `json:"Host"`
	Port     int64  `json:"Port"`
	Name     string `json:"Name"`
}

type Logging struct {
	LogSavePath      string `json:"LogSavePath"`
	LogSaveName      string `json:"LogSaveName"`
	LogFileExtension string `json:"LogFileExtension"`
	TimeFormat       int64  `json:"TimeFormat"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("mlp")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error in reading the mlp.yaml file.")
	}
	err = viper.Unmarshal(&config)
	return
}
