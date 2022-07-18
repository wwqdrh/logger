package pprofx

import (
	"context"
	"runtime"

	"github.com/pyroscope-io/client/pyroscope"
)

// appName: simple.golang.app
// scopeUrl: http://pyroscope-server:4040
func Start(appName, scopeUrl string, options TypeOptions) {
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	pyroscope.Start(pyroscope.Config{
		ApplicationName: appName,
		// replace this with the address of pyroscope server
		ServerAddress: scopeUrl,
		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,
		// optionally, if authentication is enabled, specify the API key:
		// AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN"),
		ProfileTypes: options,
	})
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
