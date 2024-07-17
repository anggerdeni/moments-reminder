package main

import (
	"context"
	"fmt"
	"log"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("bismillah")
	driveService, err := drive.NewService(context.Background(), option.WithCredentialsFile("sa.json"))
	if err != nil {
		log.Fatal(err.Error())
	}

	// https://stackoverflow.com/questions/70749061/google-drive-api-error-message-shared-drive-not-found-xyz
	driveId := "1vmL0KpSfi1qDbPddwtxlpA1QIEmMkWq7"
	fileService := driveService.Files.List()
	fileService.Q(fmt.Sprintf("'%s' in parents", driveId))

	files, err := fileService.Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, f := range files.Files {
		fmt.Println(f.Name)
	}
}
