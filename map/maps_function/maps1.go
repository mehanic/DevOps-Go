package main

import (
	"fmt"
	"strings"
)

func buildConnectionString(params map[string]string) string {
	parts := []string{}
	for k, v := range params {
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(parts, ";")
}

func main() {
	myParams := map[string]string{
		"server":   "mpilgrim",
		"database": "master",
		"uid":      "sa",
		"pwd":      "secret",
	}

	fmt.Println(buildConnectionString(myParams))
}

// pwd=secret;server=mpilgrim;database=master;uid=sa
