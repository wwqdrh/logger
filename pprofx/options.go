package pprofx

import "github.com/pyroscope-io/client/pyroscope"

var AllTypeOptions = []pyroscope.ProfileType{
	pyroscope.ProfileCPU,
	pyroscope.ProfileAllocObjects,
	pyroscope.ProfileAllocSpace,
	pyroscope.ProfileInuseObjects,
	pyroscope.ProfileInuseSpace,
	pyroscope.ProfileGoroutines,
	pyroscope.ProfileMutexCount,
	pyroscope.ProfileMutexDuration,
	pyroscope.ProfileBlockCount,
	pyroscope.ProfileBlockDuration,
}

type TypeOptions = []pyroscope.ProfileType

type PprofOption struct {
	Server  string // scope url
	AppName string // 应用名字
	Options []pyroscope.ProfileType
}

type Option func(*PprofOption)

func NewPprofOption(server, appname string, options ...Option) *PprofOption {
	profOption := &PprofOption{
		Server:  server,
		AppName: appname,
	}
	for _, item := range options {
		item(profOption)
	}
	return profOption
}

func WithPprofType(pprofType ...pyroscope.ProfileType) Option {
	return func(po *PprofOption) {
		po.Options = append(po.Options, pprofType...)
	}
}
