package common

import (
	"log"
	"os"
	"strings"

	"github.com/Unknwon/goconfig"
)

//配置文件
var HelpConfigFile, SystemConfigFile *goconfig.ConfigFile

//初始化配置文件
func InitConfig(path string) error {
	//加载帮助文件
	helpConfigFile, err := goconfig.LoadFromReader(strings.NewReader(CONFIG_HELP_CONTENT_EN))
	if err != nil {
		log.Println("加载帮助文件失败！", err)
		return err
	}
	//加载系统配置文件
	systemConfigFile, err := goconfig.LoadFromReader(strings.NewReader(CONFIG_SYSTEM_DEFAULT))
	if err != nil {
		log.Println("加载默认配置文件失败！", err)
		return err
	}
	//如果没有指定配置文件，则加载默认配置文件
	if len(strings.TrimSpace(path)) == 0 {
		file, err := os.Open(CONFIG_SYSTEM_FILE)
		defer file.Close()
		//默认配置文件不存在，则创建
		if err != nil && os.IsNotExist(err) {
			log.Println("配置文件文件不存在，创建：", CONFIG_SYSTEM_FILE)
			goconfig.SaveConfigFile(systemConfigFile, CONFIG_SYSTEM_FILE)
		} else {
			path = CONFIG_SYSTEM_FILE
		}
	}
	//如果用户指定了配置文件，或存在默认配置，则覆盖系统预设配置文件
	if len(strings.TrimSpace(path)) > 0 {
		log.Println("加载系统配置文件，文件路径：", path)
		userConfigFile, err := goconfig.LoadConfigFile(path)
		if err != nil {
			log.Println("加配置文件失败！", err)
			return err
		}
		for _, sec := range systemConfigFile.GetSectionList() {
			m, _ := systemConfigFile.GetSection(sec)
			for k, _ := range m {
				if nv, err := userConfigFile.GetValue(sec, k); err == nil {
					systemConfigFile.SetValue(sec, k, nv)
				}
			}
		}
		log.Println("加载系统配置文件完成")
	}
	HelpConfigFile = helpConfigFile
	SystemConfigFile = systemConfigFile
	return nil
}

//获取帮助配置
func GetHelpConfig() *goconfig.ConfigFile {
	return HelpConfigFile
}

//获取系统配置
func GetSystemConfig() *goconfig.ConfigFile {
	return SystemConfigFile
}
