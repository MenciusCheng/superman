package dragons

type Option func(*Dragons)

func ConfigPath(path string) Option {
	return func(o *Dragons) {
		o.ConfigPath = path
	}
}
