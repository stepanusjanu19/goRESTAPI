package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/stepanusjanu19/goRESTAPI/api/models"
)

var users = []models.User{
	models.User{
		Username: "Stepanus Janu",
		Email:    "kitoyu22@gmail.com",
		Password: "password",
	},
	models.User{
		Username: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "GO BASIC WEB",
		Content: "Welcome REST CLIENT Go",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

var itemgroup = []models.ItemGroup{
	models.ItemGroup{
		Item_Group_Name: "Baju & Sweater",
	},
	models.ItemGroup{
		Item_Group_Name: "Makanan & Minuman",
	},
}
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}, &models.ItemGroup{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.ItemGroup{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(user_id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].USER_ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

		err = db.Debug().Model(&models.ItemGroup{}).Create(&itemgroup[i]).Error
		if err != nil {
			log.Fatalf("cannot seed item group table: %v", err)
		}
	}
}