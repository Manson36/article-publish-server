package practice_gorm

import "github.com/jinzhu/gorm"

type User struct {
	//Default table name is `users`
	Role string
	db   *gorm.DB
}

//Set User's table name to be `profile`
func (User) TableName() string {
	return "profile"
}

func (u User) TableName2() string {
	if u.Role == "admin" {
		return "admin_users"
	} else {
		return "users"
	}
}

//Disable table name's pluralization, if set to true, `User`'s table name will be `user`
func (u User) TableName3() {
	u.db.SingularTable(true)
}

//change default table name
func TableName4() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}
}
