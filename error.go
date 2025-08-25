package main

// Generic error handler that lets the program panic if error
func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}

	return x
}
