package middleware

type MiddleWare struct {
	Runtime Runtime
}

func New() *MiddleWare {
	return &MiddleWare{}
}
