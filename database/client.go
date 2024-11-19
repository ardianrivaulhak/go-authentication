package database
import (
  "log"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "go-authentication/models"
)

var Instance *gorm.DB
var dbError error
func Connect(connectionString string) () {
  
  Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
  if dbError != nil {
    log.Fatal(dbError)
    panic("Cannot connect to DB")
  }
  log.Println("Connected to Database!")
}
func Migrate() {
  Instance.AutoMigrate(&models.User{}, &models.Categories{})
  log.Println("Database Migration Completed!")
}