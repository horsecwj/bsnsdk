package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/contracts"
)

func TestGetBlockByNumber(t *testing.T) {

	fmt.Println(client.GetBlockByNumber(301751))
}
func TestGetEvents(t *testing.T) {

	fmt.Println(client.GetBlockEvents(301751))
}
func TestGetEvent(t *testing.T) {
	//0x10c5facd78a4bf8a25b243b3d1d684ed4183674b6256648015ed08a3ef022940
	events, _ := client.GetTxEvents(common.HexToHash("0x10c5facd78a4bf8a25b243b3d1d684ed4183674b6256648015ed08a3ef022940"))

	for _, e := range events {
		res, ok := e.(*contracts.DDC721Transfer)
		if !ok {
			continue
		}
		fmt.Println(res.DdcId)
	}
}
