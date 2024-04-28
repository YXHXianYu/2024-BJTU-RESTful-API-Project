package config

type Config struct {
	Database struct {
		DSN string
	}
}

func GetConfig() *Config {
	return &Config{
		Database: struct {
			DSN string
		}{
			DSN: "root:20021012@tcp(127.0.0.1:3306)/restful_api?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
}
