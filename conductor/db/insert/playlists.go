package insert

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/fetch"
	"conductor/db/serializer"
	"conductor/db/update"
	"errors"
)

func Playlist(playlists ...*model.Playlist) error {
	errMessage := ""
	if playlists != nil {
		for _, playlist := range playlists {
			existingPlaylist := fetch.Playlist(playlist.Id)
			if existingPlaylist != nil {
				err := update.Playlists.Update(playlist, true, true)
				if err != nil {
					errMessage += err.Error() + "; "
				}
			} else {
				_, err := db.X.InsertOne(db.PlaylistCollection, serializer.PlaylistSerializer.Serialize(playlist))
				if err != nil {
					errMessage += err.Error() + "; "
				}
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
