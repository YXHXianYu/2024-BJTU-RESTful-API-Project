package config

<<<<<<< HEAD
import "github.com/gin-contrib/cors"

func GetDSN() string {
	return "root:20021012@tcp(127.0.0.1:3306)/restful_api?charset=utf8mb4&parseTime=True&loc=Local"
}

func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		AllowCredentials: true,
=======
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
>>>>>>> 21de4b8415dfe06201b4e0052e7c443b0dae38a8
	}
}
