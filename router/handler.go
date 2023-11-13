package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type config struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
}

type oauth2Body struct {
	config
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

func loadEnv() config {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	return config{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectUri:  os.Getenv("REDIRECT_URI"),
		ResponseType: os.Getenv("RESPONSE_TYPE"),
		Scope:        os.Getenv("SCOPE"),
	}
}

func createOauth2Url() string {
	cfg := loadEnv()

	baseURL := "https://accounts.google.com/o/oauth2/v2/auth"
	params := url.Values{}
	params.Add("client_id", cfg.ClientId)
	params.Add("redirect_uri", cfg.RedirectUri)
	params.Add("response_type", cfg.ResponseType)
	params.Add("scope", cfg.Scope)
	oauth2URL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	return oauth2URL
}

func createOauth2Body(code string) *bytes.Buffer {
	cfg := loadEnv()

	body := oauth2Body{
		config:    cfg,
		Code:      code,
		GrantType: "authorization_code",
	}
	encodedReqBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(encodedReqBody)
}

func Oauth2(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, createOauth2Url())
}

func Oauth2Callback(c *gin.Context) {
	req, err := http.NewRequest(
		"POST", "https://oauth2.googleapis.com/token", createOauth2Body(c.Query("code")))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	encodedResBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	c.String(200, string(encodedResBody))
}
