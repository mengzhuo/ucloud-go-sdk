# ucloud-go-sdk
Ucloud API 的 go语言SDK


## 安装
    go get gitcafe.com/mengzhuo/ucloud-go-sdk
    // Or..
    go get github.com/mengzhuo/ucloud-go-sdk

## 使用例子
    
    import (
        ucloud "gitcafe.com/mengzhuo/ucloud-go-sdk"
        )

    u := ucloud.NewUcloudApiClient("https://api.ucloud.cn", publicKey, privateKey, "1", "1")
    response, err := u.Do(&ucloud.SendSms{"Hello World", []string{"10086"}})
    if err != nil || !response.OK() {
        // 一些错误处理
    }
    // 业务代码……

## Response处理
由于Golang是静态语言，仅能通过interface的方式保证接口的统一
因此，如需要对请求的数据进行读取，便需要type assertion
例如：
    
	rsp, err := u.Do(&DescribeUHostInstance{Region: "cn-east-01"})
	if !rsp.OK() || err != nil {
        // Bla……
	}
	fmt.Println(rsp.Data().([]*UHostSet)[0])

## Request种类
目前包括了UHost, UMon, UCDN, UDB, ULB, UNet, Sms接口
请求方式与[Ucloud API文档](http://docs.ucloud.cn/api/apilist.html)一致，这里就不再赘述了
