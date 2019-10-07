package pak

import (
	"github.com/Unknown/goconfig"
	"xorm"
)

/*goconfig
基于windows下的ini配置文件的解析器

提供和windows一样的操作方式
支持递归读取分区
支持自增键
支持对注释的读写
支持直接返回指定类型的键值
支持多个文件的覆盖加载
*/
func config() {
	cfg, err := goconfig.LoadConfigFile("conf.ini", "conf2.ini") //多文件加载, 后配置文件优先级更高

	err = cfg.AppendFiles("conf3.ini") //追加配置文件(优先级更高)

	val, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key") //读取 DEFAULT 区的内容

	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION, "key", "value") //写, 若为新键值对返回true

	ok := cfg.DeleteKey("must", "key") //删除指定键

	comment := cfg.GetSectionComments("super")         //获取分区注释, 读写注释都需要加 注释的标识符 # 或 ;
	isInsert = cfg.SetSectionComments("super", "# 注释") //写注释, 若为新键值对返回true

	vInt, err := cfg.Int("must", "int")          //在must分区获取key为int, 并强制转为int类型
	vBool := cfg.MustBool("must", "bool", false) //返回对应值并强转, 失败返回该类型默认值(或者返回缺省值)

	val, err = cfg.GetValue("parent.child", "key") //若子分区没有对应键值, 返回父分区的

	val = cfg.GetSection("must") //获取整个分区, 返回一个map

	isInsert = cfg.SetValue("auto", "-", "value") //设置自增键

	err = goconfig.SaveConfigFile(cfg, "conf.ini") //保存操作

}
