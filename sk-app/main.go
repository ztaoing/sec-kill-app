/**
* @Author:zhoutao
* @Date:2020/7/6 下午2:34
 */

package main

import (
	"github.com/ztaoing/sec-kill-app/sk-app/setup"
	"github.com/ztaoing/sec-kill-pkg/pkg/bootstrap"
)

func main() {
	setup.InitZk()
	setup.InitRedis()
	setup.InitHTTP(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
}
