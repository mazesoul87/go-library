package gouuid_test

import (
	"fmt"
	"gopkg.in/dtapps/go-library.v2/utils/gouuid"
	"testing"
)

func TestGetUuId(t *testing.T) {
	fmt.Println(gouuid.GetUuId())
}
