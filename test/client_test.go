package test

import (
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	clientBuilder = app.DDCSdkClientBuilder{}
	client        = clientBuilder.
			SetSignEventListener(new(SignListener)). //注册实现了签名接口的结构体
			SetGasPrice(1).                          //设置gasPrice
			SetGatewayURL("https://opbningxia.bsngate.com:18602/api/c23867c9099d411880c52fc9284da8ad/evmrpc").
			SetGatewayAPIKey("x-api-key").
			SetGatewayAPIValue("00194cd7affd494a9aa1a2f60fa5da42").
			SetAuthorityAddress("0xFa1d2d3EEd20C4E4F5b927D9730d9F4D56314B29").
			SetChargeAddress("0x0B8ae0e1b4a4Eb0a0740A250220eE3642d92dc4D").
			SetDDC721Address("0x354c6aF2cB870BEFEA8Ea0284C76e4A46B8F2870").
			SetNft("0x11e81E617d656a8D17aAcC1dD68d9B3Fe9E03888").
			SetDDC1155Address("0x0E762F4D11439B1130D402995328b634cB9c9973").
			RegisterLog("./log.log"). //设置日志输出的文件路径
			Build()

	opts = new(bind.TransactOpts)

	platform    = "0x2E10C44D7bE60AbB8BEd88B0BA4103E851Bbef2b" //ddcID：2、4、6
	platformPri = "0xe206069e3ce17a6d02c12bc3d00643165e60a0ee239f8bf6c069cd71169865fd"
	//无gas
	v1    = "0x2E10C44D7bE60AbB8BEd88B0BA4103E851Bbef2b"
	v1Pri = "0xe206069e3ce17a6d02c12bc3d00643165e60a0ee239f8bf6c069cd71169865fd"
	//无gas
	v2    = "0x2E10C44D7bE60AbB8BEd88B0BA4103E851Bbef2b" //ddcID:5
	v2Pri = "0xe206069e3ce17a6d02c12bc3d00643165e60a0ee239f8bf6c069cd71169865fd"
	//账户和私钥的映射
	senderPri = map[string]string{
		platform: platformPri,
		v1:       v1Pri,
		v2:       v2Pri,
	}
)
