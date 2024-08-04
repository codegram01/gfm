package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0) // hide timestamp
	// to read arg by flag, we must parse first
	flag.Parse()

	err := execCommand()
	if err != nil {
		log.Fatal(err)
	}
}

func execCommand() error {
	command := getArg(0, "Enter command")
	switch command {
	case "ls": // list file
		files, err := os.ReadDir(".")
		if err != nil {
			return err
		}
		for _, file := range files {
			fmt.Printf("%s  ", file.Name())
		}
		fmt.Println()
	case "rm": // delete file
		path := getArg(1, "Enter path")
		return os.RemoveAll(path)
	case "pwd": // current path
		currentPath, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(currentPath)
	case "mv": // move file 
		path1 := getArg(1, "Enter path 1")
		path2 := getArg(2, "Enter path 2")

		return os.Rename(path1, path2)
	case "mkdir": // create dir
		path := getArg(1, "Enter Dir")
		return os.MkdirAll(path, 0777)
	case "cat":
		path := getArg(1, "Enter path")
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		os.Stdout.Write(data)
	case "touch": // create empty file
		path := getArg(1, "Enter path")
		
		_, err := os.Create(path)
		return err
	case "cp": // copy file
		// path old file
		path1 := getArg(1, "Enter path 1")
		// path new file
		path2 := getArg(2, "Enter path 2")

		data, err := os.ReadFile(path1)
		if err != nil {
			return err
		}
		return os.WriteFile(path2, data, 0777)
	case "echo": // add content to file
		content := getArg(1, "Enter content")
		path := getArg(2, "Enter path")
		
		return os.WriteFile(path, []byte(content), 0777)
	default:
		log.Fatal("Unknow command")
	}
	return nil
}

// this function must get arg
// if user not enter arg , it will break and show error
func getArg(indexArg int, mess string) string {
	val := flag.Arg(indexArg)
	if val == "" {
		log.Fatal(mess)
	}

	return val
}
