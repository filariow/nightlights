module github.com/filariow/nightlights

go 1.19

replace (
	device/sam => /usr/lib/tinygo/src/device/sam
	internal/reflectlite => /usr/lib/tinygo/src/internal/reflectlite
	internal/task => /usr/lib/tinygo/src/internal/task
	machine => /usr/lib/tinygo/src/machine
	os => /usr/lib/tinygo/src/os
	reflect => /usr/lib/tinygo/src/reflect
	runtime => /usr/lib/tinygo/src/runtime
	runtime/interrupt => /usr/lib/tinygo/src/runtime/interrupt
	runtime/volatile => /usr/lib/tinygo/src/runtime/volatile
	sync => /usr/lib/tinygo/src/sync
	testing => /usr/lib/tinygo/src/testing
)
