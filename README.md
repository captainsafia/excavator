# excavator
excavator is a Go library that fetches data from the user's shell history into a 
string array.

## Usage
```
import "fmt"
import "github.com/captainsafia/excavator"

commands, err := excavator.GetHistory("bash", 30)
if err != nil {
    fmt.Println(err.Error())
}
fmt.Println(commands)
```
