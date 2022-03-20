package main

import (
	"ultimate_timer/config"
	"ultimate_timer/handler"
	"ultimate_timer/infra"
	"ultimate_timer/services"
	"ultimate_timer/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// "errors"

	// "encoding/json"
	// "ultimate_timer/services"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// func main() {
// 	e := echo.New()

// 	defer db.Close()
// 	// Middleware
// 	// e.Use(services.Logger)
// 	// e.Use(middleware.Recover())

// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.GET("/preset/:id", FindPresetByID)
// 	e.POST("/preset", CreatePreset)
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// var (
// 	db  *gorm.DB
// 	err error
// )

// func init() {
// 	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("db connected: ", &db)
// 	db.Set("gorm:table_options", "ENGINE=InnoDB")
// 	db.AutoMigrate(&Preset{})
// 	db.LogMode(true)
// }


func main() {
    taskRepository := infra.NewPresetRepository(config.NewDB())
    taskUsecase := usecase.NewPresetUsecase(taskRepository)
    taskHandler := handler.NewPresetHandler(taskUsecase)

    e := echo.New()
	// Middleware
	e.Use(services.Logger)
	e.Use(middleware.Recover())
    handler.InitRouting(e, taskHandler)
    e.Logger.Fatal(e.Start(":8080"))
}

// type Preset struct {
// 	gorm.Model
// 	Name             string      `db:"name" json:"name"`
// 	DisplayOrder     int         `db:"display_order" json:"display_order"`
// 	LoopCount        int         `db:"loop_count" json:"loop_count"`
// 	WaitsConfirmEach bool        `db:"waits_confirm_each" json:"waits_confirm_each"`
// 	WaitsConfirmLast bool        `db:"waits_confirm_last" json:"waits_confirm_last"`
// 	TimerUnits       []TimerUnit `db:"timer_unit" json:"timer_unit"`
// }

// type TimerUnit struct {
// 	Order    int           `json:"order"`
// 	Duration int `json:"duration"`
// 	PresetID uint          `json:"-"`			//hides in json response
// }

// func CreatePreset(c echo.Context) error {
// 	preset := new(Preset)
// 	if err := c.Bind(preset); err != nil {
// 		return err
// 	}
// 	db.NewRecord(preset) // just checks if it IS a NEW RECORD (primary key is not duplicated)
// 	db.Create(&preset)
// 	return c.JSON(http.StatusOK, preset)
// }

// func FindPresetByID(c echo.Context) error {
// 	preset := Preset{}
// 	id := c.Param("id")
// 	db.First(&preset, id).Related(&preset.TimerUnits)
// 	// db.First(&preset, id).Model(&preset).Related(&preset.TimerUnits)  works as well, WHY?
// 	return c.JSON(http.StatusOK, preset)
// }
