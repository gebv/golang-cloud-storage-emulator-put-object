package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

var (
	c *storage.Client
)

var projectID = os.Getenv("GCP_PROJECT")

func main() {
	// asserts
	if os.Getenv("STORAGE_EMULATOR_HOST") == "" && os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		panic("env STORAGE_EMULATOR_HOST is required")
	}

	fmt.Println("Assert envs")

	var err error
	c, err = storage.NewClient(context.Background())

	if err != nil {
		log.Fatalln("failed to connect the cloud storage", err)
	}

	fmt.Println("Connected")

	const bucketName = "somebucket"

	if os.Getenv("STORAGE_EMULATOR_SKIP_CREATE_BUCKET") == "" {
		createBucketIfNeed(bucketName)
	} else {
		fmt.Println("Skip the creation of the bucket")
	}

	fmt.Println("Assert bucket")

	const fileName = "path/to/file"

	// put object with metadata
	fw := c.Bucket(bucketName).Object(fileName).NewWriter(context.Background())
	fw.Metadata = map[string]string{"a": "b"}
	fw.Write([]byte("123"))
	err = fw.Close()
	if err != nil {
		log.Fatalln("failed to push file data to storage:", err)
	}

	fmt.Println("Object is created")

	// to read file data
	fr, err := c.Bucket(bucketName).Object(fileName).NewReader(context.Background())
	if err != nil {
		log.Fatalln("failed to create reader for file:", err)
	}
	fileData, err := ioutil.ReadAll(fr)
	if err != nil {
		log.Fatalln("failed to read the file:", err)
	}
	fmt.Printf("file data: %q\n", fileData)

	// to read file metadata
	atrrs, err := c.Bucket(bucketName).Object(fileName).Attrs(context.Background())
	if err != nil {
		log.Fatalln("failed to get the attrs:", err)
	}
	fmt.Printf("file metadata: %#v\n", atrrs.Metadata)
}

func createBucketIfNeed(name string) {
	b := c.Bucket(name)
	_, err := b.Attrs(context.Background())
	if err == nil {
		return
	}
	if err == storage.ErrBucketNotExist {
		err = b.Create(context.TODO(), projectID, &storage.BucketAttrs{
			Name: name,
		})
		if err != nil {
			log.Fatalln("failed to create bucket:", err)
		}
	} else if err != nil {
		log.Fatalln("failed to get attrs of bucket", err)
	}
}
