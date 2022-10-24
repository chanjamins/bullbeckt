package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
加载配置文件
*/
var Conf = new(AppConfig)

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mdoe"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    string `mapstructure:"maxsize"`
	MaxAge     string `mapstructure:"max_age"`
	MaxBackups string `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init(filePath string) error {

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("初始化配置文件失败，，，")
		return err
	}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("序列化失败。。。")
		return err
	}
	//动态更新
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了...")
		//重新序列化
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("重新序列化失败...")
			return
		}
	})
	return nil
}
