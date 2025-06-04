package gorm

func Query() {
	//-------查询单条数据----------

	// 查询第一条
	t := Teacher{}
	res := DB.Table("teachers").First(&t)
	Println(res.RowsAffected, res.Error, t)

	// 查询最后一条
	t = Teacher{}
	res = DB.Last(&t)
	Println(res.RowsAffected, res.Error, t)

	// 无排序，取第一条
	t = Teacher{}
	res = DB.Take(&t)
	Println(res.RowsAffected, res.Error, t)

	// works because model is specified using `db.Model()`
	//
	// 将结果填充到集合，不支持 **serializer** 特殊类型处理，无法完成类型转换,
	// 2025/06/03 13:57:14 go0base/gorm/query.go:23 sql: Scan error on column index 8,
	//    name "birthday": converting driver.Value type time.Time
	//    ("2025-06-02 23:32:44 +0800 CST") to a int64: invalid syntax
	//    [0.176ms] [rows:1]
	//    SELECT * FROM `teachers` WHERE `teachers`.`deleted_at` IS NULL ORDER BY `teachers`.`id` LIMIT 1
	// 这类字段需要omit
	result := map[string]interface{}{}
	DB.Model(&Teacher{}).Omit("Birthday", "Roles", "JobInfo2").First(&result)
	Println(res.RowsAffected, res.Error, result)

	// 基于表名，不支持 First、Last
	result = map[string]interface{}{}
	DB.Table("teachers").Take(&result)
	Println(res.RowsAffected, res.Error, result)

	//-------- 查询多条记录 --------
	var teachers []Teacher
	res = DB.Where("name=?", "nick").Or("name=?",
		"king").Order("id desc").Limit(10).Find(&teachers)
	Println(res.RowsAffected, res.Error, teachers)
}
