/*
 * @Author: NyanCatda
 * @Date: 2022-11-19 04:46:55
 * @LastEditTime: 2022-11-19 05:03:05
 * @LastEditors: NyanCatda
 * @Description: 推测版本并打印信息
 * @FilePath: \motd-cli\internal\PrintMOTD\Guess.go
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
 * @description: 尝试推断并打印服务器信息
 * @param {string} Address 服务器地址
 * @return {*}
 */
func Guess(Address string) {
	// 尝试以基岩版服务器模式获取信息
	BEAddress := ServerPort.Completion(Address, 19132)
	// 获取IP地址
	Host, Port := ServerPort.Split(BEAddress)
	IPAddress, err := net.ResolveIPAddr("ip", Host)
	if err != nil {
		fmt.Println("Unable to connect to the server, please check if the address is correct")
		return
	}

	PrintTryingTips := func(Host string, IP string, Port int) {
		if Host == IP {
			fmt.Printf("Trying to get MOTD info for %s\n", Host)
		} else {
			fmt.Printf("Trying to get MOTD info for %s [%s:%d]\n", Host, IP, Port)
		}
	}

	// 获取服务器信息
	BEMOTD, err := MotdBEAPI.MotdBE(BEAddress)
	if err == nil {
		if BEMOTD.Status != "offline" {
			// 基岩版服务器
			PrintTryingTips(BEAddress, IPAddress.String(), Port)

			PrintContent := fmt.Sprintf("MOTD info: Delay=%vms people=%d/%d agreement=%v game_version=%s level=%s mode=%s motd=%s", BEMOTD.Delay, BEMOTD.Online, BEMOTD.Max, BEMOTD.Agreement, BEMOTD.Version, BEMOTD.LevelName, BEMOTD.GameMode, strings.Replace(BEMOTD.Motd, "\n", "\\n", -1))
			fmt.Println(PrintContent)
			return
		}
	}

	// 尝试以Java版服务器模式获取信息
	JEAddress := ServerPort.Completion(Address, 25565)
	// 仅分割地址
	Host, Port = ServerPort.Split(JEAddress)
	// 获取服务器信息
	JEMOTD, err := MotdBEAPI.MotdJava(JEAddress)
	if err == nil {
		if JEMOTD.Status != "offline" {
			// Java版服务器
			PrintTryingTips(JEAddress, IPAddress.String(), Port)

			PrintContent := fmt.Sprintf("MOTD info: Delay=%vms people=%d/%d agreement=%v game_version=%s motd=%s", JEMOTD.Delay, JEMOTD.Online, JEMOTD.Max, JEMOTD.Agreement, JEMOTD.Version, strings.Replace(JEMOTD.Motd, "\n", "\\n", -1))
			fmt.Println(PrintContent)
			return
		}
	}

	fmt.Println("Unable to connect to the server, please check if the address is correct")
}
