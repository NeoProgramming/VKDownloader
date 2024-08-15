package core

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
	"vkdownloader/cfmt"
)

type RecordAttrs int

func InitVK() {
	cfmt.PrintlnFunc("VK API library initializing...")
	App.vk = api.NewVK(App.config.AccessToken)
	cfmt.PrintlnLine("VK API library initialized")
}

func extractAccessToken(urlStr string) string {
	u, _ := url.Parse(urlStr)
	parameters, _ := url.ParseQuery(u.Fragment)
	accessToken := parameters.Get("access_token")
	return accessToken
}

func LargestPhoto(photo *object.PhotosPhoto) string {
	maxUrl := ""
	maxSize := 0.0
	cfmt.PrintlnFunc("LargestPhoto")
	for _, s := range photo.Sizes {
		size := s.Width * s.Height
		cfmt.PrintlnLine("file=", s.URL, " size=", size, " w=", s.Width, " h=", s.Height)
		if size > maxSize {
			maxSize = size
			maxUrl = s.URL
		}
	}
	cfmt.PrintlnFunc("LargestPhoto end: file=", maxUrl, " size=", maxSize)
	return maxUrl
}

type Photo struct {
	url string
	id  int
}

type Album struct {
	id    int
	title string
}

func (app *Application) getAlbumsList(oid string) []Album {

	result, err := app.vk.PhotosGetAlbums(api.Params{
		"owner_id": oid,
	})
	if err != nil {
		fmt.Println("Error fetching albums:", err)
		return nil
	}

	// Print the result
	var albums []Album
	for _, album := range result.Items {
		fmt.Printf("Album ID: %d, Title: %s\n", album.ID, album.Title)
		a := Album{album.ID, album.Title}
		albums = append(albums, a)
	}
	return albums
}

func (app *Application) getPhotosList(oid string, aid string) []Photo {
	cfmt.PrintlnFunc("loadPhotos: oid=", oid, " aid=", aid)
	offset := 0
	totalCount := 0
	app.totalItems = 0
	app.currentItem = 0
	var imageLinks []Photo

	for {
		tStart := time.Now()

		cfmt.PrintlnLine("Request for photos: offset=", offset)

		photos, err := app.vk.PhotosGet(api.Params{
			"owner_id": oid,
			"album_id": aid,
			"offset":   offset,
			"count":    100, // 1000 is too big?,
		})
		if err != nil {
			cfmt.PrintlnErr("loadPhotos:", err)
			return nil
		}
		received := len(photos.Items)
		totalCount += received
		offset += received
		app.totalItems = photos.Count
		app.currentItem += received
		cfmt.PrintlnLine("Received: ", received, " RTotal: ", photos.Count, "CTotal: ", totalCount)

		// adding downloaded photos to the database
		for _, photo := range photos.Items {
			maxUrl := LargestPhoto(&photo)
			image := Photo{maxUrl, photo.ID}
			imageLinks = append(imageLinks, image)

			cfmt.PrintlnLine("URL: ", maxUrl)
		}

		// if the received number of elements is less than the number in the package, then the package is the last
		if totalCount >= photos.Count {
			break
		}

		tElapsed := time.Since(tStart)
		if tElapsed < 350*time.Millisecond {
			time.Sleep(350*time.Millisecond - tElapsed)
		}
	}
	return imageLinks
}

func downloadPhoto(url string, filename string) error {
	// Create a new file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed creating file: %w", err)
	}
	defer file.Close()

	// Send a GET request to the photo URL
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed getting photo: %w", err)
	}
	defer response.Body.Close()

	// Copy the body content to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("failed copying photo to file: %w", err)
	}

	return nil
}
