package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func main() {
	os.Setenv("STORAGE_EMULATOR_HOST", "127.0.0.1:9199")
	defer os.Unsetenv("STORAGE_EMULATOR_HOST")

	c, err := storage.NewClient(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	const bucketName = "somebucket"

	b := c.Bucket(bucketName)

	const fileName = "filename123"

	// put object with metadata
	fw := b.Object(fileName).NewWriter(context.Background())
	fw.Metadata = map[string]string{"a": "b"}
	fw.Write([]byte("123"))
	err = fw.Close()
	if err != nil {
		log.Fatalln("failed to push file data to storage:", err)
	}

	file := b.Object(fileName)

	// to read file data
	fr, err := file.NewReader(context.Background())
	if err != nil {
		log.Fatalln("failed to create reader for file:", err)
	}
	fileData, err := ioutil.ReadAll(fr)
	if err != nil {
		log.Fatalln("failed to read the file:", err)
	}
	fmt.Printf("file data: %q\n", fileData)

	// to read file metadata
	atrrs, err := file.Attrs(context.Background())
	if err != nil {
		log.Fatalln("failed to get the attrs:", err)
	}
	fmt.Printf("file metadata: %#v\n", atrrs.Metadata)
}
