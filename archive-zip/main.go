package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Save to zip
	files := []string{"file1.txt", "file2.txt"}
	err := createZip("output.zip", files)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("ZIP file created successfully!")
	}

	// Read from zip
	err = extractZip("output.zip", "extracted_files")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("ZIP extracted successfully!")
	}

}

func createZip(zipFileName string, files []string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		err := addFileToZip(zipWriter, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	zipEntry, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(zipEntry, file)
	return err
}

func extractZip(zipFileName, destDir string) error {
	zipReader, err := zip.OpenReader(zipFileName)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		err := extractFile(file, destDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func extractFile(f *zip.File, destDir string) error {
	filePath := filepath.Join(destDir, f.Name)

	if f.FileInfo().IsDir() {
		return os.MkdirAll(filePath, os.ModePerm)
	}

	destFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	srcFile, err := f.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
