package core

var WorkDir string
var EnvName string

const (
	Development = "development"
	Production  = "production"
)

type DeferHandle func()

var deferHandles []DeferHandle

func DeferRun() {
	for _, foo := range deferHandles {
		foo()
	}
}

func DeferAdd(handle DeferHandle) {
	deferHandles = append(deferHandles, handle)
}
