package photos

type ItemFactory struct {
	AccessToken   string
	DirectoryName string
	AlbumId       string
}

type AlbumFactory struct {
	AccessToken string
	AlbumName   string
}

type MediaInfo struct {
	DirectoryName string
	UploadToken   string
}

type SimpleMediaItem struct {
	FileName    string `json:"fileName"`
	UploadToken string `json:"uploadToken"`
}

type NewMediaItem struct {
	Description     string          `json:"description"`
	SimpleMediaItem SimpleMediaItem `json:"simpleMediaItem"`
}

type batchReqBody struct {
	AlbumId       string         `json:"albumId"`
	NewMediaItems []NewMediaItem `json:"newMediaItems"`
}

type AlbumTitle struct {
	Title string `json:"title"`
}

type NewAlbum struct {
	Album AlbumTitle `json:"album"`
}

type Album struct {
	Id                    string `json:"id"`
	Title                 string `json:"title"`
	ProductURL            string `json:"productURL"`
	MediaItemCount        string `json:"mediaItemCount"`
	CoverPhotoBaseUrl     string `json:"coverPhotoBaseUrl"`
	CoverPhotoMediaItemId string `json:"coverPhotoMediaItemId"`
}

type Albums struct {
	Albums        []Album `json:"albums"`
	NextPageToken string  `json:"nextPageToken"`
}
