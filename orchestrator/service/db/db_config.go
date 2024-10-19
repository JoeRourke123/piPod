package db

import (
	"encoding/json"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"golang.org/x/oauth2"
	"log"
)

func GetSpotifyToken() *oauth2.Token {
	spotifyTokenConfig, err := db.FindFirst(spotifyTokenQuery())

	if err != nil {
		return nil
	}

	if spotifyTokenConfig == nil {
		return nil
	}

	var token *oauth2.Token
	tokenMap := spotifyTokenConfig.AsMap()
	tokenJson, _ := json.Marshal(tokenMap)
	err = json.Unmarshal(tokenJson, &token)
	if err != nil {
		log.Println("error fetching spotify token from db: ", err)
		return nil
	}

	return token
}

func SetSpotifyToken(token *oauth2.Token) {
	err := db.Delete(spotifyTokenQuery())
	if err != nil {
		log.Println("error deleting old spotify tokens:", err)
	}

	tokenJson, _ := json.Marshal(token)
	tokenMap := make(map[string]interface{})
	json.Unmarshal(tokenJson, &tokenMap)

	spotifyTokenConfig := document.NewDocument()
	spotifyTokenConfig.SetAll(tokenMap)
	spotifyTokenConfig.Set(spotifyTokenKey, true)

	_, err = db.InsertOne(configCollection, spotifyTokenConfig)
	if err != nil {
		log.Println("error creating spotify token:", err)
	}
}

func spotifyTokenQuery() *query.Query {
	return query.NewQuery(configCollection).Where(query.Field(spotifyTokenKey).Exists())
}
