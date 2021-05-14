package common

import "encoding/json"

//通过类型转换
func SwapTo(request, category interface{}) (err error)  {
	//转换成切片
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	//转换成字符串
	err = json.Unmarshal(dataByte, category)
	return
}

