package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string     `mapstructure:"env"`
	LogFile    string     `mapstructure:"logfile"`
	Storage    Storage    `mapstructure:"storage"`
	HTTPServer HTTPServer `mapstructure:"http"`
	JWT        JWT        `mapstructure:"jwt"`
	OTP        OTP        `mapstructure:"otp"`
	Smsc       Smsc       `mapstructure:"smsc"`
}

type HTTPServer struct {
	Port string `mapstructure:"address"`
}

type JWT struct {
	Secret  string        `mapstructure:"secret"`
	Expires time.Duration `mapstructure:"expiration"`
}

type OTP struct {
	Length  int           `mapstructure:"length"`
	Expires time.Duration `mapstructure:"expiration"`
}

type Smsc struct {
	Login    string `mapsstructure:"login"`
	Password string `mapstructure:"password"`
}

type Storage struct {
	Driver   string `mapstructure:"driver"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

func MustLoadConfig() Config {
	configPath := flag.String("config", "./config.yaml", "path to configure the project")
	flag.Parse()
	if *configPath == "" {
		log.Fatal("could not get config path")
	}

	//check if config file exists

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %v", err)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*configPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper could not read config file:%v", err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("viper could not unmarshal to config struct:%v", err)

	}
	return cfg
}
