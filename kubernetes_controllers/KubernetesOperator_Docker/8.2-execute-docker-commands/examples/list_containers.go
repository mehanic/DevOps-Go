package examples

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"time"
)

func ListContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithTimeout(10*time.Second)) //собираем клиент с опциями из env переменных, в данном случае установим timeout
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{}) //забираем список контейнеров
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
		fmt.Println(container.Image)
		fmt.Println(container.Image)
	}
}

//
//┌─────(free)─────(~/GOLANG/Slurm_course/golang-main/8.2)
//└> $ go run main.go
//{"status":"Pulling from library/alpine","id":"latest"}
//{"status":"Pulling fs layer","progressDetail":{},"id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":34765,"total":3408729},"progress":"[\u003e                                                  ]  34.77kB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":143018,"total":3408729},"progress":"[==\u003e                                                ]    143kB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":310555,"total":3408729},"progress":"[====\u003e                                              ]  310.6kB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":605467,"total":3408729},"progress":"[========\u003e                                          ]  605.5kB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":888057,"total":3408729},"progress":"[=============\u003e                                     ]  888.1kB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":1125625,"total":3408729},"progress":"[================\u003e                                  ]  1.126MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":1371385,"total":3408729},"progress":"[====================\u003e                              ]  1.371MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":1617145,"total":3408729},"progress":"[=======================\u003e                           ]  1.617MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":1813753,"total":3408729},"progress":"[==========================\u003e                        ]  1.814MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":2059513,"total":3408729},"progress":"[==============================\u003e                    ]   2.06MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":2296110,"total":3408729},"progress":"[=================================\u003e                 ]  2.296MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":2541870,"total":3408729},"progress":"[=====================================\u003e             ]  2.542MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":2738478,"total":3408729},"progress":"[========================================\u003e          ]  2.738MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":2885934,"total":3408729},"progress":"[==========================================\u003e        ]  2.886MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":3082542,"total":3408729},"progress":"[=============================================\u003e     ]  3.083MB/3.409MB","id":"4abcf2066143"}
//{"status":"Downloading","progressDetail":{"current":3279150,"total":3408729},"progress":"[================================================\u003e  ]  3.279MB/3.409MB","id":"4abcf2066143"}
//{"status":"Download complete","progressDetail":{},"id":"4abcf2066143"}
//{"status":"Extracting","progressDetail":{"current":65536,"total":3408729},"progress":"[\u003e                                                  ]  65.54kB/3.409MB","id":"4abcf2066143"}
//{"status":"Extracting","progressDetail":{"current":3408729,"total":3408729},"progress":"[==================================================\u003e]  3.409MB/3.409MB","id":"4abcf2066143"}
//{"status":"Pull complete","progressDetail":{},"id":"4abcf2066143"}
//{"status":"Digest: sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b"}
//{"status":"Status: Downloaded newer image for alpine:latest"}
//
//┌─────(free)─────(~/GOLANG/Slurm_course/golang-main/8.2)
//└> $
//
//
