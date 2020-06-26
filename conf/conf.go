package conf

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	confPath string
	Debug    bool
	WebName  string
)

type Log struct {
	Filename  string `yaml:"FileName" `
	MaxSize   int    `yaml:"MaxSize" `
	LocalTime bool   `yaml:"LocalTime" `
	Compress  bool   `yaml:"Compress" `
	Level     string `yaml:"Level" `
}

type DB struct {
	Prefix   string `yaml:"prefix"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Server struct {
	Address string `yaml:"address"`
	WebName string `yaml:"webName"`
}

type Config struct {
	LogConf    *Log    `yaml:"log"`
	DBConf     *DB     `yaml:"db"`
	ServerConf *Server `yaml:"server"`
}

func init() {
	flag.StringVar(&confPath, "conf", "app.yaml", "default config path")
	flag.BoolVar(&Debug, "debug", false, "env")

}

func Init() (c *Config, err error) {
	var f *os.File
	var b []byte
	if f, err = os.Open(confPath); err != nil {
		return
	}

	if b, err = ioutil.ReadAll(f); err != nil {
		return
	}

	if err = yaml.Unmarshal(b, &c); err != nil {
		return
	}

	WebName = "写代码的放羊小子"
	if len(c.ServerConf.WebName) > 0 {
		WebName = c.ServerConf.WebName
	}

	return
}
