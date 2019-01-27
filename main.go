package main

func main() {
	a := App{}
	a.Initialize("root", "1234", "test2")
	a.Run(":8080")
}
