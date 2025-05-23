package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"
)

func main() {
	fmt.Println("=============Processing file infos in directory===============")
	directory_name := flag.String("fileName", "./DIR_util", "a string")
	flag.Parse()
	dir_path := *directory_name
	file_infos := Read_directory_files_folders(dir_path)
	fmt.Println(file_infos)
	files := Process_all_file_infos(file_infos)
	log.Println("Listing files")
	files_as_string := Process_all_file_infos_string(files)
	fmt.Println(files_as_string)
	fmt.Println("=============Copy file source to destination=================")
	source_path := "./go.mod"
	dest_path := "./DIR_util/go.mod"
	Copy_file_source_destination(source_path, dest_path)
	Remove_created_files(dest_path)
	fmt.Println("===============================================================")

}




type MyFileInfo struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	IsDir bool   `json:"isdir"`
}

func (fl MyFileInfo) get_file_details_json() string {
	message := ""
	log.Println("Converting file info structure to json.")
	data, err := json.Marshal(fl)
	if err != nil {
		message = fmt.Sprintf("Error occured %s", err)
		log.Fatal(message)
		os.Exit(1)
	}
	log.Println("Returning file data to json string.")
	return string(data)
}

func process_all_file_infos(infos []MyFileInfo) []string {
	myfile_data := []string{}
	log.Println("Processing all the file infos.")
	for _, file_obj := range infos {
		myfile_data = append(myfile_data, file_obj.get_file_details_json())
	}
	return myfile_data
}

func Read_directory_files_folders(directory_path string) []os.FileInfo {
	message := ""
	log.Println("Collecting all the file infos.")
	fileInfos, err := ioutil.ReadDir(directory_path)
	if err != nil {
		message = fmt.Sprintf("Error occured %s", err)
		log.Println(message)
		os.Exit(1)
	}
	log.Println("Returning list of file infos.")
	return fileInfos
}

func Process_all_file_infos(file_infos []os.FileInfo) []MyFileInfo {
	infos := []MyFileInfo{}
	log.Println("Processing the file infos for the files obtained from directory.")
	for _, file_info := range file_infos {
		myfile_info := MyFileInfo{Name: file_info.Name(), Size: file_info.Size(), IsDir: file_info.IsDir()}
		infos = append(infos, myfile_info)
	}
	return infos
}

func Process_all_file_infos_string(infos []MyFileInfo) []string {
	return process_all_file_infos(infos)
}


// Copy_file_source_destination copies a file from the source path to the destination path
func Copy_file_source_destination(sourcePath, destPath string) {
    sourceFile, err := os.Open(sourcePath)
    if err != nil {
        log.Fatal(err)
    }
    defer sourceFile.Close()

    destFile, err := os.Create(destPath)
    if err != nil {
        log.Fatal(err)
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, sourceFile)
    if err != nil {
        log.Fatal(err)
    }
}

// Remove_created_files removes the specified file
func Remove_created_files(filePath string) {
    err := os.Remove(filePath)
    if err != nil {
        log.Fatal(err)
    }
}

// =============Processing file infos in directory===============
// 2025/01/19 14:25:15 Collecting all the file infos.
// 2025/01/19 14:25:15 Error occured open ./DIR_util: no such file or directory
// exit status 1
