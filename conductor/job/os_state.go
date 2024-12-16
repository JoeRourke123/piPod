package job

import (
	"conductor/common/model"
	"conductor/db/insert"
	"conductor/util/logger"
	"context"
	"net/http"
	"strconv"
	"time"
)

const (
	FORCE_INTERNET_DISABLED = false
)

var (
	UpdateOsState = Job{
		Interval: time.Second * 10,
		Handler: func(ctx context.Context) {
			isInternetEnabled := checkForInternet()
			newOsState := model.OsState{
				IsInternetEnabled: isInternetEnabled,
			}
			err := insert.OsState(&newOsState)
			if err != nil {
				logger.Error(ctx, "error updating os state", err, logger.FromTag("UpdateOsState"), logger.DbWriteTag)
			} else {
				logger.Info(ctx, "os state updated, isInternetEnabled="+strconv.FormatBool(isInternetEnabled), logger.FromTag("UpdateOsState"), logger.DbWriteTag)
			}
		},
	}
)

func checkForInternet() bool {
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
