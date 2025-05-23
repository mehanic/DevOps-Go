package JenkinsUtil

import (
	"context"
	"fmt"
	"time"

	"github.com/bndr/gojenkins"
)

func GetContextJenkins(url, admin, password string) (*gojenkins.Jenkins, context.Context) {
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, url, admin, password)
	_, err := jenkins.Init(ctx)

	if err != nil {
		fmt.Println(err)
		panic("Unable to connect to remote jenkins server.")
	}
	return jenkins, ctx
}

func TriggerBuild(ctx context.Context, name string, params map[string]string, jenkins *gojenkins.Jenkins) {
	queueid, err := jenkins.BuildJob(ctx, name, params)
	if err != nil {
		panic(err)
	}
	build, err := jenkins.GetBuildFromQueueID(ctx, queueid)
	if err != nil {
		fmt.Println(err)
	}

	for build.IsRunning(ctx) {
		time.Sleep(5000 * time.Millisecond)
		build.Poll(ctx)
	}

	fmt.Printf("build number %d with result: %v\n", build.GetBuildNumber(), build.GetResult())
}
