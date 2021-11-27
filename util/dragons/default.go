package dragons

var Default = New()

func Init(options ...Option) {
	Default.Init(options...)
}

func Shutdown() error {
	return Default.Shutdown()
}

func Scan(v interface{}) error {
	return Default.Scan(v)
}
