package behaviour

import (
	"context"
	"time"
)

func Setup() {
	ctx := context.Background()
	behaviourWrapper(ctx, OsUpdatesBehaviour, 5*time.Second)
	time.Sleep(10 * time.Second)
	behaviourWrapper(ctx, SpotifyCacheBehaviour, 1*time.Hour)
}

func behaviourWrapper(ctx context.Context, f func(context.Context), duration time.Duration) {
	go func() {
		for {
			f(ctx)
			time.Sleep(duration)
		}
	}()
}
