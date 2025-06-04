package gorm

// https://gorm.io/zh_CN/docs/transactions.html
import (
	"errors"

	"gorm.io/gorm"
)

func Transaction() {
	t := teacherTemp
	t1 := teacherTemp

	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}

		if err := tx.Create(&t1).Error; err != nil {
			return err
		}

		return nil
	})

}

func NestTransaction() {

	t := teacherTemp
	t1 := teacherTemp
	t2 := teacherTemp
	t3 := teacherTemp

	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}

		// 回滚子事务，不影响大事务最终结果
		tx.Transaction(func(tx1 *gorm.DB) error {
			tx1.Create(&t1)
			return errors.New("rollback t1")
		})

		tx.Transaction(func(tx2 *gorm.DB) error {
			if err := tx2.Create(&t2).Error; err != nil {
				return err
			}
			return nil
		})

		if err := tx.Create(&t3).Error; err != nil {
			return err
		}

		return nil
	})

}
