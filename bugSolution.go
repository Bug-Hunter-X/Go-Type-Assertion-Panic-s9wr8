```go
func main() {
    var i interface{}
    i = 5
    switch j := i.(type) {
    case int:
        fmt.Println(j)
    case string:
        fmt.Println("String value:", j)
    default:
        fmt.Println("Unknown type")
    }
    i = "hello"
    switch j := i.(type) {
    case int:
        fmt.Println(j)
    case string:
        fmt.Println("String value:", j)
    default:
        fmt.Println("Unknown type")
    }
}
```