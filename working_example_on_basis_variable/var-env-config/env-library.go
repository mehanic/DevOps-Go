package main

import (
	"fmt"
	"os"

	"github.com/qiangxue/go-env"
)

type Config struct { //Определяем структуру для выгрузки конфигураций
	Host string
	Port int
}

func Env_library() {
	_ = os.Setenv("APP_HOST", "127.0.0.1") // задаем переменные окружения
	_ = os.Setenv("APP_PORT", "8080")

	var cfg Config
	if err := env.Load(&cfg); err != nil { // сканим в Config
		panic(err)
	}
	fmt.Println(cfg.Host)
	fmt.Println(cfg.Port)
}

func main() {
	Env_library()
}

// 2024/10/01 19:56:57 set Host with $APP_HOST="127.0.0.1"
// 2024/10/01 19:56:57 set Port with $APP_PORT="8080"
// 127.0.0.1
// 8080
