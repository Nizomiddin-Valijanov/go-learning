package main

import (
	"fmt"
	"os"
)

func main() {
	first_path, second_path := "text.txt", "new.txt"
	fmt.Println(readMyFile(first_path))
	fmt.Println(createWriteFile(second_path))
	deleteFile(second_path)
}

func readMyFile(filePath string) string {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File '%s' not found\n", filePath)
		} else {
			fmt.Println("Error reading file:", err)
		}
	}
	return string(fileData)
}

func createWriteFile(path string) string {
	data := []byte("New file and text")
	err := os.WriteFile(path, data, 0644)

	if err != nil {
		fmt.Println("Fayl yarata olmayman \n", err)
	}
	str := readMyFile(path)
	return str
}

func deleteFile(path string) {
	os.Remove(path)
}
