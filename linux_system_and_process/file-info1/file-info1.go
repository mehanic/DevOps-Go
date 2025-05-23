package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Name string
}

func (fi *FileInfo) SetName(filename string) {
	fi.Name = filename
}

type MP3FileInfo struct {
	FileInfo
	Tags map[string]string
}

var tagDataMap = map[string]struct {
	start     int
	end       int
	parseFunc func([]byte) string
}{
	"title":   {3, 33, stripNulls},
	"artist":  {33, 63, stripNulls},
	"album":   {63, 93, stripNulls},
	"year":    {93, 97, stripNulls},
	"comment": {97, 126, stripNulls},
	"genre":   {127, 128, byteToString},
}

func stripNulls(data []byte) string {
	return strings.TrimSpace(string(bytes.ReplaceAll(data, []byte("\x00"), []byte(" "))))
}

func byteToString(data []byte) string {
	return fmt.Sprintf("%d", data[0])
}

func (mp3 *MP3FileInfo) Parse(filename string) {
	mp3.Tags = make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	// Seek to the last 128 bytes to get the ID3v1.0 tags
	file.Seek(-128, io.SeekEnd)
	tagdata := make([]byte, 128)
	_, err = file.Read(tagdata)
	if err != nil {
		return
	}

	// Check if the tag is present
	if string(tagdata[:3]) == "TAG" {
		for tag, bounds := range tagDataMap {
			start, end := bounds.start, bounds.end
			mp3.Tags[tag] = bounds.parseFunc(tagdata[start:end])
		}
	}
}

func listDirectory(directory string, fileExtList []string) []*MP3FileInfo {
	var files []*MP3FileInfo
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ext := strings.ToLower(filepath.Ext(path))
		for _, extMatch := range fileExtList {
			if ext == extMatch {
				mp3 := &MP3FileInfo{}
				mp3.SetName(path)
				mp3.Parse(path)
				files = append(files, mp3)
			}
		}
		return nil
	})
	return files
}

func main() {
	// Change to your directory with MP3 files
	directory := "/music/_singles/"
	files := listDirectory(directory, []string{".mp3"})

	for _, file := range files {
		fmt.Printf("File: %s\n", file.Name)
		for key, value := range file.Tags {
			fmt.Printf("%s: %s\n", key, value)
		}
		fmt.Println()
	}
}

