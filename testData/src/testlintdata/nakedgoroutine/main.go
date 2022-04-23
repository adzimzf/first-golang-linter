package main

func main() {
	go func() {
		print("Hello")
	}()
}
