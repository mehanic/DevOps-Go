package main

import (
	"context"
	_ "fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"os"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv,
		client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	// Pull the latest image
	out, err := cli.ImagePull(ctx, "nginx:latest", ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
	// Create the container
	resp, err := cli.ContainerCreate(ctx, &container.Config{Image: "nginx:latest"}, &container.HostConfig{
		Binds: []string{"/path/on/host:/path/on/container"},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, ContainerStartOptions{}); err != nil {
		panic(err)
	}
	// Attach to the container logs
	out, err = cli.ContainerLogs(ctx, resp.ID, ContainerLogsOptions{ShowStdout: true, Follow: true})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

