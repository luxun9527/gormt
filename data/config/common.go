package config

import (
	"fmt"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/tools"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// CfgBase base config struct
type CfgBase struct {
	// SerialNumber       string `json:"serial_number" yaml:"serial_number"`             // version.版本号
	// ServiceName        string `json:"service_name" yaml:"service_name"`               // service name .service名字
	// ServiceDisplayname string `json:"service_displayname" yaml:"service_displayname"` // display name .显示名
	// SerciceDesc        string `json:"sercice_desc" yaml:"sercice_desc"`               // sercice desc .service描述
	IsDev bool `json:"is_dev" yaml:"is_dev"` // Is it a development version?是否是开发版本
}

var _map = Config{
	CfgBase: CfgBase{
		IsDev: false,
	},
	DBInfo: DBInfo{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "test",
	},
	OutDir:        "./model",
	URLTag:        "json",
	Language:      "中 文",
	DbTag:         "gorm",
	UserGormModel: false,
	Simple:        false,
	IsWEBTag:      false,
	// SingularTable: true,
	IsForeignKey:         true,
	IsOutSQL:             false,
	IsOutFunc:            true,
	IsGUI:                false,
	TablePrefix:          "",
	SelfTypeDef:          make(map[string]string),
	WebTagType:           0,
	TableNames:           "",
	IsColumnName:         true,
	IsOutFileByTableName: false,
}

var ConfigPath string

func init() {
	//ConfigPath = path.Join(tools.GetCurrentDirectory(), "config.yml") // 先找本程序文件夹
	//if !tools.CheckFileIsExist(ConfigPath) {                          // dont find it
	//	ConfigPath = path.Join(tools.GetModelPath(), "config.yml")
	//	if !tools.CheckFileIsExist(ConfigPath) {
	//		mylog.ErrorString("config.yml not exit. using default config")
	//	}
	//}

	onInit()
	dev.OnSetDev(_map.IsDev)
}

func onInit() {
	err := InitFile(ConfigPath)
	if err != nil {
		fmt.Println("Load config file error: ", err.Error())
		return
	}
}

// InitFile default value from file .
func InitFile(filename string) error {
	// if _, e := os.Stat(filename); e != nil {
	// 	fmt.Println("init default config file: ", filename)
	// 	if err := SaveToFile(); err == nil {
	// 		InitFile(filename)
	// 		return nil
	// 	} else {
	// 		fmt.Println("shit,fail", err)
	// 	}
	// 	// os.Exit(0)
	// }
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bs, &_map); err != nil {
		fmt.Println("read config file error: ", err.Error())
		return err
	}
	return nil
}

// GetServiceConfig Get service configuration information
// func GetServiceConfig() (name, displayName, desc string) {
// 	name = _map.ServiceName
// 	displayName = _map.ServiceDisplayname
// 	desc = _map.SerciceDesc
// 	return
// }

// GetIsDev is is dev
func GetIsDev() bool {
	return _map.IsDev
}

// SetIsDev is is dev
func SetIsDev(b bool) {
	_map.IsDev = b
}

// SaveToFile save config info to file
func SaveToFile() error {
	d, err := yaml.Marshal(_map)
	if err != nil {
		return err
	}
	tools.WriteFile(ConfigPath, []string{
		string(d),
	}, true)
	return nil
}
