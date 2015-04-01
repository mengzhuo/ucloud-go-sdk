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
    
    rsp, err := u.Do(&ucloud.DescribeUHostInstance{Region: "cn-east-01"})
    if err != nil {
        panic(err)
    }
    for _, h := range rsp.Data().([]*ucloud.UHostSet) {
        fmt.Println(h)
        for _, v := range h.DiskSet {
            fmt.Println(v)
        }
    }
    /* 返回
    &{uhost-ahdvfh Normal 71bf5434-9ae1-4b1e-ad6d-39cc814a1423 uimage-55j4fq UCloud Debian 6.0 64位 Default  uhost0 Running 1427473439 Trial 1427905439 2 4096 [0xc208129080 0xc208129260] [0xc20813c140]}
    &{Boot 71bf5434-9ae1-4b1e-ad6d-39cc814a1423 20}
    &{Data b9a0cb79-6d4c-4303-95a2-c97d7c5cf939 20}
    */

## Request种类
目前包括了

    Host
    UMon
    UCDN
    UDB
    ULB
    UNet
    Sms
    UDisk

请求方式与[Ucloud API文档](http://docs.ucloud.cn/api/apilist.html)一致，这里就不再赘述了

## LICENSE
The MIT License
