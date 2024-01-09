package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func loadFileFromPath(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return []byte{}, err
	}

	return bs, nil
}

func main() {
	// Get path from command line args.
	if len(os.Args) < 2 {
		fmt.Print("Usage : ./cli <path_to_file>")
	}

	path := os.Args[1]

	buf, err := loadFileFromPath(path)
	if err != nil {
		panic(err)
	}
	meta, err := GetSGXMetadata(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(meta)

	return
}
