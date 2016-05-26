# Struct map
This module is used to convert struct into map recursively

### Example
```go
import stm "github.com/alileza/go-structmap"

type MyStruct struct{}

func main(){
    myStruct := &MyStruct{}
    fmt.Println(stm.StructToMap(*myStruct))

    // if you put true on the second parameter
    // it will convert all of your datatype into string
    // unless (bool, nil)
    fmt.Println(stm.StructToMap(*myStruct, true))
}
```
