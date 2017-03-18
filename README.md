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

// Output
[sudo npm start git status git add client/package.json git commit -m "Update react-scripts" git push git status git diff client/src/components/login.js git add client/src/components/login.js git commit -m "Close #7" git status git push npm start git status git diff git add server.js git commit -m "Fix typo" git push git status sudo npm start sudo npm start sudo npm start sudo npm start cd dev/hangouts-clone-react/ ls git status git diff ls cat .env  cp .env example.env vim example.env ]
```
