package models

import (
	"errors"
	"time"
	"ultimate_timer/services"
)

type Preset struct {
	BaseModel
	Name         string `db:"name" json:"name"`
	DisplayOrder int `db:"display_order" json:"display_order"`
	WaitsConfirm bool   `db:"waits_confirm" json:"waits_confirm"`
	LoopCount    int    `db:"loop_count" json:"loop_count"`
	Unit         []Unit
}

type Unit struct {
	Duration time.Duration `db:"duration" json:"duration"`
	PresetID string        `db:"preset_id" json:"preset_id"`
}

/*
constructor is for preparing objects (returns struct and error)
setter is for updating new objects (returns error if failed to udpate)
*/

// constructor
func NewPreset(
	name, memo string,
	price int,
	purchaseDate time.Time) (*Preset, error) {

	if name == "" {
		return nil, errors.New("項目名を入力してください")
	}

	now := time.Now()
	id := lib.GenUuid()

	preset := &Preset{
		BaseModel: BaseModel{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:         name,
		Memo:         memo,
		Price:        price,
		PurchaseDate: purchaseDate,
		// SmallCategoryId: smallCategoryId,
		// UserId:          userId,
	}

	return item, nil
}

// setter
func (p *Preset) Set(
	name, memo string,
	price int,
	purchaseDate time.Time) error {

	if name == "" {
		return errors.New("項目名を入力してください")
	}

	now := time.Now()

	i.UpdatedAt = now
	i.Name = name
	i.Memo = memo
	i.Price = price
	i.PurchaseDate = purchaseDate
	// i.SmallCategoryId = smallCategoryId

	return nil
}
