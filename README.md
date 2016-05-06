# ucloud-go-sdk
Ucloud API 的 go语言SDK

![Build Status](https://travis-ci.org/mengzhuo/ucloud-go-sdk.svg)
[![Godoc Reference](https://godoc.org/github.com/mengzhuo/ucloud-go-sdk?status.png)](https://godoc.org/github.com/mengzhuo/ucloud-go-sdk)

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

    UHost
    UMon
    UCDN
    UDB
    ULB
    UNet
    Sms
    UDisk

请求参数与[Ucloud API文档](http://docs.ucloud.cn/api/apilist.html)一致

具体文档可见[Godoc](https://godoc.org/github.com/mengzhuo/ucloud-go-sdk)

## LICENSE
The MIT License

## 代码生成器
本代码部分由Python生成，具体见tools文件夹

[如何生成代码的博客文章](https://mengzhuo.org/%E5%A6%82%E4%BD%95%E4%B8%80%E4%B8%8B%E5%8D%88%E5%86%993000%E8%A1%8C%EF%BC%9F%E8%AE%B0%E6%9F%90%E4%BA%91%E7%9A%84Golang-API-SDK%E7%94%9F%E4%BA%A7%E8%BF%87%E7%A8%8B.html)
