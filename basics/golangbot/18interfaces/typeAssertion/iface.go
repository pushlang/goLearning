//TYPE ASSERTION. i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.

package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}
func main() {  
    var s interface{} = 56
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)
}