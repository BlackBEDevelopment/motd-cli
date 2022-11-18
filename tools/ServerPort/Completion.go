/*
 * @Author: NyanCatda
 * @Date: 2022-11-19 03:49:01
 * @LastEditTime: 2022-11-19 03:53:20
 * @LastEditors: NyanCatda
 * @Description: 补全服务器端口
 * @FilePath: \motd-cli\tools\ServerPort\Completion.go
 */
package ServerPort

import (
	"strconv"
	"strings"
)

/**
 * @description: 判断是否有端口，如果没有则自动补全
 * @param {string} Address 服务器地址
 * @param {int} Port 需要补全的端口
 * @return {string} 补全后的地址
 */
func Completion(Address string, Port int) string {
	// 判断是否有端口
	AddressArr := strings.Split(Address, ":")

	if len(AddressArr) == 1 {
		return Address + ":" + strconv.Itoa(Port)
	} else {
		return Address
	}
}
