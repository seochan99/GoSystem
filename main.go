// package main
package main

import (
	"fmt"

	"github.com/seochan99/goSystem/something"
)

// func main
// main은 컴파일을 위해서 존재
func main() {
	// export 하고싶으면 대문자로 시작
	fmt.Println("Hello, World!")
	
	println("한글도 잘 나오나?")
	something.SayHello()
	
}
