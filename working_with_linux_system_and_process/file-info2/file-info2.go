package main

import (
//	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileInfo stores file metadata.
type FileInfo struct {
	Name string
}

// MP3FileInfo extends FileInfo to store MP3 ID3v1.0 tags.
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

// stripNulls removes null bytes and trims the resulting string.
func stripNulls(data []byte) string {
	return strings.TrimSpace(strings.ReplaceAll(string(data), "\x00", " "))
}

// byteToString converts a byte to a string.
func byteToString(data []byte) string {
	return fmt.Sprintf("%d", data[0])
}

// Parse extracts ID3v1.0 tags from an MP3 file.
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

// listDirectory gets a list of file info objects for files with particular extensions.
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
				mp3.Name = path
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

