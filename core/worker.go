package core

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"vkdownloader/cfmt"
)


func (app *Application) worker(album string) {
	
	// loop until the execution flag is cleared
	cfmt.PrintlnFunc("Worker running...")
	
	defer app.wg.Done()

	// https://vk.com/album-123456_123456789
	const prefix = "https://vk.com/album"
	if !strings.HasPrefix(album, prefix) {
		cfmt.PrintlnErr("bad album url")
		app.running = false
		return
	}
	album = strings.TrimPrefix(album, prefix)
	parts := strings.Split(album, "_")
	if len(parts) != 2 {
		cfmt.PrintlnErr("bad album url")
		app.running = false
		return
	}
	
	owner_id, _ := strconv.Atoi(parts[0])
	album_id, _ := strconv.Atoi(parts[1])
	
	fmt.Println("owner_id: ", parts[0])
	fmt.Println("album_id: ", parts[1])
	
	fileName := "album" + parts[0] + "_" + parts[1]
	albumPath := path.Join(app.config.SavePath, fileName)
	err := os.Mkdir(albumPath, 0755)
	if err!= nil && !os.IsExist(err) {
		fmt.Println("Error creating directory:", err)
		app.running = false
		return
	}
	
	imageLinks := app.getPhotosList(owner_id, album_id)
	fmt.Println("images count: ", len(imageLinks))
	for i, item := range imageLinks {
		fileName := "photo" + parts[0] + "_" + strconv.Itoa(item.id) + ".jpg"
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
