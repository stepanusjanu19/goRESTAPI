package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ITEM_ID uint64 `gorm:"primary_key;auto_increment" json:"item_id"`
	Item_Name string `gorm:"size:255;not null;unique" json:"item_name"`
	Brand_Item_Name string `gorm:"size:255;not null;unique" json:"brand_item_name"`
	Stok_Item uint64 `gorm:"not null" json:"stok_item"`
	ItemGroup ItemGroup `json:"itemgroup"`
	ItemGroupID uint64 `gorm:"not null" json:"item_group_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (i *Item) Prepare()  {
	i.ITEM_ID = 0
	i.Item_Name = html.EscapeString(strings.TrimSpace(i.Item_Name))
	i.Brand_Item_Name = html.EscapeString(strings.TrimSpace(i.Brand_Item_Name))
	i.Stok_Item = 0
	i.ItemGroup = ItemGroup{}
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
}

func (i *Item) Validate() error  {
	if i.Item_Name == "" {
		return errors.New("Required Item Name")
	}
	if i.Brand_Item_Name == "" {
		return errors.New("Required Brand Item Name")
	}
	if i.ItemGroupID < 1 {
		return errors.New("Required Item Group ID")
	}
	return nil
}

func (i *Item) SaveItem(db *gorm.DB) (*Item, error)  {
	var err error
	err = db.Debug().Model(&Item{}).Create(&i).Error
	if err != nil {
		return &Item{}, err
	}
	if i.ITEM_ID != 0 {
		err = db.Debug().Model(&ItemGroup).Where("item_id = ?", i.ItemGroupID).Take(&i.ItemGroup).Error
		if err != nil {
			return &Item{}, err
		}
	}
	return i, nil
}

func (i *Item) FindAllItems(db *gorm.DB) (*[]Item, error)  {
	var err error
	items := []Item{}
	err = db.Debug().Model(&Item{}).Limit(100).Find(&items).Error
	if err != nil {
		return &[]Item{}, err
	}
	if len(items) > 0 {
		for i, _ := range items {
			err := db.Debug().Model(&Item{}).Where("item_id = ?", items[i].ItemGroupID).Take(&items[i].ItemGroup).Error
			if err != nil {
				return &[]Item{}, err
			}
		}
	}
	return &items, nil
}

func (i *Item) FindItemByID(db *gorm.DB, iid uint64) (*Item, error)  {
	var err error
	err = db.Debug().Model(&Item{}).Where("item_id = ?", iid).Take(&i).Error
	if err != nil {
		return &Item{}, err
	}
	if i.ITEM_ID != 0 {
		err = db.Debug().Model(&ItemGroup{}).Where("item_id = ?", i.ItemGroupID).Take(&i.ItemGroup).Error
		if err != nil {
			return &Item{}, err
		}
	}
	return i, nil
}

func (i *Item) UpdateItem(db *gorm.DB) (*Item, error)  {
	var err error
	err = db.Debug().Model(&Item{}).Where("item_id = ?", i.ITEM_ID).Updates(Item{Item_Name: i.Item_Name, Brand_Item_Name: i.Brand_Item_Name, Stok_Item: i.Stok_Item, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Item{}, err
	}
	if i.ITEM_ID != 0 {
		err = db.Debug().Model(&ItemGroup{}).Where("item_id = ?", i.ItemGroupID).Take(&i.ItemGroup).Error
		if err != nil {
			return &Item{}, err
		}
	}
	return i, nil
}

func (i *Item) DeleteItem(db *gorm.DB, iid uint64, igid uint64) (int64, error)   {
	db = db.Debug().Model(&Item{}).Where("item_id = ? and item_group_id = ?", iid, igid).Take(&Item{}).Delete(&Item{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Item Not Found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}






