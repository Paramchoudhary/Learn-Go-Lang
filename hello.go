package main


import (
    "fmt"
    "os"
)


func main(){

	namee := "Param"
	fmt.Println("Hello world" ,namee);
	if len(os.Args) < 2 {
        fmt.Println("Usage: go run hello.go <your_name>")
        os.Exit(1)
    }

    name := os.Args[1]
    fmt.Printf("Hello, %s!\n", name)
}

