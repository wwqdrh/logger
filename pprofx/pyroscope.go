package pprofx

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/pyroscope-io/client/pyroscope"
	"github.com/wwqdrh/logger"
)

// appName: simple.golang.app
// scopeUrl: http://pyroscope-server:4040
func Start(ctx context.Context, appName, scopeUrl string, options TypeOptions) {
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	prof, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: appName,
		// replace this with the address of pyroscope server
		ServerAddress: scopeUrl,
		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,
		// optionally, if authentication is enabled, specify the API key:
		// AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN"),
		ProfileTypes: options,
	})
	if err != nil {
		logger.DefaultLogger.Error(err.Error())
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-ctx.Done():
		if err := prof.Stop(); err != nil {
			logger.DefaultLogger.Error(err.Error())
		}
	case <-quit:
		if err := prof.Stop(); err != nil {
			logger.DefaultLogger.Error(err.Error())
		}
	}
}

// 便于进行筛选
// 例如sql或者controller或者其他
func AddTag(labels []string, fn func()) {
	// these two ways of adding tags are equivalent:
	pyroscope.TagWrapper(context.Background(), pyroscope.Labels(labels...), func(c context.Context) {
		fn()
	})

	// pprof.Do(context.Background(), pprof.Labels("controller", "slow_controller"), func(c context.Context) {
	// 	slowCode()
	// })
}
