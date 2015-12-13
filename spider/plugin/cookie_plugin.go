package plugin

import (
	"github.com/zhangxiaoyang/goDataAccess/spider/common"
	"net/http/cookiejar"
)

type GetCookieFunc func(*common.Request) *cookiejar.Jar

type CookiePlugin struct {
	getCookieFunc GetCookieFunc
}

func NewCookiePlugin(getCookieFunc GetCookieFunc) *CookiePlugin {
	return &CookiePlugin{getCookieFunc: getCookieFunc}
}

func (this *CookiePlugin) Do(pluginType PluginType, args ...interface{}) {
	if pluginType == PreDownloaderType {
		req := args[0].(*common.Request)
		req.Jar = this.getCookieFunc(req)
	}
}
