/*
 * @Author: NyanCatda
 * @Date: 2022-11-19 03:14:05
 * @LastEditTime: 2022-11-19 04:59:40
 * @LastEditors: NyanCatda
 * @Description: 主文件
 * @FilePath: \motd-cli\main.go
 */
package main

import (
	"fmt"

	"github.com/BlackBEDevelopment/motd-cli/internal/Parameter"
	"github.com/BlackBEDevelopment/motd-cli/internal/PrintMOTD"
)

func main() {
	// 获取参数
	Parameter, err := Parameter.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 判断版本参数
	switch Parameter.Version {
	case "be":
		PrintMOTD.GetBEServe(Parameter.Address)
		return
	case "je":
		PrintMOTD.GetJEServe(Parameter.Address)
		return
	default:
		// 如果没有指定版本，则使用推断模式
		PrintMOTD.Guess(Parameter.Address)
	}
}
