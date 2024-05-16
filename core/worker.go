package core

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"vkdownloader/cfmt"
)

func (app *Application) processDownloadOwner(owner_url string) {
	cfmt.PrintlnFunc("DownloadOwner running...")
	defer app.wg.Done()
	oid := parseOwnerUrl(owner_url)
	albums := app.getAlbumsList(oid)
	for _, album := range albums {
		app.processDownloadAlbum(oid, strconv.Itoa(album.id))
	}
}

func parseOwnerUrl(owner_url string) string {
	// https://vk.com/club12345
	const prefix1 = "https://vk.com/club"
	const prefix2 = "https://vk.com/public"
	const prefix3 = "https://vk.com/id"

	if strings.HasPrefix(owner_url, prefix1) {
		return "-" + strings.TrimPrefix(owner_url, prefix1)
	} else if strings.HasPrefix(owner_url, prefix2) {
		return "-" + strings.TrimPrefix(owner_url, prefix2)
	} else if strings.HasPrefix(owner_url, prefix3) {
		return strings.TrimPrefix(owner_url, prefix3)
	} else {
		cfmt.PrintlnErr("bad owner_url url")
	}
	return ""
}

func parseAlbumUrl(album_url string) []string {
	// https://vk.com/album-123456_123456789
	const prefix = "https://vk.com/album"
	if !strings.HasPrefix(album_url, prefix) {
		cfmt.PrintlnErr("bad album_url url")
		return nil
	}
	album_url = strings.TrimPrefix(album_url, prefix)
	parts := strings.Split(album_url, "_")
	if len(parts) != 2 {
		cfmt.PrintlnErr("bad album_url url")
		return nil
	}

	fmt.Println("owner_id: ", parts[0])
	fmt.Println("album_id: ", parts[1])
	return parts
}

func (app *Application) processDownloadAlbum(oid string, aid string) {

	// loop until the execution flag is cleared
	cfmt.PrintlnFunc("DownloadAlbum running...")

	defer app.wg.Done()

	fileName := "album" + oid + "_" + aid
	albumPath := path.Join(app.config.SavePath, fileName)
	err := os.Mkdir(albumPath, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating directory:", err)
		app.running = false
		return
	}

	imageLinks := app.getPhotosList(oid, aid)
	fmt.Println("images count: ", len(imageLinks))
	for i, item := range imageLinks {
		fileName := "photo" + oid + "_" + strconv.Itoa(item.id) + ".jpg"
		fullPath := path.Join(albumPath, fileName)
		fmt.Println(i, "-->", fullPath)

		downloadPhoto(item.url, fullPath)
		//	if i > 2 {
		//		break
		//	}
	}

	if app.running == true {
		cfmt.PrintlnImp("Worker stopped naturally")
	}
	app.running = false
}
