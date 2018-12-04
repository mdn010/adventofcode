package common

func Checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
