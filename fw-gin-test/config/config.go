package config

type Config struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          int64
	Name          string
	Mode          string
	DemoMsg       string
	EnableDP      bool
}

// var GinConfig = new(Config)
var GinConfig = &Config{
	ReadTimeout:   1000,
	WriterTimeout: 1000,
	Host:          "127.0.0.1",
	Port:          8901,
	Name:          "go0base",
	Mode:          "dev",
	DemoMsg:       "dddddd",
	EnableDP:      false,
}
