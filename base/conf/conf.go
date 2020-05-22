package conf

import (
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"time"
	log "xxblog/base/logger"
)

var (
	confPath string
	Conf     *Config
)

func Init() {
	flag.StringVar(&confPath, "c", "", "config file path")
	flag.Parse()
	if confPath == "" {
		viper.SetConfigName("config")       // name of config file (without extension)
		viper.SetConfigType("yml")          // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("/etc/xxblog/") // path to look for the config file in
		viper.AddConfigPath("./base/conf/") // call multiple times to add many search paths
		viper.AddConfigPath("../base/conf") // optionally look for config in the working directory
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

	//初始化日志
	initLog()
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

func initLog() {
	c := log.New()
	c.SetDivision("time")    // 设置归档方式，"time"时间归档 "size" 文件大小归档，文件大小等可以在配置文件配置
	c.SetTimeUnit(log.Day)   // 时间归档 可以设置切割单位
	c.SetEncoding("console") // 输出格式 "json" 或者 "console"
	if Conf.Debug {
		c.Debug()
	}
	if !Conf.Log.Stdout {
		c.CloseConsoleDisplay()
	}
	c.SetInfoFile(Conf.Log.Info)
	c.SetErrorFile(Conf.Log.Error)
	c.InitLogger()
}
