package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/samber/lo"
)

// Redacter defines how to log an object
type Redacter interface {
	Redact() string
}

// Server is an server logging middleware.
func LogServer(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				operation = info.Operation()
			}

			reply, err = handler(ctx, req)

			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}

			level, stack := extractErrors(err)
			_ = log.WithContext(ctx, logger).Log(level,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime).Seconds(),
			)

			return
		}
	}
}

// extractError returns the string of the error
func extractErrors(err error) (log.Level, string) {
	if err != nil {
		var kerr *errors.Error
		text := ""
		if errors.As(err, &kerr) {
			stack := getStack(kerr.Unwrap())
			text += fmt.Sprintf(`%s msg="%s" meta=%s`, stack, kerr.Unwrap(), kerr.GetMetadata())
		} else {
			stack := getStack(err)
			text = fmt.Sprintf(`%s msg="%s"`, stack, err.Error())
		}
		return log.LevelError, text
	}
	return log.LevelInfo, ""
}

func getStack(err error) (stack string) {
	text := fmt.Sprintf("%+v", err)
	lines := strings.Split(text, "\n")

	lines = lo.Filter(lines, func(cnt string, _ int) bool {
		return strings.Index(cnt, "\t") != -1
	})

	lines = lo.Slice(lines, 0, 4)
	lines = lo.Map(lines, func(item string, index int) string { return strings.TrimSpace(item) })
	lines = lo.Compact(lines)

	stack = strings.Join(lines, ",")

	return
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	j, err := json.Marshal(req)
	if err == nil {
		return string(j)
	}

	if redacter, ok := req.(Redacter); ok {
		return redacter.Redact()
	}
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}
