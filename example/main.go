package main

import (
	"github.com/whuanle/lsm"
	"github.com/whuanle/lsm/config"
	"github.com/whuanle/lsm/kv"
)

type TestValue struct {
	A int64
	B int64
	C int64
	D string
}

func main() {
	lsm.Start(config.Config{
		DataDir:    `E:\项目\lsm数据测试目录`,
		Level0Size: 1,
		PartSize:   4,
		Threshold:  500,
	})
	lsm.Set("a", "tes")
	testV := TestValue{
		A: 1,
		B: 1,
		C: 3,
		D: "00000000000000000000000000000000000000",
	}
	data, _ := kv.Convert(testV)
	value := kv.Value{
		Key:     "abcdef",
		Value:   data,
		Deleted: false,
	}
	lsm.Set("b", value)
	for {

	}
	//
	//testV := TestValue{
	//	A: 1,
	//	B: 1,
	//	C: 3,
	//	D: "00000000000000000000000000000000000000",
	//}
	//data, _ := kv.Convert(testV)
	//value := kv.Value{
	//	Key:     "abcdef",
	//	Value:   data,
	//	Deleted: false,
	//}
	//js, _ := json.Marshal(value)
	//fmt.Println(len(js))
}
