package File_Ops

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Copy_file_source_destination(source_path string, dest_path string) {
	message := ""
	source_file, err := os.Open(source_path)
	if err != nil {
		message = fmt.Sprintf("Unable to open file in path %s", source_path)
		log.Fatal(message)
		os.Exit(1)
	}
	defer source_file.Close()
	dest_file, err := os.Create(dest_path)
	if err != nil {
		message = fmt.Sprintf("Unable to create file in path %s", dest_path)
		log.Fatal(message)
		os.Exit(1)
	}
	defer dest_file.Close()
	bytes_copied, err := io.Copy(dest_file, source_file)
	if err != nil {
		message = fmt.Sprintf("Error occured %s", err)
		log.Fatal(message)
		os.Exit(1)
	}
	message = fmt.Sprintf("Copied %d bytes", bytes_copied)
	log.Println(message)
	return

}

func Remove_created_files(file_path string) {
	message := ""
	err := os.Remove(file_path)
	if err != nil {
		message = fmt.Sprintf("Unable to remove the file %s", file_path)
		log.Fatal(message)
		os.Exit(1)
	}
	return
}
