package main

func main() {
	x := "Hello World"
	ch := make(chan string)
	// @expectedflow: false
	sink(x) // sink, no leak
	x = source()
	ch <- x
	go f(ch)
}

func f(ch chan string) {
	y := <-ch
	// @expectedflow: true
	sink(y) // sink, leak
}

func sink(s string) {
}

func source() string {
	return "secret"
}
