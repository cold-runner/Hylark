package main

import (
	"os"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/rawsql"
)

var (
	SQL_FILE_NAME = os.Getenv("SQL_FILE_NAME")
	SERVICE_NAME  = os.Getenv("SERVICE_NAME")
)

func main() {
	// 构造生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:           "gorm_gen/" + SERVICE_NAME,
		ModelPkgPath:      "gorm_gen/" + SERVICE_NAME + "/model",
		Mode:              gen.WithQueryInterface | gen.WithoutContext | gen.WithDefaultQuery,
		FieldNullable:     false,
		FieldCoverable:    true,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	db, _ := gorm.Open(rawsql.New(rawsql.Config{
		//SQL:      rawsql,                      // 建表sql
		FilePath: []string{
			SQL_FILE_NAME, // 建表sql文件
			//"./test_sql", // 建表sql目录
		},
	}))
	g.UseDB(db)

	//generate all table from database
	g.ApplyBasic(g.GenerateAllTable(
		gen.FieldGORMTagReg("^id$", func(tag field.GormTag) field.GormTag {
			return tag.
				Append("default", "uuid_generate_v4()").
				Set("type", "uuid")
		}),
		gen.FieldType("id", "uuid.UUID"),
	)...)

	g.Execute()
}
