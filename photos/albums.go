package photos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (album AlbumFactory) CreateAlbum() string {
	ENDPOINT := "https://photoslibrary.googleapis.com/v1/albums"

	body := NewAlbum{
		Album: AlbumTitle{
			Title: album.AlbumName,
		},
	}
	encodedReqBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST", ENDPOINT, bytes.NewBuffer(encodedReqBody))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+album.AccessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	encodedResBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	newAlbum := Album{}
	json.Unmarshal(encodedResBody, &newAlbum)

	return newAlbum.Id
}

func (album AlbumFactory) ListAlbums() {
	ENDPOINT := "https://photoslibrary.googleapis.com/v1/albums"

	req, err := http.NewRequest("GET", ENDPOINT, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+album.AccessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	encodedResBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var albumList Albums
	json.Unmarshal(encodedResBody, &albumList)

	for _, v := range albumList.Albums {
		fmt.Println(v.Title)
		fmt.Println(v.Id)
	}
}

func (album AlbumFactory) GetAlbumInfo(albumId string) {
	ENDPOINT := fmt.Sprintf("https://photoslibrary.googleapis.com/v1/albums/%s", albumId)

	req, err := http.NewRequest("GET", ENDPOINT, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+album.AccessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	encodedResBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(encodedResBody))
}
