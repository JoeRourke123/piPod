package podcasts

import (
	"bytes"
	"conductor/common/model"
	"conductor/db/fetch"
	"conductor/db/insert"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func RssFeed(podcast *model.Album) string {
	cachedRssUrl := fetch.RssFeedCache(podcast.Id)
	if cachedRssUrl != "" {
		return cachedRssUrl
	}

	castosResponse, err := searchCastos(podcast.Name, podcast.Artist)
	if err == nil && castosResponse != nil && len(castosResponse.Data) > 0 {
		insert.RssFeedCache(podcast.Id, castosResponse.Data[0].Url)
		return castosResponse.Data[0].Url
	}

	return ""
}

func searchCastos(query ...string) (*CastosResponse, error) {
	url := "https://castos.com/wp-admin/admin-ajax.php"
	searchQuery := strings.Join(query, " ")

	body := fmt.Sprintf(`------WebKitFormBoundaryLmDJsRsCoKZFRuXt
Content-Disposition: form-data; name="search"

%s
------WebKitFormBoundaryLmDJsRsCoKZFRuXt
Content-Disposition: form-data; name="action"

feed_url_lookup_search
------WebKitFormBoundaryLmDJsRsCoKZFRuXt--`, searchQuery)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8,fr;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundaryLmDJsRsCoKZFRuXt")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var castosResponse CastosResponse
	if err := json.Unmarshal(responseBody, &castosResponse); err != nil {
		return nil, err
	}

	return &castosResponse, nil
}
