package main

import (
	"fmt"
	"os"
)

func delete_directory(path string) error {
	dr, err := os.Open(path) //open the path using os package function
	if err != nil {
		return err
	}
	defer dr.Close()
	names, err := dr.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(path + "/" + name) //remove the path
		if err != nil {
			return err
		}
	}
	return os.Remove(path)
}

func main() {
	path := "/path/to/directory"

	err := delete_directory(path)
	if err != nil {
		fmt.Println("Error deleting directory:", err)
		os.Exit(1)
	}
	fmt.Println("Directory is deleted successfully") //if directory is deleted print the success message
}

//-----------

func delete_directory1(path string) error {
	errs := os.RemoveAll(path) //use the os package function to delete dirg
	if errs != nil {
		return fmt.Errorf("Error deleting directory: %v", errs)
	}
	return nil
}

func main1() {
	path := "/path/to/directory"

	errs := delete_directory1(path)
	if errs != nil {
		fmt.Println(errs) //print the error if directory is not deleted successfully
		os.Exit(1)
	}
	fmt.Println("Directory is deleted successfully") //print the success message if directory is deleted successfully
}
