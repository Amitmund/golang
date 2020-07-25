# Go Quick Notes

## `package` command, to Start with 

```go
package main
```
## `import` command to import packages:

```go
import (
    "fmt"
)
```


## `func` to define a function.

```go
func main() {
    ...
}

```

## Example code

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello world.")
}
```

## Few data types

```go
fmt.Println("go" + "lang")
fmt.Println("1+1 =", 1+1)
fmt.Println("7.0/3.0 =" 7.0/3.0)

fmt.Println(true && false)
fmt.Println(true || false)
fmt.Println(!true)
```

## Variables

```go
var a = "initial"
var b, c int = 1, 2
var d = true
var e int
// := syntax is shorthand for declaring and initializing a variable. 
f := "apple"
```

## `const`, few example

```go
const s string = "constant"

fmt.Println(int64(d))
fmt.Println(math.Sin(n))

```

## `for` example 0

```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
```

## `for` example 1

initial/condition/after example

```go
for j := 7; j <= 9; j++ {
    fmt.Println(j)
}
```

## `for` example 2
without condition continious loop, till `break` or `return`

```go
for{
    fmt.Println("loop")
    break
}
```

## `for` example 3
You can also `continue` to the next iteration of the loop.

```go
for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
```


## `if/else` Example

