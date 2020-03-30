package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	_ "github.com/go-sql-driver/mysql"
)

const (
	EnvBucketName = "BUCKET_NAME"
	EnvObjectName = "OBJECT_NAME"

	EnvHost     = "DB_HOST"
	EnvUser     = "DB_USER"
	EnvPassword = "DB_PASSWORD"
	EnvDatabase = "DB_NAME"
)

var (
	sessionID = "backup"
)

func main() {
	sqlHC()
	ctx := context.Background()
	loadFile(ctx)
}

func sqlHC() {
	for i := 0; i < 15; i++ {
		if _, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", getEnv(EnvUser), getEnv(EnvPassword), getEnv(EnvHost), getEnv(EnvDatabase))); err == nil {
			log.Printf("Waiting for %s:3306", getEnv(EnvHost))
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

}

func loadFile(ctx context.Context) {
	log.Print("Load file")
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Creates a Bucket instance.
	bucket := client.Bucket(getEnv(EnvBucketName))
	rdr, err := bucket.Object(getEnv(EnvObjectName)).NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to read object: %v", err)
	}
	defer rdr.Close()

	f, err := os.Create(getFile())
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(rdr)
	for scanner.Scan() {
		b := scanner.Bytes()
		b = append(b, '\n')
		if _, err := f.Write(b); err != nil {
			log.Fatalf("Failed to read object with scanner: %v", err)
		}
	}
}

func getFile() string {
	return fmt.Sprintf("/tmp/%s.sql", sessionID)
}

func getEnv(name string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}
	log.Fatalf("%s not set as environment variable", name)
	return ""
}
