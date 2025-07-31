package pkg

import (
	"Music_api/appconfig"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

func Pay(outTradeNo, totalAmount string) (payURL string) {
	var privateKey = appconfig.Conf.ALiPay.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	client, err := alipay.New(appconfig.Conf.ALiPay.AppId, privateKey, false)
	if err != nil {
		logx.Error(err)
		return
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = appconfig.Conf.ALiPay.NotifyURL
	p.ReturnURL = appconfig.Conf.ALiPay.ReturnURL
	p.Subject = "支付宝沙箱支付业务"
	p.OutTradeNo = outTradeNo
	p.TotalAmount = totalAmount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		logx.Error(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	payURL = url.String()
	fmt.Println(payURL)
	return
}
