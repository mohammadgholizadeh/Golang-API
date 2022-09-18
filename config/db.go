package config

import (
	"api-app/models"
	"fmt"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"strconv"
)

//var db *gorm.DB

func Connect() *gorm.DB {
	dsn := "host=localhost user=mohammad password=321 dbname=logininfo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		return db
	}
}

func Check_from_db(U *models.User, DB *gorm.DB) bool {
	DB.AutoMigrate(&models.User{})
	var user = models.User{}
	DB.First(&user, models.User{Username: U.Username})

	if user.Pwd != U.Pwd {
		return false
	} else {
		return true
	}
}

func Save_to_db(U *models.User, DB *gorm.DB) int {
	DB.AutoMigrate(&models.User{})
	var user = models.User{}
	var user1 = models.User{}
	DB.First(&user, models.User{Id: U.Id})
	DB.First(&user1, models.User{Username: U.Username})

	// if the ID & Username exists
	if user.Id == U.Id {
		return 0
	}
	if user1.Username == U.Username {
		return 1
	}
	//U.Role = "admin"
	//DB.Select("username", "pwd", "id", "role", "email", "phone").Create(&U)
	DB.Select("username", "pwd", "id", "email", "phone").Create(&U)
	return 2
}

func Update_db(U *models.User, DB *gorm.DB) int {
	var user = models.User{}
	if err := DB.Model(&user).Where("username = ?", U.Username).Update("pwd", U.Pwd).Error; err != nil {
		return 0
	} else {
		return 2
	}

}

func Get_from_db(U *models.User, DB *gorm.DB) (string, int, string, int, string, error) {
	var user = models.User{}
	if err := DB.First(&user, models.User{Username: U.Username}).Error; err != nil {
		return "", 0, "", 0, "", err
	} else {
		return user.Username, user.Id, user.Email, user.Phone, user.Role, err
	}

}

func Save_Tickets(m *models.Tickets, DB *gorm.DB) bool {
	DB.AutoMigrate(&models.Tickets{})
	if err := DB.Select("UserId", "subject", "message", "ReadByUser", "ReadByAdmin").Create(&m).Error; err != nil {
		return false
	} else {
		return true
	}
}

func Get_all_tickets(DB *gorm.DB) ([]models.Tickets, error) {
	var res []models.Tickets
	if err := DB.Find(&res).Error; err != nil {
		return nil, err
	} else {
		return res, err
	}
}

func Get_Tickets(m *models.Tickets, DB *gorm.DB) ([]models.Tickets, error) {
	var res []models.Tickets
	if err := DB.Where("user_id = ?", strconv.Itoa(m.UserId)).Find(&res).Error; err != nil {
		return nil, err
	} else {
		return res, err
	}
}

func Get_Ticket(m *models.Tickets, DB *gorm.DB) ([]models.Tickets, error) {
	var res []models.Tickets
	fmt.Println(m.UserId)
	if err := DB.Where("id = ?", strconv.Itoa(m.Id)).Find(&res).Error; err != nil {
		return nil, err
	} else {
		return res, err
	}
}

func Update_Ticket(t *models.Tickets, DB *gorm.DB) bool {
	var ticket = models.Tickets{}
	if err := DB.Model(&ticket).Where("id = ?", t.Id).Updates(map[string]interface{}{"read_by_user": t.ReadByUser, "read_by_admin": t.ReadByAdmin}).Error; err != nil {
		return false
	} else {
		return true
	}

}

func Save_Messages(m *models.Message, DB *gorm.DB) bool {
	DB.AutoMigrate(&models.Message{})
	if err := DB.Select("AdminId", "TicketId", "Text").Create(&m).Error; err != nil {
		return false
	} else {
		return true
	}
}

func Get_messages(m *models.Message, DB *gorm.DB) ([]models.Message, error) {
	var res []models.Message
	if err := DB.Where("ticket_id = ?", strconv.Itoa(m.TicketId)).Find(&res).Error; err != nil {
		return nil, err
	} else {
		return res, err
	}
}
