package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	dataFilePath   = "data.txt"
	resultFilePath = "result.txt"
)

func main() {
	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int
	var res strings.Builder
	connector := ","
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res.WriteString(line)
		res.WriteString(connector)
		count++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	resStr := strings.TrimRight(strings.TrimSpace(res.String()), connector)
	fmt.Println("line count:", count)

	err = WriteToFile(resultFilePath, resStr)
	if err != nil {
		panic(err)
	}
}

func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		fmt.Println("write succeed!")
		defer f.Close()
	}
	return err
}
