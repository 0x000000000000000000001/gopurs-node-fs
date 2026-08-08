package main
import "fmt"
import "gopurs/output/gopurs_runtime"

func main() {
    v := gopurs_runtime.Box(2)
    fmt.Printf("%v\n", v)
    fmt.Printf("%+v\n", v)
    fmt.Printf("%#v\n", v)
}
