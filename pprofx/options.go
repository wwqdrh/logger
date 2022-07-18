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
