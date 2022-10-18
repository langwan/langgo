---
title: "sqlite"
---
# sqlite

sqlite 内部封装了 grom

## 配置

```yaml
sqlite:
  path: "app.db"
```

## 使用

```go
package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/sqlite"
	"gorm.io/gorm"
	"os"
)

type account struct {
	gorm.Model
	Name string
}

func main() {
	dbPath := "./app.db"
	os.Remove(dbPath)
	defer os.Remove(dbPath)
	langgo.Run(&sqlite.Instance{})
	err := sqlite.Get().AutoMigrate(&account{})
	if err != nil {
		panic(err)
	}
	acc := account{Name: "langgo"}
	sqlite.Get().Create(&acc)
	acc2 := account{}
	sqlite.Get().First(&acc2, "name=?", acc.Name)
	fmt.Printf("acc2 name is %s\n", acc2.Name)
}
```
