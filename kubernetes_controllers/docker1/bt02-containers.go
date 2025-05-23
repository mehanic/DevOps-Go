package main

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/olekukonko/tablewriter"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	data := [][]string{}

	for _, v := range containers {
		if v.Status[:2] == "Up" {
			v.Status = "Running"
		} else if v.Status[:6] == "Exited" {
			v.Status = "Stopped"
		}
		data = append(data, []string{
			v.ID[:12],
			v.Image,
			v.Status,
			Convert2String(v.Ports),
			strings.Trim(strings.Join(v.Names, ","), "/"),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"CONTAINER ID", "IMAGE", "STATUS", "PORTS", "NAMES"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func Convert2String(ports []types.Port) string {
	ps := []string{}
	for _, p := range ports {
		ps = append(ps, strconv.Itoa(int(p.PrivatePort))+":"+strconv.Itoa(int(p.PublicPort)))
	}
	return strings.Join(ps, ", ")
}
