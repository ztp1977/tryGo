package main

import (
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"database/sql"
	"github.com/Sirupsen/logrus"
	"reflect"
	"fmt"
	"github.com/k0kubun/pp"
)

type User struct {
	gorm.Model
	Birthday     time.Time
	Age          int
	Name         string  `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	//Num          int     `gorm:"AUTO_INCREMENT"`

	CreditCard        CreditCard      // One-To-One relationship (has one - use CreditCard's UserID as foreign key)
	Emails            []Email         // One-To-Many relationship (has many - use Email's UserID as foreign key)

	BillingAddress    Address         // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
	BillingAddressID  sql.NullInt64

	ShippingAddress   Address         // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
	ShippingAddressID int

	IgnoreMe          int `gorm:"-"`   // Ignore this field
	//Languages         []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

type Email struct {
	ID      int
	UserID  int     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column
	Email   string  `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	gorm.Model
	UserID  uint
	Number  string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/storage?charset=utf8&loc=Local")

	logrus.SetLevel(logrus.DebugLevel)


	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	db.LogMode(true)
	defer db.Close()

	// create table
	// 構造体が事前に知っていれば、tableを作成できます。
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Email{})
	db.AutoMigrate(&User{})
	//db.DropTableIfExists(&User{}, "products")

	// index create
	//db.Model(&User{}).AddIndex("idx_user_name", "name")

	// insert rows
	//user := &User{
	//	Age: 10,
	//	Name: "test1",
	//}
	//db.Model(&User{}).Create(user)

	logrus.Debug(db.RowsAffected)

	users := []User{}
	db.Table("users").Find(&users)
	if db.GetErrors() != nil {
		logrus.Error(db.GetErrors())
		panic("error")
	}

	for _,v := range users {
		logrus.Debug(v.ID)
	}

	// rows
	rows, _ := db.Table("users").Select("? as fake, ? as fack", gorm.Expr("name"), gorm.Expr("name")).Rows()
	if !rows.Next() {
		logrus.Error("should have returned at least one row")
	} else {
		columns, _ := rows.Columns()
		if !reflect.DeepEqual(columns, []string{"fake"}) {
			logrus.Error("should have returned at least one row")
		}
		logrus.Debug(columns)
	}

	pp.Printf("%v", rows)

	for rows.Next()  {
		var fake string
		var fack string
		rows.Scan(&fake, &fack)
		logrus.Debug(fake, "----", fack)
	}
	rows.Close()

	// rawのSQLを投げる
	//db.Raw("select * from users where ID > ?", 1)
	res := []map[string]interface{}{}
	db.Raw("SELECT * FROM users where id > ?", 2).Scan(&res)
	logrus.Error(res)
	res2 := []map[string]string{}
	db.Raw("SELECT * FROM users where id > ?", 2).Scan(&res2)
	logrus.Error(res2)

	rows, _ = db.Raw("SELECT * FROM users where id > ?", 2).Rows()
	cols, _ := rows.Columns()

	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			pp.Printf("%v", col)
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(cols[i] , ": ", value)
		}
		fmt.Println("-----------------------------------")
	}

	tx, err = db.Begin()
	rows :=


	//// scopeの使い方
	//scope := &gorm.Scope{}
	//scope.Raw("select * from users Where id > 0").Exec()
	//logrus.Debug(scope.SQL)

	//scope.Exec()
	//logrus.Debug(scope)



}
