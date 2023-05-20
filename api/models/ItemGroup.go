package models 

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ItemGroup struct {
	ITEM_GROUP_ID uint64 `gorm:"primary_key;auto_increment" json:"item_group_id"`
	Item_Group_Name string `gorm:"size:255;not null;unique" json:"item_group_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ig *ItemGroup) Prepare()  {
	ig.ITEM_GROUP_ID = 0
	ig.Item_Group_Name = html.EscapeString((strings.TrimSpace(ig.Item_Group_Name)))
	ig.CreatedAt = time.Now()
	ig.UpdatedAt = time.Now()
}

func (ig *ItemGroup) Validate() error  {
	if ig.Item_Group_Name == "" {
		return errors.New("Required Item Group Name")
	}
	return nil
}

func (ig *ItemGroup) SaveItemGroup(db *gorm.DB) (*ItemGroup, error) {
	var err error
	err = db.Debug().Model(&ItemGroup{}).Create(&ig).Error
	if err != nil {
		return &ItemGroup{}, err
	}
	return ig, nil
}

func (ig *ItemGroup) FindAllItemGroup(db *gorm.DB) (*[]ItemGroup, error)  {
	var err error
	itemgroup := []ItemGroup{}
	err = db.Debug().Model(&ItemGroup{}).Limit(100).Find(&itemgroup).Error
	if err != nil {
		return &[]ItemGroup{}, err
	}
	return &itemgroup, err
}

func (ig *ItemGroup) FindItemGroupByID(db *gorm.DB, igid uint64) (*ItemGroup, error)  {
	var err error
	err = db.Debug().Model(ItemGroup{}).Where("item_group_id = ?", igid).Take(&ig).Error
	if err != nil {
		return &ItemGroup{}, err
	}
	if gorm.IsRecordNotFoundError(err){
		return &ItemGroup{}, errors.New("Item Group Not Found")
	}
	return ig, err
}

func (ig *ItemGroup) UpdateItemGroup(db *gorm.DB, igid uint64) (*ItemGroup, error)  {
	var err error
	db = db.Debug().Model(&ItemGroup{}).Where("item_group_id = ?", igid).Take(&ItemGroup{}).UpdateColumns(
		map[string]interface{}{
			"item_group_name" : ig.Item_Group_Name,
			"updated_at" : time.Now(),
		},
	)
	if db.Error != nil {
		return &ItemGroup{}, err
	}

	err = db.Debug().Model(&ItemGroup{}).Where("item_group_id = ?", igid).Take(&ig).Error

	if err != nil {
		return &ItemGroup{}, err
	}
	return ig, nil
}

func (ig *ItemGroup) DeleteItemGroup(db *gorm.DB, igid uint64) (int64, error) {
	db = db.Debug().Model(&ItemGroup{}).Where("item_group_id = ?", igid).Take(&ItemGroup{}).Delete(&ItemGroup{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}