package fetch

import (
	"conductor/common/model"
	"conductor/db"
	"github.com/ostafen/clover/v2/document"
	"golang.org/x/oauth2"
	"time"
)

func getConfig(key string) *document.Document {
	config, err := db.X.FindFirst(db.ConfigQuery(key))
	if err != nil {
		return nil
	}

	return config
}

func SpotifyToken() *oauth2.Token {
	doc := getConfig(db.SpotifyTokenKey)

	expiry, ok := doc.Get("Expiry").(time.Time)
	if !ok {
		expiryStr, ok2 := doc.Get("Expiry").(string)

		if ok2 {
			expiry, _ = time.Parse(time.RFC3339, expiryStr)
			ok = ok2
		}
	}

	if doc != nil && ok {
		token := oauth2.Token{
			AccessToken:  doc.Get("AccessToken").(string),
			TokenType:    doc.Get("TokenType").(string),
			Expiry:       expiry,
			RefreshToken: doc.Get("RefreshToken").(string),
		}

		return &token
	} else {
		return nil
	}
}

func OsState() *model.OsState {
	doc := getConfig(db.OsStateKey)

	if doc == nil {
		return nil
	}

	osState := model.OsState{
		IsInternetEnabled: doc.Get("IsInternetEnabled").(bool),
	}

	return &osState
}

func InternetEnabled() bool {
	osState := OsState()
	if osState == nil {
		return false
	} else {
		return osState.IsInternetEnabled
	}
}

func RssFeedCache(podcastId string) string {
	doc := getConfig("rssFeedCache")

	if doc == nil || doc.Has(podcastId) == false {
		return ""
	}

	return doc.Get(podcastId).(string)
}
