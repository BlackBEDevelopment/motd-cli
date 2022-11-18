/*
 * @Author: NyanCatda
 * @Date: 2022-11-19 03:16:27
 * @LastEditTime: 2022-11-19 04:45:25
 * @LastEditors: NyanCatda
 * @Description: 命令行参数获取
 * @FilePath: \motd-cli\internal\parameter\Get.go
 */
package Parameter

import (
	"errors"
	"os"
)

type Parameter struct {
	Version string // 版本，be, je
	Address string // 地址
}

/**
 * @description: 获取参数
 * @return {Parameter} 参数
 * @return {error} 错误
 */
func Get() (Parameter, error) {
	ConsoleParameter := os.Args

	if len(ConsoleParameter) < 2 {
		return Parameter{}, errors.New("Missing parameters")
	}

	if ConsoleParameter[1] == "be" || ConsoleParameter[1] == "je" {
		// 如果有指定版本，则直接返回
		return Parameter{
			Version: ConsoleParameter[1],
			Address: ConsoleParameter[2],
		}, nil
	} else {
		// 如果没有指定版本，则使用推断模式
		return Parameter{
			Version: "",
			Address: ConsoleParameter[1],
		}, nil
	}
}
