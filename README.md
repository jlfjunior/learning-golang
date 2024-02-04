# Learning Golang

### What is Golang?
Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from packages, whose properties allow efficient management of dependencies.


### How to do comments in Go?

Single comments
``` Go
//var x = 0
```

Block of comments
``` Go
/* 
Usually used for comments that used more than 2 lines!
Like this!
*/
```

### Reserved Keywords
- break: used to stop a loop execution prematurely.
- default: used in switch statament to define a default behavior when none of the other match. 
- func:  used to declare function. Functions can be public or private
- interface: 
- select
- case 
- defer
- go           
- map          
- struct
- chan         
- else         
- goto         
- package      
- switch
- const        
- fallthrough  
- if           
- range        
- type
- continue     
- for          
- import       
- return       
- var

### Variable declarations

Explicit declaration
``` Go
var age int
var age = 10
var salary, descont int
var salary, descont = 100, 0.10
var animal = "Dog"
var (
	age    int
	animal string
)
```

Short or Implicit declaration
``` Go
age := 10
salary, descont := 200, 0.10
```

### Function declarations

Package function. You just need to use lowercase in the first letter.
``` Go
func sum(a, b int) int {
	return a + b
}
```
Exported function. You just need to use uppercase in the first letter.
``` Go
func Sum(a, b int) int {
	return a + b
}
```

With Generics
``` Go
func sum[T int | float64](a, b T) T {
	return a + b
}
```