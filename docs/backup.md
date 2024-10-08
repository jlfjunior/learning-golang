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
- interface: Defines a contract or behavior that a type must fulfill.
- select
- case 
- defer
- go           
- map          
- struct: is used to define a data type that groups together variables, fields or members
- chan         
- else         
- goto         
- package      
- switch
- const: used to define variable that cannot be the value changed       
- fallthrough  
- if           
- range: used to iterates in for loops over a slice or map.      
- type
- continue     
- for: is used for iterative control flow. Golang don't has While, but it doesn't mean that dont have infinite loop.
- import: is used to include external packages in your program.       
- return: used to terminate the execution of a function and return control to the calling function.
- var: is used to declare variables.

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