package conf

import (
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"time"
)

var (
	confPath string
	Conf     *Config
)

func init() {
	flag.StringVar(&confPath, "c", "", "config file path")
	flag.Parse()
	if confPath == "" {
		viper.SetConfigName("config")       // name of config file (without extension)
		viper.SetConfigType("yml")          // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("/etc/xxblog/") // path to look for the config file in
		viper.AddConfigPath("./conf/")      // call multiple times to add many search paths
		viper.AddConfigPath(".")            // optionally look for config in the working directory
	} else {
		viper.SetConfigFile(confPath)
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	c := &Config{}
	opt := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))
	if err := viper.Unmarshal(&c, opt); err != nil {
		fmt.Println(err)
	}
	Conf = c
}

// Config config
type Config struct {
	Debug bool   `yaml:"debug"`
	Log   *Log   `yaml:"log"`
	Mysql *Mysql `yaml:"mysql" mapstructure:"mysql"`
	Redis *Redis `yaml:"redis" mapstructure:"redis"`
	Init  *InitInfo
}
type InitInfo struct {
	DB bool
}

type Log struct {
	Dir    string
	Error  string `yaml:"error"`
	Info   string
	Stdout bool
}

type Mysql struct {
	DSN               string
	MaxConnection     int
	MaxIdleConnection int
	MaxLifeTime       time.Duration
}

type Redis struct {
	Network      string        `yaml:"network"`
	Addr         string        `yaml:"addr"`
	Password     string        `yaml:"password"`
	DB           int           `yaml:"db"`
	DialTimeout  time.Duration `yaml:"dialTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	PoolSize     int
	MinIdleConn  int
	IdleTimeout  time.Duration
	PubSubChan   string
}
