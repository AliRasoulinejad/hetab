package config

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var C *Config

type Config struct {
	Logger     Logger      `yaml:"logger"`
	HTTPServer HTTPServer  `yaml:"http_server"`
	Database   SQLDatabase `yaml:"database"`
}

type Logger struct {
	Level string `yaml:"level"`
}

func (c *Config) String() string {
	return fmt.Sprintf("HTTP\t\t|%s\n", c.HTTPServer)
}

type HTTPServer struct {
	Address           string        `yaml:"address"`
	Port              int           `yaml:"port"`
	Listen            string        `yaml:"listen"`
	ReadTimeout       time.Duration `yaml:"read_Timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
}

func (i HTTPServer) String() string {
	// return fmt.Sprintf("Listen to: %s", i.Listen)
	return fmt.Sprintf("Listen to: %s", fmt.Sprintf("%s:%d", i.Address, i.Port))
}

type SQLDatabase struct {
	Driver        string        `yaml:"driver"`
	Host          string        `yaml:"host"`
	Port          int           `yaml:"port"`
	DB            string        `yaml:"db"`
	User          string        `yaml:"user"`
	Password      string        `yaml:"password"`
	TimeZone      string        `yaml:"time_zone"`
	MaxConn       int           `yaml:"max_conn"`
	IdleConn      int           `yaml:"idle_conn"`
	Timeout       time.Duration `yaml:"timeout"`
	DialRetry     int           `yaml:"dial_retry"`
	DialTimeout   time.Duration `yaml:"dial_timeout"`
	ReadTimeout   time.Duration `yaml:"read_timeout"`
	WriteTimeout  time.Duration `yaml:"write_timeout"`
	UpdateTimeout time.Duration `yaml:"update_timeout"`
	DeleteTimeout time.Duration `yaml:"delete_timeout"`
	QueryTimeout  time.Duration `yaml:"query_timeout"`
}

func (d SQLDatabase) DSN() (dsn string) {
	switch d.Driver {
	case "postgres":
		// dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", d.Host, d.User, d.Password, d.DB, d.Port, d.TimeZone)
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", d.Host, d.User, d.Password, d.DB, d.Port)
	default:
		log.Fatalf("SQLDatabase driver is not supported: %s", d.Driver)
	}

	return
}

// String represents most usable values for the SQLDatabase config
func (d SQLDatabase) String() (str string) {
	switch d.Driver {
	case "postgres":
		// str = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", d.Host, d.User, d.Password, d.DB, d.Port, d.TimeZone)
		str = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", d.Host, d.User, d.Password, d.DB, d.Port)
	default:
		log.Fatalf("SQLDatabase driver is not supported: %s", d.Driver)
	}

	return
}

func Init(filename string) {
	c := new(Config)
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigFile(filename)
	v.SetConfigType("yml")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("failed on config `%s` unmarshal: %v", filename, err)
	}

	err = v.Unmarshal(c)
	if err != nil {
		log.Fatalf("load configurations failed, error: %v", err)
	}

	C = c
	log.Infof("config file opened successfully. filename: %s", filename)
}
