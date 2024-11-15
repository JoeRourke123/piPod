package additional_info

import (
	sptfy "github.com/zmb3/spotify/v2"
	"orchestrator/service/spotify"
	"orchestrator/ui"
	"orchestrator/ui/model"
)

func AlbumAdditionalInfo(album sptfy.SavedAlbum) []model.AdditionalInfo {
	return []model.AdditionalInfo{
		{
			Text: spotify.GetArtistString(album.Artists),
			Icon: ui.USER_SOUND,
			Bold: true,
		},
		{
			Text: spotify.GetAlbumLength(album.Tracks),
			Icon: ui.HOURGLASS_MEDIUM,
			Bold: false,
		},
		{
			Text: album.ReleaseDate,
			Icon: ui.CALENDAR,
			Bold: false,
		},
	}
}
