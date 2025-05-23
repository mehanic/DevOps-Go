package main

import (
	"context"

	"github.com/PacktPublishing/Go-for-DevOps/chapter/7/cobra/app/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Execute(ctx)
}
