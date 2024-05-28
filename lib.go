package workspacelib

import "fmt"

func AddNums(a, b int) int {
	return a + b
}

func Hello(msg string) string {
	return fmt.Sprintf("Hello %v", msg)
}
