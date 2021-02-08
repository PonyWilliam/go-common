package common

import "encoding/json"

//工具类，主要用于关系映射
func SwapTo(request,some interface{})(err error){
	dataByte,err := json.Marshal(request)
	if err!=nil{
		return
	}
	err = json.Unmarshal(dataByte,some)
	return
}