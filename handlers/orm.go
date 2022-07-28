package handlers

import "gorm.io/gorm"

type callbackFunc func() error

func runWithTransaction(db *gorm.DB, callback callbackFunc) error {
	tran := db.Begin()
	defer tran.Commit()
	err := callback()
	return err
}

func Create(db *gorm.DB, v interface{}) error {
	return runWithTransaction(db, func() error {
		if err := db.Create(v).Error; err != nil {
			db.Rollback()
			return err
		}
		return nil
	})
}

func SaveChanges(db *gorm.DB, v interface{}) error {
	return runWithTransaction(db, func() error {
		if err := db.Save(v).Error; err != nil {
			db.Rollback()
			return err
		}
		return nil
	})
}

func Delete(db *gorm.DB, v interface{}) error {
	return runWithTransaction(db, func() error {
		if err := db.Delete(v).Error; err != nil {
			db.Rollback()
			return err
		}
		return nil
	})
}

func GetById(db *gorm.DB, v interface{}) error {
	return runWithTransaction(db, func() error {
		if err := db.Last(v, "id").Error; err != nil {
			db.Rollback()
			return err
		}
		return nil
	})
}

func GetAll(db *gorm.DB, v interface{}) error {
	return runWithTransaction(db, func() error {
		if err := db.Find(v).Error; err != nil {
			db.Rollback()
			return err
		}
		return nil
	})
}

func FindByQueryfunc(db *gorm.DB, v interface{}, params map[string]interface{}) error {
	return runWithTransaction(db, func() error {
		if err := db.Where(params).Find(v).Error; err != nil {
			db.Rollback()
			return err
		}
		return nil
	})
}
