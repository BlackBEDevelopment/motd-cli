/*
 * @Author: NyanCatda
 * @Date: 2022-11-19 03:31:30
 * @LastEditTime: 2022-11-19 04:43:42
 * @LastEditors: NyanCatda
 * @Description: 打印基岩版服务器状态
 * @FilePath: \motd-cli\internal\PrintMOTD\BedrockEdition.go
 */
package PrintMOTD

import (
	"fmt"
	"net"
	"strings"

	"github.com/BlackBEDevelopment/MCBE-Server-Motd/MotdBEAPI"
	"github.com/BlackBEDevelopment/motd-cli/tools/ServerPort"
)

/**
 * @description: 打印基岩版服务器状态
 * @param {string} Address 服务器地址
 * @return {*}
 */
func GetBEServe(Address string) {
	// 补全端口
	Address = ServerPort.Completion(Address, 19132)

	// 获取IP地址
	Host, Port := ServerPort.Split(Address)
	IPAddress, err := net.ResolveIPAddr("ip", Host)
	if err != nil {
		fmt.Println("Unable to connect to the server, please check if the address is correct")
		return
	}

	if Host == IPAddress.String() {
		fmt.Printf("Trying to get MOTD info for %s\n", Address)
	} else {
		fmt.Printf("Trying to get MOTD info for %s [%s:%d]\n", Address, IPAddress.String(), Port)
	}

	// 获取服务器信息
	MOTD, err := MotdBEAPI.MotdBE(Address)
	if err != nil {
		fmt.Println("Unable to connect to the server, please check if the address is correct")
		return
	}

	if MOTD.Status == "offline" {
		fmt.Println("Server is offline")
		return
	} else {
		PrintContent := fmt.Sprintf("MOTD info: Delay=%vms people=%d/%d agreement=%v game_version=%s level=%s mode=%s motd=%s", MOTD.Delay, MOTD.Online, MOTD.Max, MOTD.Agreement, MOTD.Version, MOTD.LevelName, MOTD.GameMode, strings.Replace(MOTD.Motd, "\n", "\\n", -1))
		fmt.Println(PrintContent)
		return
	}
}
