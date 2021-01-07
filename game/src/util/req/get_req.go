package req

import (
	"LollipopGo2.8x/conf/g"
	"fmt"
	"strings"
)

//添加参数到get请求
func AddParamsToGetReq(tpType g.TransportType, strArr []string, paramsMap map[string]string) string {
	urlPath := getUrlPath(tpType, strArr)
	if len(paramsMap) <= 0 || paramsMap == nil { //如果没有参数需要添加直接返回当前路径
		return urlPath
	}
	urlPath = urlPath + "?" //如果参数个数大于等于0,路径后缀加上?
	paramList := make([]string, 0)
	for k, v := range paramsMap {
		paramList = append(paramList, fmt.Sprintf("%s=%s", k, v))
	}
	tempStr := strings.Join(paramList, "&")
	return fmt.Sprintf("%s%s", urlPath, tempStr)
}
