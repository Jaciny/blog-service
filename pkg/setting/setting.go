package setting

import "github.com/spf13/viper"

//自定义结构体，继承Viper配置实体
type Setting struct {
	vp *viper.Viper
}

//此处加载配置文件，将文件读取到vp结构体中
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
