package liface

type IServer interface {
	Start()
	Stop()
	Serve()
}
