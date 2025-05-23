package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New() //другой веб фреймворк, для упрощения примеров

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {

		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

//
//┌─────(free)─────(~/GOLANG/Slurm_course/golang-main/8.1)
//└> $ go run main.go

//____    __
/// __/___/ /  ___
/// _// __/ _ \/ _ \
///___/\__/_//_/\___/ v4.6.1
//High performance, minimalist Go web framework
//https://echo.labstack.com
//____________________________________O/_______
//O\
//⇨ http server started on [::]:8080
//^[[{"time":"2024-03-18T02:55:59.048529073+01:00","id":"","remote_ip":"127.0.0.1","host":"127.0.0.1:8080","method":"GET","uri":"/metrics","user_agent":"Prometheus/2.28.0","status":404,"error":"code=404, message=Not Found","latency":106939,"latency_human":"106.939µs","bytes_in":0,"bytes_
//out":24}                                                                                                                                    {"time":"2024-03-18T02:55:59.551391758+01:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8080","method":"GET","uri":"/","user_agent":"Blackbox Exporter/0.22.0","status":200,"error":"","latency":103129,"latency_human":"103.129µs","bytes_in":0,"bytes_out":17}
//{"time":"2024-03-18T02:56:14.048108108+01:00","id":"","remote_ip":"127.0.0.1","host":"127.0.0.1:8080","method":"GET","uri":"/metrics","user_agent":"Prometheus/2.28.0","status":404,"error":"code=404, message=Not Found","latency":38728,"latency_human":"38.728µs","bytes_in":0,"bytes_out":24}
//{"time":"2024-03-18T02:56:14.550163178+01:00","id":"","remote_ip":"127.0.0.1","host":"localhost:8080","method":"GET","uri":"/","user_agent":"Blackbox Exporter/0.22.0","status":200,"error":"","latency":24661,"latency_human":"24.661µs","bytes_in":0,"bytes_out":17}
//
