package scan

import (
	. "getitle/src/fingers"
	. "getitle/src/pkg"
)

// -v
// 信息收集插件,通过匹配http服务的favicon md5值判断CMS
func faviconScan(result *Result) {
	var err error
	conn := HttpConn(RunOpt.Delay)
	url := result.GetURL()
	resp, err := conn.Get(url + "/favicon.ico")
	if err != nil {
		Log.Debugf("request favicon %s/favicon.ico %s", result.GetURL(), err.Error())
		return
	}
	Log.Debugf("request favicon %s/favicon.ico %d", result.GetURL(), resp.StatusCode)
	if resp.StatusCode != 200 {
		return
	}
	content := GetBody(resp)

	// MD5 hash匹配
	md5h := Md5Hash(content)
	if Md5Fingers[md5h] != "" {
		result.AddFramework(&Framework{Name: Md5Fingers[md5h], Version: "ico"})
		return
	}

	// mmh3 hash匹配,指纹来自kscan
	mmh3h := Mmh3Hash32(content)
	if Mmh3Fingers[mmh3h] != "" {
		result.AddFramework(&Framework{Name: Mmh3Fingers[mmh3h], Version: "ico"})
		return
	}
	return
}
