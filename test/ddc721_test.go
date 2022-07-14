package test

import (
	"fmt"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var ddc721Service = client.GetDDC721Service()
var nftServerice = client.GetNftService()

func TestMint(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	temp := ddc721Service
	fmt.Println(temp)
	tx, err := ddc721Service.Mint(opts, platform, "www.123.com")
	if err != nil {
		log.Error.Println(err.Error())
		return
	}
	fmt.Println(tx.Hash())
}

func TestSafeMint(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.SafeMint(opts, platform, "www.123.com", []byte{}))
}

func TestMyMint(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	res, err := nftServerice.MyMint(opts, platform, "www.123.com")
	if err != nil {
		return
	}
	txHash := res.Hash().Hex()
	fmt.Println(txHash)
}
func TestDDC721BalanceOf(t *testing.T) {
	fmt.Println(ddc721Service.BalanceOf(platform))
}

func TestApprove(t *testing.T) {
	fmt.Println(ddc721Service.Approve(opts, v1, 1))
}

func TestGetApprove(t *testing.T) {
	fmt.Println(ddc721Service.GetApprove(3))
}

func TestIsApprovedForAll(t *testing.T) {
	fmt.Println(ddc721Service.IsApprovedForAll(platform, v1))
}

func TestTransferFrom(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.TransferFrom(opts, platform, v2, 5))
}

func TestSafeTransferFrom(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.SafeTransferFrom(opts, platform, v2, 5, []byte{}))
}

func TestOwnerOf(t *testing.T) {
	fmt.Println(ddc721Service.OwnerOf(1000))
}

func TestDDCURI(t *testing.T) {
	fmt.Println(ddc721Service.DdcURI(2))
}

func TestSymbol(t *testing.T) {
	fmt.Println(ddc721Service.Symbol())
}
