package main

func main() {
	x := "hello, world"
	ch := make(chan string)
	go f(ch)
	sink(x) // sink, no leak
	x = source()
	ch <- x
}
func f(ch_1 chan string) {
	y := <-ch_1
	sink(y) // sink, leak
}
func sink(s string) {}
func source() string {
	return "secret"
}
