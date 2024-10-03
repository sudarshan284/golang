package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Welcome to the file system utility CLI tool!!!")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nEnter command (ls,mkdir,mv,exit)")
		command, _ := reader.ReadString('\n')
		command = trimNewline(command)
		switch command {
		case "ls":
			listFiles()
		case "mkdir":
			createFiles(reader)
		case "mv":
			moveFile(reader)
		case "exit":
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Invalid command,Please try again")
		}
	}
}

func listFiles() {
	files, err := filepath.Glob("*")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Files in the current directory:")
	for _, file := range files {
		fmt.Println("-", file)
	}
}

func createFiles(reader *bufio.Reader) {
	fmt.Print("Enter directory name")
	dirName, _ := reader.ReadString('\n')
	dirName = trimNewline(dirName)
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Directory created successfully: ", dirName)
}

func moveFile(reader *bufio.Reader) {
	fmt.Print("Enter source file name :")
	srcFilename, _ := reader.ReadString('\n')
	srcFilename = trimNewline(srcFilename)
	fmt.Print("Enter the destination directory : ")
	destDir, _ := reader.ReadString('\n')
	destDir = trimNewline(destDir)
	err := os.Rename(srcFilename, filepath.Join(destDir, srcFilename))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("File moved successfully.")
}

func trimNewline(s string) string {
	return s[:len(s)-1]
}
