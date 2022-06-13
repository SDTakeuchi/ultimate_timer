package infra

import (
	"reflect"
	"testing"
	"ultimate_timer/domain/model"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestTimerPresetRepository_Create(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	TimerPresetRepository{DB: db}

	id := "e6ee385e-4ef0-4eb6-ac3f-d6dbbbd503cd"
	createdAt := "2022-01-01 12:34:56"
	updatedAt := "2022-01-01 12:34:56"
	name := "testable name"
	displayOrder := 1
	loopCount := 22
	waitsConfirmEach := true
	waitsConfirmLast := false
	timerUnitDuration1 := 1234
	timerUnitOrder1 := 333

	// Mock設定
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "timerPresets" (
				"id",
				"created_at",
				"updated_at",
				"name",
				"display_order",
				"loop_count",
				"waits_confirm_each",
				"waits_confirm_last"
				) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
			RETURNING "users"."id"`)).
		WithArgs(
			id,
			createdAt,
			updatedAt,
			name,
			displayOrder,
			loopCount,
			waitsConfirmEach,
			waitsConfirmLast,
			timerUnitDuration1,
			timerUnitOrder1,
		).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id))

	type fields struct {
		Conn *gorm.DB
	}
	type args struct {
		timerPreset *model.TimerPreset
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.TimerPreset
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &TimerPresetRepository{
				Conn: tt.fields.Conn,
			}
			got, err := pr.Create(tt.args.timerPreset)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimerPresetRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimerPresetRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
