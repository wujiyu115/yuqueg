# yuque_go

example:

```golang
package main

import (
	"encoding/json"
)

func main() {
	yu := NewService("token")
	d, err := yu.User.Get("")
	if err != nil {
		l.Info(err)
		return
	}
	jsonString, _ := json.Marshal(d)
	l.Info(string(jsonString))
}
```