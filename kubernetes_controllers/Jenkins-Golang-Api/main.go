package main

import (
	"fmt"
	jnkn "jenkins/JenkinsUtil"
	//"github.com/bndr/gojenkins"
)

func main() {
	jenkins, ctx := jnkn.GetContextJenkins("http://localhost:8080/", "bapan", "learner#12")
	job_names, err := jenkins.GetAllJobNames(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if len(job_names) == 0 {
		fmt.Println("No jobs currently.")
	}
	for index, value := range job_names {
		fmt.Println(index, value)
	}
	params := make(map[string]string)
	params["Name"] = "Baps"
	jnkn.TriggerBuild(ctx, "bapanproj", params, jenkins)
	// builds, err := jenkins.GetAllBuildIds(ctx, "bapanproj")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// for index, ids := range builds {
	// 	fmt.Println(index, ids)
	// }
}
