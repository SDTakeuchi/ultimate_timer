package main

import (
	// "ultimate_timer/services"
	"net/http"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"


	// "errors"
	"time"
	"fmt"
	// "encoding/json"
	// "ultimate_timer/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
    e := echo.New()

	defer db.Close()
	// Middleware
	// e.Use(services.Logger)
	// e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	e.GET("/preset/:id", FindPresetByID)
	e.POST("/preset", CreatePreset)
    e.Logger.Fatal(e.Start(":8080"))
}



var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&Preset{})
	db.LogMode(true)
}

type Preset struct {
	gorm.Model
	Name             string `db:"name" json:"name"`
	DisplayOrder     int    `db:"display_order" json:"display_order"`
	LoopCount        int    `db:"loop_count" json:"loop_count"`
	WaitsConfirmEach bool   `db:"waits_confirm_each" json:"waits_confirm_each"`
	WaitsConfirmLast bool   `db:"waits_confirm_last" json:"waits_confirm_last"`
	TimerUnits        []TimerUnit `gorm:"ForeignKey:PresetID" db:"timer_unit" json:"timer_unit"`
}

type TimerUnit struct {
	Order int `json:"order"`
	Duration time.Duration `json:"duration"`
	PresetID uint
}


func CreatePreset(c echo.Context) error {
    preset := new(Preset)
	now := time.Now()
	preset.CreatedAt = now
	preset.UpdatedAt = now
    if err := c.Bind(preset); err != nil {
        return err
    }
	db.NewRecord(preset)
	db.Create(&preset)
    return c.JSON(http.StatusOK, preset)
}

func FindPresetByID(c echo.Context) error {
	preset := Preset{}
	id := c.Param("id")
	db.First(&preset, id).Related(&preset.TimerUnits)
	// db.First(&preset, id).Model(&preset).Related(&preset.TimerUnits)  works as well, WHY?
	return c.JSON(http.StatusOK, preset)
}
