In Go, files are identified by their filename extension. Go source code files have the .go extension. For example, a file named main.go is a Go source code file.

In addition to .go files, Go projects may also include other types of files, such as documentation files, configuration files, and test files. These files may have different filename extensions, depending on their purpose.

Go uses a package system to organize code. Each Go source code file belongs to a package, which is specified at the top of the file with the package keyword. The package name is used to identify the package when it is imported by another package. For example, a file with the following package declaration belongs to the main package:


```
package main
```

The main package is a special package in Go that is used to build executable programs. When you build a Go program, the main package is compiled into an executable binary.

In addition to the main package, Go also has a number of standard library packages that provide common functionality, such as input/output, networking, and data manipulation. These packages are imported using their package name, like this:


```
import "fmt"
```

This imports the fmt package, which provides input/output functions.

You can also create your own custom packages by defining them in their own Go source code files and using the package keyword to specify the package name. Custom packages can be imported by other packages using their package name.


Q: How would we write a function in go to take the name of Zeekus from STDIN and output Hello Mr. Zeekus ? 

A: To write a Go function that reads a name from standard input (stdin) and outputs a greeting using that name, you can use the fmt package's Scanf function to read the input and the Printf function to print the output.

Here is an example of a function that does this:

```
package main

import "fmt"

func greet(name string) {
  fmt.Printf("Hello Mr. %s\n", name)
}

func main() {
  var name string
  fmt.Scanf("%s", &name)
  greet(name)
}
```

This function defines a greet function that takes a name parameter and prints a greeting using that name. The main function reads a name from stdin using the Scanf function and passes it to the greet function.

To run this program, you can save it to a file with the .go extension (e.g., main.go) and use the go run command to execute it:

```
$ go run main.go
Zeekus
Hello Mr. Zeekus
```


When you run the program, it will prompt you to enter a name. When you enter a name and press Enter, the program will output the greeting with the name you entered.
