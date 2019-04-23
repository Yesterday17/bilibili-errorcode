bilibili-errorcode
===

基于`[来源请求]``[数据删除]`得到的`[数据删除]`站的错误代码对应表。

## Get Started
```go
package example

import (
	"fmt"

	ecode "github.com/Yesterday17/bilibili-errorcode"
)

func main() {
	var code = ecode.ErrorCode(-403)
	fmt.Printf("Region: %s\n", code.GetRegion())
	fmt.Printf("Code: %d\n", code.GetDetail().Code)
	fmt.Printf("Message: %s\n", code.GetDetail().Message)
	fmt.Printf("Description: %s\n", code.GetDetail().Description)
}
```

## LICENSE
[MIT](LICENSE)