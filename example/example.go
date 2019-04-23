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
