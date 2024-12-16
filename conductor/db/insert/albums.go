package insert

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/fetch"
	"conductor/db/serializer"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Album(albums ...*model.Album) error {
	errMessage := ""
	if albums != nil {
		for _, album := range albums {
			if fetch.AlbumExists(album.Id) {
				// Do not recreate album if it already exists
				continue
			}

			_, err := db.X.InsertOne(db.AlbumCollection, serializer.AlbumSerializer.Serialize(album))
			if err != nil {
				errMessage += err.Error() + "; "
			}
		}
	} else {
		errMessage += "album to insert is nil; "
	}

	if errMessage != "" {
		return errors.New("could not create album: " + errMessage)
	} else {
		return nil
	}
}

func AlbumArtwork(id string, url string) error {
	// Create the directory if it doesn't exist
	dir := ".db/artwork"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create directory: %v", err)
	}

	// Check if the file already exists
	filePath := filepath.Join(dir, id+".jpg")
	if _, err := os.Stat(filePath); err == nil {
		// File already exists, skip downloading
		return nil
	}

	// Fetch the image from the URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not fetch image: %v", err)
	}
	defer resp.Body.Close()

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	// Copy the image data to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("could not save image: %v", err)
	}

	return nil
}
