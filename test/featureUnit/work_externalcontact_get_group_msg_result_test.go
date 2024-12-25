package featureUnit

import (
	"context"
	"testing"

	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
)

func Test_ExternalContact_Get_Group_MSG_Result(t *testing.T) {
	msgID := "msg_ViZBwAA72X9XCh4Cx5ku9OVFb2thQ" // walle
	// msgID := "msg_ViZBwAAT_5pzEfUhaT2xbSbFxhTgw"  // matt
	response, _ := Work.ExternalContactGroupChat.Get(context.Background(), msgID, 0)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}

	fmt.Dump(response)
}
