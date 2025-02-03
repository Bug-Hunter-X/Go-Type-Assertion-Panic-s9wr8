```go
func main() {
    var i interface{}
    i = 5
    j := i.(int)
    fmt.Println(j)
    i = "hello"
    j = i.(int) // This will panic
    fmt.Println(j)
}
```