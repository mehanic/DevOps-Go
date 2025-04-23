package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "strings"

    "github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
    // Get all the PDF filenames in the current directory
    files, err := os.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    var pdfFiles []string

    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".pdf") {
            pdfFiles = append(pdfFiles, file.Name())
        }
    }

    // Sort the filenames alphabetically
    sort.Strings(pdfFiles)

    // Combine all the PDFs except the first page of each
    outputFile := "allminutes.pdf"
    err = combinePDFs(pdfFiles, outputFile)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("PDFs combined successfully into %s\n", outputFile)
}

// combinePDFs merges all PDF files skipping the first page of each
func combinePDFs(pdfFiles []string, outputFile string) error {
    var allPages []string

    // Extract pages except the first from each PDF
    for _, file := range pdfFiles {
        // Define temp output for each single PDF minus first page
        tmpOutput := fmt.Sprintf("tmp_%s", file)

        // Extract the pages (excluding the first one)
        err := api.ExtractPagesFile(file, tmpOutput, []string{"2-end"}, nil)
        if err != nil {
            return fmt.Errorf("failed to extract pages from %s: %w", file, err)
        }

        // Check if the temporary file has content
        info, err := os.Stat(tmpOutput)
        if err != nil || info.Size() == 0 {
            // If the file is empty, skip adding it to the allPages slice
            os.Remove(tmpOutput) // Clean up immediately
            continue
        }

        allPages = append(allPages, tmpOutput)
    }

    // Check if we have any pages to merge
    if len(allPages) == 0 {
        return fmt.Errorf("no pages to merge")
    }

    // Merge the resulting pages into one file
    err := api.MergeCreateFile(allPages, outputFile, false, nil)
    if err != nil {
        return fmt.Errorf("failed to merge PDF files: %w", err)
    }

    // Clean up temporary files
    for _, tmpFile := range allPages {
        os.Remove(tmpFile)
    }

    return nil
}

// 2025/01/21 14:11:58 failed to extract pages from allminutes.pdf: The file could not be opened because it is empty.
// exit status 1
