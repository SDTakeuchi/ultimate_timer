package infra

import (
	"reflect"
	"testing"
	"ultimate_timer/domain/model"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestPresetRepository_Create(t *testing.T) {
	type fields struct {
		Conn *gorm.DB
	}
	type args struct {
		preset *model.Preset
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Preset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PresetRepository{
				Conn: tt.fields.Conn,
			}
			got, err := pr.Create(tt.args.preset)
			if (err != nil) != tt.wantErr {
				t.Errorf("PresetRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PresetRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
