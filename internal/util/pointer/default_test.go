package pointer

import (
	"encoding/json"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/server/http/response"
	"testing"
)

func TestDefault(t *testing.T) {
	var value = Default(response.BaseDataList[int]{})

	bs, _ := json.Marshal(value)

	fmt.Println(string(bs))
}
