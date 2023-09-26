package middleware

type MiddleWare struct {
	Runtime   Runtime
	Validator *Validator
}

func New(valid *Validator) *MiddleWare {
	return &MiddleWare{
		Validator: valid,
	}
}
