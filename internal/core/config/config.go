package config

type Config struct {
	Env      string `yaml:"env" env-default:"local"`
	Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
