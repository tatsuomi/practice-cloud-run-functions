package main

import "fmt"

func main() {
	fmt.Print(Hello("World"))
}

// Hello: Hello World 実行用関数
func Hello(s string) string {
	return fmt.Sprintf("Hello %s", s)
}
