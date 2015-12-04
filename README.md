[![Build Status](https://travis-ci.org/zheng-ji/gophone.svg)](https://travis-ci.org/zheng-ji/gophone) [![GoDoc](https://godoc.org/github.com/zheng-ji/gophone?status.svg)](https://godoc.org/github.com/zheng-ji/gophone)

## gophone

获取手机号码的归属地，邮编，卡的类型，区号，运营商等属性

Get PhoneNum's property,
Such as Province, City, ZipCode, CardType, AreaZone, OperatorCompany


### Complie

```
go get github.com/zheng-ji/gophone
```

### Example

```go
import (
	"fmt"
	"github.com/zheng-ji/gophone"
)

func main() {

	pr, err := gophone.Find("15920544678")
	if err == nil {
		fmt.Println(pr)
	}

	pr, err = gophone.Find("15920544678")
	if err == nil {
		// 也可以单独获取该号码各个属性
		fmt.Println(pr.PhoneNum)
		fmt.Println(pr.Province)
		fmt.Println(pr.AreaZone)
		fmt.Println(pr.City)
		fmt.Println(pr.ZipCode)
	}
}
```

### OutPut

```
PhoneNum: 15920544678
AreaZone: 020
CardType: 移动虚拟运营商
City: 广州
ZipCode: 510000
Province: 广东
```

License
-------

Copyright (c) 2015 by [zheng-ji](http://zheng-ji.info) released under a MIT style license.

Thanks To the Data Provided by [@loved](https://github.com/lovedboy)

