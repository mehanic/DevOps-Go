package main

import (
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: uncomp <file>")
		return
	}

	file := os.Args[1]
	switch {
	case filepath.Ext(file) == ".tgz" || filepath.Ext(file) == ".tar.gz":
		if err := untarGz(file); err != nil {
			fmt.Println("Error:", err)
		}
	case filepath.Ext(file) == ".tar.Z":
		if err := uncompressTarZ(file); err != nil {
			fmt.Println("Error:", err)
		}
	case filepath.Ext(file) == ".tar.bz2":
		if err := untarBz2(file); err != nil {
			fmt.Println("Error:", err)
		}
	case filepath.Ext(file) == ".tar":
		if err := untar(file); err != nil {
			fmt.Println("Error:", err)
		}
	case filepath.Ext(file) == ".gz":
		if err := gunzip(file); err != nil {
			fmt.Println("Error:", err)
		}
	case filepath.Ext(file) == ".Z":
		if err := gunzip(file); err != nil {
			fmt.Println("Error:", err)
		}
	case filepath.Ext(file) == ".bz2":
		if err := uncompressBz2(file); err != nil {
			fmt.Println("Error:", err)
		}
	case filepath.Ext(file) == ".zip":
		if err := unzip(file); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unsupported file extension:", file)
	}
}

// Function to untar gzipped files
func untarGz(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	return untar(gzr)
}

// Function to untar bzipped files
func untarBz2(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	bzr := bzip2.NewReader(f)
	return untar(bzr)
}

// Function to untar files
func untar(r io.Reader) error {
	tarReader := tar.NewReader(r)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of tar archive
		}
		if err != nil {
			return err
		}

		// Create the file/directory based on the header info
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(header.Name, os.FileMode(0755)); err != nil {
				return err
			}
		case tar.TypeReg:
			f, err := os.OpenFile(header.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tarReader); err != nil {
				return err
			}
			f.Close()
		}
	}
	return nil
}

// Function to gunzip files
func gunzip(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	// Create the output file
	outFile := file[:len(file)-3] // Remove .gz
	out, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write uncompressed data to the output file
	_, err = io.Copy(out, gzr)
	return err
}

// Function to uncompress .bz2 files using system command
func uncompressBz2(file string) error {
	cmd := exec.Command("bunzip2", file)
	return cmd.Run()
}

// Function to unzip .zip files
func unzip(file string) error {
	r, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Extracting: %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		outFile, err := os.Create(f.Name)
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

// Function to uncompress .tar.Z files using system command
func uncompressTarZ(file string) error {
	cmd := exec.Command("gunzip", file)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Extract the .tar file that was created
	return untar(file[:len(file)-2]) // Remove .Z
}

