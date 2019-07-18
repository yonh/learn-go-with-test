package main

func Repeat(s string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += s
	}

	return repeated
}
