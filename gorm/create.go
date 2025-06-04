package gorm

import (
	"log"
	"time"
)

// https://gorm.io/zh_CN/docs/create.html

var teacherTemp = Teacher{
	Name:         "nick",
	Age:          40,
	WorkingYears: 10,
	Email:        "nick@ovoice.com",
	Birthday:     time.Now().Unix(),

	StuNumber: struct {
		String string
		Valid  bool
	}{String: "10", Valid: true},

	Roles: []string{"普通用户", "讲师"},

	JobInfo: Job{
		Title:    "讲师",
		Location: "湖南长沙",
	},

	JobInfo2: Job{
		Title:    "讲师",
		Location: "湖南长沙",
	},
}

func CreateRecord() {
	t := teacherTemp
	res := DB.Create(&t)

	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	Println("CreateRecord():", res.RowsAffected, res.Error, t)

	// 正向选择
	t1 := teacherTemp
	res = DB.Select("Name", "Age").Create(&t1)
	Println("CreateRecord():", res.RowsAffected, res.Error, t1)

	// 反向选择
	t2 := teacherTemp
	res = DB.Select("Email", "Birthday").Create(&t2)
	Println("CreateRecord():", res.RowsAffected, res.Error, t2)

	// 批量操作 //docs/images/db-CreateInBatches.png
	var teachers1 = []Teacher{{Name: "king", Age: 40},
		{Name: "daren", Age: 40}, {Name: "nick", Age: 40},
	}
	DB.CreateInBatches(teachers1, 2)
	for _, tc := range teachers1 {
		Println(tc.ID)
	}

}
