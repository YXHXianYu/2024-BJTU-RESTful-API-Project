package config

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
	}
}
