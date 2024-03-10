package database

import (
    "Seinarukiro_XYZ/main"
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "path/filepath"
    "runtime"
    "sync"
    "time"
)

var (
    db          *gorm.DB
    initialized bool
    initMu      sync.Mutex
)

type About struct {
    ID      uint   `gorm:"primaryKey"`
    About   string `gorm:"type:text"`
    Contact string `gorm:"type:text"`
}

type Post struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"not null"`
    Date        time.Time `gorm:"not null"`
    Description string
    Image       string
    Content     string    `gorm:"type:text"`
    Filename    string    `gorm:"uniqueIndex"`
    CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
    Draft       bool      `gorm:"not null;default:false"`
}

func ConnectToDatabase() (*gorm.DB, error) {
    _, filename, _, _ := runtime.Caller(0)
    dbPath := filepath.Join(filepath.Dir(filename), "seinaru.db")
    return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}

func InitializeDatabase() error {
    initMu.Lock()
    defer initMu.Unlock()
    if initialized {
        return nil
    }

    var err error
    db, err = ConnectToDatabase()
    if err != nil {
        return err
    }

    db.AutoMigrate(&About{}, &Post{})

    db.Callback().Create().Before("gorm:creating").Register("setCreatedAt", func(db *gorm.DB) {
        db.Statement.SetColumn("CreatedAt", time.Now())
    })

    db.Callback().Update().After("gorm:after_update").Register("logPostUpdate", func(db *gorm.DB) {
        fmt.Println("Post updated!")
    })

    initialized = true
    return nil
}

func GetLatestPost() (Post, error) {
    var latestPost Post
    if err := db.Order("created_at desc").First(&latestPost).Error; err != nil {
        return Post{}, err
    }
    return latestPost, nil
}

func GetDB() *gorm.DB {
    return db
}

func init() {
    if err := InitializeDatabase(); err != nil {
        panic("failed to initialize database: " + err.Error())
    }
}
