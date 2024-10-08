# Function declarations

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