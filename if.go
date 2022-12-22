package main

import "fmt"

func main() {

	#loop 10 - 0
	for i := 10; i > 0; i-- {
          #check if we are a factor of 2
          if i %2 == 0 {
	    str := fmt.Sprintf("i:%d is a factor of 2",i)
	    fmt.Println(str)
          } else {
	    str := fmt.Sprintf("i:%d is not a factor of 2",i)
	    fmt.Println(str)
         }
       }
}
