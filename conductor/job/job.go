package job

import (
	"context"
	"time"
)

type Job struct {
	Interval time.Duration
	Handler  func(ctx context.Context)
}

func (j *Job) Start(ctx context.Context) {
	go func() {
		for {
			j.Handler(ctx)
			time.Sleep(j.Interval)
		}
	}()
}

func Init(ctx context.Context) {
	UpdateOsState.Start(ctx)
	RefreshAlbums.Start(ctx)
	RefreshPlaylists.Start(ctx)
}
