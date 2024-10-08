# Functions
In Golang, functions can be package functions or exported functions. This is similar to private and public functions in other languages.

## Package function. 
Only can be accessed within the package where it was defined. You just need to use lowercase in the first letter.
``` Go
func sum(a, b int) int {
	return a + b
}
```

## Exported function. 
You can access it from outside the package where it has been defined. You just need to use uppercase in the first letter.
``` Go
func Sum(a, b int) int {
	return a + b
}
```

## With Generics
``` Go
func sum[T int | float64](a, b T) T {
	return a + b
}
```