package photos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

func (item ItemFactory) Upload() {
	fmt.Println("Uploading...")
	mediaInfos := item.createMediaInfos()

	for _, v := range *mediaInfos {
		item.batchCreate(v)
	}

	fmt.Println("Done!")
}

func (item ItemFactory) batchCreate(mediainfo MediaInfo) {
	ENDPOINT := "https://photoslibrary.googleapis.com/v1/mediaItems:batchCreate"

	reqBody := batchReqBody{
		AlbumId: item.AlbumId,
		NewMediaItems: []NewMediaItem{
			{
				Description: "",
				SimpleMediaItem: SimpleMediaItem{
					FileName:    mediainfo.DirectoryName,
					UploadToken: mediainfo.UploadToken,
				},
			},
		},
	}
	encodedBody, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", ENDPOINT, bytes.NewReader(encodedBody))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+item.AccessToken)
	req.Header.Add("Content-type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func (item ItemFactory) uploadItem(photo io.Reader) string {
	ENDPOINT := "https://photoslibrary.googleapis.com/v1/uploads"

	req, _ := http.NewRequest("POST", ENDPOINT, photo)
	req.Header.Add("Authorization", "Bearer "+item.AccessToken)
	req.Header.Add("Content-type", "application/octet-stream")
	req.Header.Add("X-Goog-Upload-Content-Type", "image/jpeg")
	req.Header.Add("X-Goog-Upload-Protocol", "raw")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyBytes)
}

func sortByNumber(s []string) {
	sort.Slice(s, func(i, j int) bool {
		num1, _ := strconv.Atoi(strings.TrimSuffix(s[i], ".jpg"))
		num2, _ := strconv.Atoi(strings.TrimSuffix(s[j], ".jpg"))
		return num1 < num2
	})
}

func (item ItemFactory) createMediaInfos() *[]MediaInfo {
	directory, err := os.Open(item.DirectoryName)
	if err != nil {
		panic(err)
	}
	defer directory.Close()

	fileInfos, err := directory.Readdir(-1)
	if err != nil {
		panic(err)
	}

	files := make([]string, 0)
	for _, fileInfos := range fileInfos {
		files = append(files, fileInfos.Name())
	}

	sortByNumber(files)

	mediaInfos := make([]MediaInfo, 0)

	for i, file := range files {
		body, err := os.ReadFile(item.DirectoryName + "/" + file)
		if err != nil {
			panic(err)
		}

		bodyReader := bytes.NewReader(body)
		uploadToken := item.uploadItem(bodyReader)

		mediaInfos = append(mediaInfos, MediaInfo{
			UploadToken:   uploadToken,
			DirectoryName: item.DirectoryName + "+" + strconv.Itoa(i+1) + ".jpg",
		})
	}

	return &mediaInfos
}
