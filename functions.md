# Functions in Go

Functions in Go are first-class objects, meaning they can be assigned to variables, passed as arguments to other functions, and returned as function results. Go provides support for multiple return values, named return values, and variadic parameters, ensuring efficient function handling.

## Basic Function Declaration
A function in Go consists of the `func` keyword, a function name, parameter list, return type (optional), and a function body enclosed in curly braces. Functions must always have a defined signature that specifies the input parameters and return values.
```go
func add(a, b int) int {
    return a + b
}

result := add(10, 5)
fmt.Println(result) // Output: 15
```

## Void Function Declaration
A void function does not return any value. It is primarily used for operations such as logging, updating global states, or triggering side effects.
```go
func logMessage(level, message string) {
    fmt.Printf("[%s] %s\n", level, message)
}

logMessage("INFO", "Launching Go app")
// Output: [INFO] Launching Go app
```

## Multiple Return Values
Go allows functions to return multiple values, which is useful for error handling and returning related data without defining additional structures.
```go
func divide(dividend, divisor int) (int, int, error) {
    if divisor == 0 {
        return 0, 0, fmt.Errorf("division by zero is not allowed")
    }
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder, nil
}

q, r, err := divide(13, 5)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Quotient: %d, Remainder: %d", q, r)
}
// Output: Quotient: 2, Remainder: 3
```

## Named Return Values
When using named return values, Go initializes the return variables with their zero values, and the return statement without arguments automatically returns those values.
```go
func rectangleArea(length, width int) (area int) {
    area = length * width
    return
}

fmt.Println(rectangleArea(10, 5)) // Output: 50
```

## Important Note: No Function Overloading

Go does not support function overloading. Each function must have a unique name, even if parameter types differ.

```go
func printInt(a int) {
    fmt.Println(a)
}

func printString(a string) {
    fmt.Println(a)
}
```

## Similar Parameter Type Shorthand
Go allows grouping parameters of the same type for concise declarations.
```go
func addMultiple(a, b, c int) int {
    return a + b + c
}

fmt.Println(addMultiple(1, 2, 3)) // Output: 6
```

## Variadic Functions
A function that accepts a variable number of arguments of the same type.
```go
func Greeting(prefix string, who ...string) {
    for _, person := range who {
        fmt.Println(prefix, person)
    }
}

Greeting("Hello", "Alice", "Bob", "Charlie")
// Output: Hello Alice
//         Hello Bob
//         Hello Charlie

func concatenate(words ...string) string {
    return strings.Join(words, " ")
}

fmt.Println(concatenate("Go", "is", "awesome"))
// Output: Go is awesome
```

## Pass by Value
In Go, function arguments are passed by value by default, meaning a copy of the argument is made. This ensures that changes made within the function do not affect the original value.
```go
func increment(x int) {
    x++
}

num := 5
increment(num)
fmt.Println(num) // Output: 5 (unchanged)
```

## Pass by Reference
To modify the original value of a variable, Go allows passing pointers to functions, which results in the function modifying the actual value rather than a copy.
```go
func updatePrice(price *float64) {
    *price *= 1.1
}

cost := 50.0
updatePrice(&cost)
fmt.Println("Updated Price:", cost) // Output: Updated Price: 55.0
```

## Anonymous Functions
Anonymous functions are functions without a name, useful for short-lived operations, closures, and function expressions.
```go
multiply := func(a, b int) int {
    return a * b
}

fmt.Println("Product:", multiply(3, 4)) // Output: Product: 12
```

## Recursive Functions

A function that calls itself to solve problems iteratively.

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5)) // Output: 120
```

## Closures

Closures in Go provide an elegant way to encapsulate and retain state across multiple function calls. They are widely used in various applications such as maintaining state, function chaining, and configuring behaviors dynamically. 

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

next := counter()
fmt.Println(next()) // Output: 1
fmt.Println(next()) // Output: 2
```

### Caching Mechanism
Closures can cache previously computed values to avoid redundant calculations.

```go
package main
import "fmt"

func memoize() func(int) int {
    cache := make(map[int]int)
    return func(n int) int {
        if result, found := cache[n]; found {
            return result
        }
        result := n * n // Example computation
        cache[n] = result
        return result
    }
}

func main() {
    compute := memoize()
    fmt.Println(compute(4)) // Output: 16
    fmt.Println(compute(4)) // Cached Output: 16
}
```
### Function Configuration

While Go doesn't support currying natively, it can be implemented using closures.

```go
package main
import "fmt"

func multiplier(factor int) func(int) int {
    return func(input int) int {
        return input * factor
    }
}

func main() {
    timesTwo := multiplier(2)
    timesThree := multiplier(3)
    
    fmt.Println(timesTwo(5))  // Output: 10
    fmt.Println(timesThree(5)) // Output: 15
}
```

### Function Composition

Closures can allow chaining functions to build pipelines dynamically.

```go
package main
import "fmt"

func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

func addOne(n int) int { return n + 1 }
func double(n int) int { return n * 2 }

func main() {
    pipeline := compose(addOne, double)
    fmt.Println(pipeline(3)) // Output: 7
}
```

### More Examples

```go
package main

import (
    "fmt"
    "time"
)

// Closure function to track user actions
func actionLogger(userID string) func(string) {
    actions := []string{}
    
    return func(action string) {
        actions = append(actions, fmt.Sprintf("%s: %s", time.Now().Format("15:04:05"), action))
        fmt.Printf("User [%s] performed action: %s\n", userID, action)
        fmt.Println("Current action log:", actions)
    }
}

func main() {
    userLogger := actionLogger("user123")

    userLogger("Visited Homepage")
    userLogger("Clicked 'Buy Now' Button")
    userLogger("Logged Out")

    fmt.Println("Session ended.")
}

```

## Higher-Order Functions

Higher-order functions accept functions as arguments or return them. They are commonly used for code abstraction, such as applying operations or implementing middleware patterns.

```go
func apply(op func(int, int) int, a, b int) int {
    return op(a, b)
}

sum := func(x, y int) int { return x + y }
sub := func(x, y int) int { return x - y }

fmt.Println(apply(sum, 4, 6)) // Output: 10
fmt.Println(apply(sub, 10, 3)) // Output: 7
```

### Filtering Collections
```go
package main
import "fmt"

type filterFunc func(int) bool

func filter(slice []int, f filterFunc) []int {
    var result []int
    for _, val := range slice {
        if f(val) {
            result = append(result, val)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6}
    fmt.Println("Original slice:", numbers)
    
    greaterThanThree := filter(numbers, func(n int) bool {
        return n > 3
    })
    fmt.Println("Numbers greater than 3:", greaterThanThree)
}
```

### Transforming Collections
```go
package main

import (
    "fmt"
    "strings"
)

func mapStrings(input []string, transform func(string) string) []string {
    result := make([]string, len(input))
    for i, v := range input {
        result[i] = transform(v)
    }
    return result
}

func main() {
    words := []string{"hello", "world", "golang"}
    
    uppercaseWords := mapStrings(words, strings.ToUpper)
    fmt.Println("Uppercase:", uppercaseWords)
    
    prefixedWords := mapStrings(words, func(s string) string {
        return "Prefix-" + s
    })
    fmt.Println("Prefixed:", prefixedWords)
}
```

### Function Decoration
```go
func logger(f func(string) string) func(string) string {
    return func(input string) string {
        fmt.Println("Processing input:", input)
        result := f(input)
        fmt.Println("Processed output:", result)
        return result
    }
}

func toLowerCase(s string) string {
    return strings.ToLower(s)
}

func main() {
    loggableFunction := logger(toLowerCase)
    fmt.Println(loggableFunction("HELLO WORLD"))
}
```

## Defer Statement

In Go, the defer keyword defers the execution of a function until the surrounding function returns. This mechanism is crucial for handling cleanup operations such as closing files, releasing resources, or logging activities. Deferred calls execute in a Last-In-First-Out (LIFO) order, ensuring systematic resource management.

```go
func fetchData() {
    defer fmt.Println("Closing database connection")
    fmt.Println("Fetching data from database")
}

fetchData()
// Output:
// Fetching data from database
// Closing database connection
```

### Resource Management

Using defer ensures that resources like files and network connections are properly closed.

```go
package main
import (
    "fmt"
    "os"
)

func handleFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("File open error:", err)
        return
    }
    defer file.Close()
    fmt.Println("Processing", filename)
}

handleFile("config.yaml")

// Even if an error occurs during file processing, the deferred file.Close() ensures resource release.
```

### Error Recovery

`defer` helps maintain program stability and consistency during unexpected runtime errors.

```go
package main
import "fmt"

func handlePanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    panic("Unexpected error!")
}

handlePanic()
fmt.Println("Continuing execution...")

// Output:
// Recovered: Unexpected error!
// Continuing execution...
```
### Execution Sequence

Multiple defer statements follow **LIFO** order, ensuring predictable cleanup sequencing.

```go
package main
import "fmt"

func sequence() {
    defer fmt.Println("Step 1: Finalization")
    defer fmt.Println("Step 2: Cleanup")
    fmt.Println("Executing main logic")
}

sequence()

// Output:
// Executing main logic
// Step 2: Cleanup
// Step 1: Finalization

func orderWithDefer() {
    for i := 1; i <= 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}

orderWithDefer()
```

### Execution Tracing

Tracing function execution can be simplified using defer for logging entry and exit points.

```go
package main
import "fmt"

func logExecution(fnName string) {
    fmt.Println("Start:", fnName)
    defer fmt.Println("End:", fnName)
    fmt.Println("Running", fnName)
}

logExecution("operationX")

// Output:
// Start: operationX
// Running operationX
// End: operationX
```

### Benchmarking Execution Time

Using defer to measure function execution time for performance analysis.

```go
package main
import (
    "fmt"
    "time"
)

func measureExecution() {
    start := time.Now()
    defer func() {
        fmt.Println("Execution time:", time.Since(start))
    }()
    time.Sleep(2 * time.Second)
}

measureExecution()

// Output:
// Execution time: 2s
```
### Transaction Handling
defer can be used to commit or rollback database transactions.

```go
package main
import "fmt"

func manageTransaction(success bool) {
    defer func() {
        if success {
            fmt.Println("Transaction committed")
        } else {
            fmt.Println("Transaction rolled back")
        }
    }()
    fmt.Println("Processing transaction")
}

manageTransaction(false)

// Output:
// Processing transaction
// Transaction rolled back
```
