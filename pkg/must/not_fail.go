package must

func NotFailFn(fn func() error) {
	NotFail(fn())
}

func NotFail(err error) {
	if err != nil {
		panic(err)
	}
}
