package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app/handler"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/types"
)

type NftService struct {
	Base
}

func NewDDCNftService() *NftService {
	return &NftService{}
}

// Mint
// @Description: 平台方或终端用户可以通过调用该方法进行DDC的生成
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return erro

func (d *NftService) MyMint(opts *bind.TransactOpts, to, ddcURI string) (signedTx *types.Transaction, err error) {
	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDCMyNft().AwardItem(opts, common.HexToAddress(to), ddcURI)
	if err != nil {
		log.Error.Printf("failed to execute SafeMint: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}
