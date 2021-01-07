package req

import (
	"LollipopGo2.8x/conf/g"
	"fmt"
	"strings"
)

//获取url路径
func getUrlPath(tpType g.TransportType, strArr []string) string {
	urlPath := strings.Join(strArr, "")
	return fmt.Sprintf("%s://%s", tpType, urlPath)
}



