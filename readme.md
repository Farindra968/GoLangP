

# 🎉 **Golang Full Documentation (From Basic to Advanced)**

---

## 🧩 1. What is Go?

**Go (or Golang)** is a statically typed, compiled programming language developed by **Google**.
It’s known for **speed, concurrency, simplicity, and scalability**.

### ✅ Features:

* Fast compilation
* Garbage collected
* Built-in concurrency (`goroutines`)
* Great for backend APIs, microservices, and system tools

### ⚙️ Requirements:

Install Go → [https://go.dev/dl/](https://go.dev/dl/)

After installation:

```bash
go version
```

---

## 📁 2. Go Project Structure

Example folder:

```
myproject/
 ┣ main.go
 ┗ go.mod
```

### Initialize a Go module:

```bash
go mod init myproject
```

---

## 🧠 3. Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go World!")
}
```

Run it:

```bash
go run main.go
```

---

## 🔤 4. Variables and Data Types

```go
package main

import "fmt"

func main() {
    var name string = "Farindra"
    var age int = 25
    var height = 5.8 // Type inferred
    country := "Nepal" // Short variable declaration

    fmt.Println(name, age, height, country)
}
```

---

## 🔢 5. Constants

```go
const Pi = 3.14
const Greeting = "Namaste Nepal"
```

---

## ➕ 6. Operators

```go
a, b := 10, 5
fmt.Println("Add:", a+b)
fmt.Println("Sub:", a-b)
fmt.Println("Mul:", a*b)
fmt.Println("Div:", a/b)
```

---

## 🔁 7. Conditional Statements

```go
if age := 20; age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

---

## 🔂 8. Loops

Only `for` loop exists in Go.

```go
for i := 1; i <= 5; i++ {
    fmt.Println("Number:", i)
}
```

---

## 📦 9. Functions

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 5)
    fmt.Println("Sum:", result)
}
```

---

## 🧩 10. Arrays and Slices

### Array (Fixed size)

```go
var nums = [3]int{1, 2, 3}
fmt.Println(nums)
```

### Slice (Dynamic)

```go
fruits := []string{"apple", "banana"}
fruits = append(fruits, "mango")
fmt.Println(fruits)
```

---

## 📋 11. Maps (Key-Value)

```go
person := map[string]string{
    "name": "Farindra",
    "city": "Morang",
}
fmt.Println(person["name"])
```

---

## 🧱 12. Structs (Custom Data Types)

```go
type Student struct {
    Name  string
    Age   int
    Grade string
}

func main() {
    s1 := Student{Name: "Saroj", Age: 20, Grade: "A"}
    fmt.Println(s1)
}
```

---

## ⚙️ 13. Methods

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    c := Circle{Radius: 5}
    fmt.Println("Area:", c.Area())
}
```

---

## 🔄 14. Pointers

```go
x := 10
p := &x
fmt.Println(*p) // Output: 10
*p = 20
fmt.Println(x)  // Output: 20
```

---

## ⏳ 15. Goroutines (Concurrency)

```go
package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Hello from Goroutine")
}

func main() {
    go hello()
    time.Sleep(time.Second)
    fmt.Println("Main function done")
}
```

---

## 📬 16. Channels

```go
ch := make(chan string)

go func() {
    ch <- "Hello Channel"
}()

msg := <-ch
fmt.Println(msg)
```

---

## 🧩 17. Error Handling

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

---

## 🌐 18. Simple HTTP Server (Backend API)

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go Server!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
```

Run:

```bash
go run main.go
```

Visit:
👉 [http://localhost:8080](http://localhost:8080)

---

## 🗂️ 19. Go Folder Structure for Backend (Example)

```
go-backend/
 ┣ main.go
 ┣ go.mod
 ┣ handlers/
 ┃ ┗ userHandler.go
 ┣ models/
 ┃ ┗ user.go
 ┣ routes/
 ┃ ┗ routes.go
 ┗ utils/
    ┗ database.go
```

---

## ⚙️ 20. Compile and Build

To make a binary file:

```bash
go build main.go
```

To run the compiled binary:

```bash
./main
```
