package featureUnit

import (
	"context"
	"testing"

	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
)

func Test_ExternalContact_Del_Contact_Way(t *testing.T) {
	response, _ := Work.ExternalContactContactWay.Delete(context.Background(), "f3626f74a7f94784115b0b8a729c471f")

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}

	fmt.Dump(response)
}
