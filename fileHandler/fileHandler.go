package fileHandler

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile(fileName string) {
	if VerifyIfFileExists(fileName) {
		fmt.Println("File already exists!")
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println(fileName + " created successfully!")
}

func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileName + " removed successfully!")
}

func WriteFile(fileName string, texts []string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, line := range texts {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Changes done successfully in " + fileName)
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func ReadFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileSize, err := GetFileSize(fileName)
	if err != nil {
		panic(err)
	}

	if fileSize == 0 {
		fmt.Println(fileName + " is empty!")
		return
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	fmt.Println("Reading " + fileName + "...\n")
	for {
		n, err := reader.Read(buffer)

		if err != nil {
			fmt.Println("\nThere's nothing more to read!")
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}

func VerifyIfFileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func DeleteAllLines(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("All lines removed successfully in " + fileName)
}

func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
