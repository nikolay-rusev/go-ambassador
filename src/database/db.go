package database

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var DB *gorm.DB

func Connect() {
    DB, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})

    if err != nil {
        panic("Could not connect with the database!")
    }
}

func AutoMigrate() {
    DB.AutoMigrate(models.User{})
}
