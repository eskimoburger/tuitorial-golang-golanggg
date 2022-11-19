package main

func main() {
	book := make(map[string]int)
	book["fuck"] = 1
	book["you"] = 2
	println(book["fuck"], book["you"])

}
