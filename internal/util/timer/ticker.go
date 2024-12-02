package timer

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"time"
)

func NewMinuteTicker(ctx context.Context, f func(ctx context.Context, ticker *time.Ticker, now time.Time)) {
	var ticker = time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			if 0 == now.Second() {
				f(ctx, ticker, now)
			}
		}
	}
}
