package main

import ("fmt")

func main() {
	//count backwards 10 to 1
	for i := 10; i > 0; i-- {
          //check if we are a factor of 2
          if i %2 == 0 {
	    str := fmt.Sprintf("i:%d is a factor of 2",i)
	    fmt.Println(str)
	  //check if we are not a factor of 2
          } else {
	    str := fmt.Sprintf("i:%d is not a factor of 2",i)
	    fmt.Println(str)
         }
       }
}
