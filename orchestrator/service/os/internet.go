package os

import (
	"net/http"
	"time"
)

const (
	FORCE_INTERNET_DISABLED = false
)

func CheckForInternet() bool {
	if FORCE_INTERNET_DISABLED {
		return false
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
