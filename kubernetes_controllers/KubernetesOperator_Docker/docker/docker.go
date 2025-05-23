package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	//"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/docker"
)

// these are set at build time
var (
	version   string
	builddate string
)

var versioninfo VersionInfo

func init() {
	// parse buildtime variables
	versioninfo.Version = version
	i, err := strconv.ParseInt(builddate, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	versioninfo.BuildDate = tm
}

func main() {
	http.HandleFunc("/version", VersionHandler(&versioninfo))
	fmt.Printf("version %s listening on :8000\n", versioninfo.Version)
	panic(http.ListenAndServe(":8000", nil))
}

//---

// VersionInfo holds artifacts passed in
// at build time
type VersionInfo struct {
	Version   string
	BuildDate time.Time
	Uptime    time.Duration
}

// VersionHandler writes the latest version info
func VersionHandler(v *VersionInfo) http.HandlerFunc {
	t := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		v.Uptime = time.Since(t)
		vers, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(vers)
	}
}
