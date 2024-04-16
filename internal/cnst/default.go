package cnst

import (
	"time"

	"github.com/avast/retry-go/v4"
)

var DefaultRetryOpts = []retry.Option{
	retry.LastErrorOnly(true),
	retry.MaxDelay(time.Second),
	retry.Attempts(3),
	retry.Delay(time.Millisecond * 100),
}
