package internal

type Config struct {
	Host     string      `yaml:"host"`
	Port     int         `yaml:"port"`
	LogLevel string      `yaml:"log_level"`
	CacheDir string      `yaml:"cache_dir"`
	DB       MySQLConfig `yaml:"db"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "127.0.0.1",
		Port:     8080,
		LogLevel: "info",
		CacheDir: "data",
	}
}
