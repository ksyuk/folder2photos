package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"photos/photos"
	"photos/router"
)

func receiveInput(reader *bufio.Reader, message string) string {
	fmt.Printf("%s:", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error receiving input:", err)
	}
	input = strings.TrimSpace(input)

	return input
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("upload([1]) / start router(2)")
	mode := receiveInput(reader, "mode")

	if mode == "2" {
		router.Start()
	} else {
		folderName := receiveInput(reader, "folder name")
		albumName := receiveInput(reader, "album name")

		accessToken := os.Getenv("ACCESS_TOKEN")

		album := photos.AlbumFactory{AccessToken: accessToken, AlbumName: albumName}
		albumId := album.CreateAlbum()

		item := &photos.ItemFactory{
			AccessToken:   accessToken,
			DirectoryName: folderName,
			AlbumId:       albumId,
		}
		item.Upload()
	}
}
