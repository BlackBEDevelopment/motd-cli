/*
 * @Author: NyanCatda
 * @Date: 2022-11-19 04:08:10
 * @LastEditTime: 2022-11-19 04:09:15
 * @LastEditors: NyanCatda
 * @Description: 分割服务器地址与端口
 * @FilePath: \motd-cli\tools\ServerPort\Split.go
 */
package ServerPort

import (
	"strconv"
	"strings"
)

/**
 * @description: 分割服务器地址与端口
 * @param {string} Address 服务器地址 
 * @return {string} 服务器地址
 * @return {int} 服务器端口
 */
func Split(Address string) (string, int) {
	// 判断是否有端口
	AddressArr := strings.Split(Address, ":")

	if len(AddressArr) == 1 {
		return Address, 0
	} else {
		Port, err := strconv.Atoi(AddressArr[1])
		if err != nil {
			return Address, 0
		}

		return AddressArr[0], Port
	}
}
