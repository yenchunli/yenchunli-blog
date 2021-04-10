---
title: "Write a imgur upload service with go"
date: 2021-04-10T21:43:38+08:00
draft: false
tags: ["go", "imgur", "gin"]
---

## Background

Most of the services need a image upload services. Some applications like Ptt, Dcard and HackMD use Imgur to store pictures. Imgur provides free spaces for users to upload their images, and it also provides API so that programers can upload images by sending POST request. In this tutorial, we are going to build a imgur upload service.

## Packages Used

- gonic/gin: http server
- bytes: store images
- encoding/json: process response
- io: transfer io.Reader to []byte
- "mime/multipart": form data(request)

## Code

### 1. Create a Client struct and constructor to store/initialize token and imgur API URL

```go
type Client struct {
	token string
	uploadApiUrl string
}

func NewClient(token string, uploadApiUrl string) *Client{
	return &Client{token: token, uploadApiUrl: uploadApiUrl}
}
```

### 2. Create a UploadImage function

We consider to choose type `[]byte` instead of `io.Reader` for params.
The main reason is that we should pass the real data, not the FileReader pointer.
Also `[]byte` is easier to write test.

#### Workflows

1. if image is nil, return error
2. Create a multipart writer for writing form-data
3. CreateFormFile creates fields for form
4. Copy data at the end of the writer
5. Send a POST request to Imugr with the authorization token
6. Decode the response
7. Return the imgur image url

#### Functions

|Package| Function |
|-|-|
|mime/multipart |func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)   |
|io             |func Copy(dst Writer, src Reader) (written int64, err error)                     |
|http           |func NewRequest(method, url string, body io.Reader) (*Request, error)            |
|io/ioutil      |func ReadAll(r io.Reader) ([]byte, error)                                        |

#### Code

```go
func (client *Client) UploadImage(image []byte) (string, error){
    if image == nil {
		return "", errors.New("No Image")
	}
	var buf = new(bytes.Buffer)
    writer := multipart.NewWriter(buf)

    part, _ := writer.CreateFormFile("image", "filename")

	imgReader := bytes.NewReader(image)
    io.Copy(part, imgReader)

    writer.Close()
    req, _ := http.NewRequest("POST", client.uploadApiUrl, buf)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Authorization", "Bearer "+client.token)

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

	dec := json.NewDecoder(bytes.NewReader(body))
	var img imageInfoDataWrapper
	if err := dec.Decode(&img); err != nil {
		return "", errors.New("Fail to decode")
	}

	if !img.Success {
		return "", errors.New("Fail")
	}
    
	return img.Data.Link, nil
}
```

## Tests

We set token by environment variable.

There are two tests:

1. Test upload nil image
2. Test upload real image

`io.ReadFile()` reads file in []byte as input file.

Need to check if the return URL is correct (https://i.imgur.com/xxxxx).

```go
package upload

import (
	"testing"
	"os"
	"io/ioutil"
	"regexp"
)

// Test if upload a nil image, it should return error
func TestUploadNilImage(t *testing.T) {

	client := NewClient("", "")

	_, err := client.UploadImage(nil)

	if err == nil {
		t.Error("UploadImage() should have an error")
		t.Fail()
	}
}

// Test if upload real image, it should return success(200)
func TestUploadRealImage(t *testing.T) {
	token := os.Getenv("IMGUR_UPLOAD_TOKEN")
	if token == "" {
		t.Skip("IMGUR_UPLOAD_TOKEN is not set.")
	}

	client := NewClient(os.Getenv("IMGUR_UPLOAD_TOKEN"), "https://api.imgur.com/3/upload")

	// Read File to byte
	file, err := ioutil.ReadFile("logo.png")
	if err != nil {
		t.Skip("Can't read logo.png for test")
	}

	url, err := client.UploadImage(file)
	if err != nil {
		t.Errorf("UploadImage() failed with error: %v", err)
		t.Fail()
	}
	if matched, _ := regexp.MatchString(`https://i.imgur.com/`, url); !matched {
		t.Error("UploadImage() did not return imgur url")
		t.Fail()
	}

}
```

## Usage

we need to pass `[]byte` to `UploadImage()`.

`c.FormData()`returns `*multipart.FileHeader`.

`*multipart.FileHeader.Open()` returns `io.Reader`

We use `ioutil.ReadAll()` to read `io.Reader` to `[]byte`

```go 
// imgurUploadServer.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yenchunli/arts-nthu-backend/pkg/upload"
	"io/ioutil"
	"os"
	"mime/multipart"
	"net/http"
)

func main() {

	r := gin.Default()
	r.POST("/api/v1/upload", func(c *gin.Context) {
		type request struct {
			image *multipart.FileHeader `form:image binding:"required"`
		}
		var req request
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "miss data",
			})
			return
		}


		file, err := c.FormFile("image") 	// *Multipart.FileHeader
		if file.Size <=0 {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		reader, err := file.Open()			// io.Reader
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		buf, err := ioutil.ReadAll(reader)	// bytes[]
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		client := upload.NewClient(os.Getenv("IMGUR_UPLOAD_TOKEN"), "https://api.imgur.com/3/upload")
		imgurUrl, _ := client.UploadImage(buf)


		c.JSON(200, gin.H{
			"url": imgurUrl,
		})
		return
	})
	r.Run() 
	
}
```