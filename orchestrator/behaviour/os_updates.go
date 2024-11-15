package behaviour

import (
	"context"
	"orchestrator/service/db"
	"orchestrator/service/os"
	"orchestrator/ui/model"
)

func OsUpdatesBehaviour(ctx context.Context) {
	isInternetEnabled := os.CheckForInternet()

	osUpdates := model.OsUpdates{
		IsInternetEnabled: isInternetEnabled,
	}

	//logger.Info(ctx, "os updates fetched, isInternetEnabled="+strconv.FormatBool(isInternetEnabled), logger.FromTag("OsUpdatesBehaviour"), logger.DbWriteTag)

	db.SetLatestOsUpdates(osUpdates)
}
