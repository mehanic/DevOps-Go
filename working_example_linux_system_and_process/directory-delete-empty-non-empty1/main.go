package main

import (
	"fmt"
	"os"
)

func delete_directory1(path string) error {
	errs := os.RemoveAll(path) //use the os package function to delete dirg
	if errs != nil {
		return fmt.Errorf("error deleting directory: %v", errs)
	}
	return nil
}

func main() {
	path := "/path/to/directory"

	errs := delete_directory1(path)
	if errs != nil {
		fmt.Println(errs) //print the error if directory is not deleted successfully
		os.Exit(1)
	}
	fmt.Println("Directory is deleted successfully") //print the success message if directory is deleted successfully
}
