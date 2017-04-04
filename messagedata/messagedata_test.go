package messagedata

import (
	"fmt"
	"testing"
)

func TestGetId(t *testing.T) {
	//	fmt.Println(Message())
	tstMsg := messagedata.Message
	fmt.Println(tstMsg.getId())
}
