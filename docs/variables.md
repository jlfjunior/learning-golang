# Constantes, Variable and Type declaration

## Constantes
Are immutable values declared with const keyword. Once declared, a constante cannot be changed or reassigned.

``` Go
const PI = 3.1415
```

## Variables
Can be declared using var keyword or  by inference using the symbol := . Once declared, a variable can have its value updated.

``` Go
//Long declaration 
var name = "Jose"
name = "Jose Junior"

//Short declaration
age := 10
age = 29
```

## Primitive Types

### Strings
It represents a sequence of characters. Strings are used to store text.
``` Go
var name string = "Jose Junior"
```

### Integer 
It represents an integer. The int type can vary in length. It can be 32 or 64 bits.
``` Go
var year int = 2024
```

### Float
It represents a floating point number, allowing decimal values to be stored accurately. Also, it can be 32 or 64 bits.
``` Go
var salary float64 = 2000.99
```

### Boolean 
It represents a boolean value, which can be true or false.
``` Go
var isEmployeed bool = true
```

### Byte 
It represents an 8-bit value (1 byte). It is often used to store byte data and is equivalent to the uint8 type.
``` Go
var name byte = 'A'
```
### Rune
It represents a 32-bit value used to store Unicode characters. It is equivalent to the int32 type and is useful for representing characters from different languages 
``` Go
var a rune = 'a'
```