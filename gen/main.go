package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

const MySQLDSN = "root:@Waymon102092!@tcp(121.41.59.235:3306)/sainiao_wann?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./gen/gen_query",
		ModelPkgPath: "./gen_model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 设置 JSON 标签命名策略为小驼峰
	g.WithJSONTagNameStrategy(func(columnName string) string {
		return toCamelCase(columnName)
	})

	// 替换时间类型为自定义时间类型
	g.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		"datetime": func(columnType gorm.ColumnType) string {
			return "local_time.Time"
		},
	})

	gormdb, _ := gorm.Open(mysql.Open(MySQLDSN))
	g.UseDB(gormdb) // 复用现有的 gorm db

	g.ApplyBasic(
		// 根据表 `wx_profession` 生成 struct `WxProfession`
		g.GenerateModel("wx_profession"),
		g.GenerateModel("wx_user_star"),
		g.GenerateModel("wx_user_infos"),
		g.GenerateModel("wx_user_info_details"),
		g.GenerateModel("wx_user_info_details_audits"),
		g.GenerateModel("wx_user_labels"),
		g.GenerateModel("wx_user_label_users"),
		g.GenerateModel("wx_red_moons_auths"),
		g.GenerateModel("wx_command_process"),
		g.GenerateModel("wx_command"),
		g.GenerateModel("wx_coupon_packages"),
		g.GenerateModel("wx_coupons_package"),
		g.GenerateModel("wx_coupons"),
		g.GenerateModel("wx_red_moons_user_info"),
		g.GenerateModel("sys_config"),
		g.GenerateModel("wx_red_moons"),
		g.GenerateModel("wx_red_moon_details"),
		g.GenerateModel("wx_user_logs"),
	)
	g.Execute()
}

// 工具函数：将下划线命名转换为小驼峰命名
func toCamelCase(input string) string {
	words := strings.Split(input, "_")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	result := strings.Join(words, "")
	if strings.HasSuffix(strings.ToLower(result), "id") {
		return result + ",string"
	}
	return result
}
