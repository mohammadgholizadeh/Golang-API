package controllers

import (
	"net/http"
	"strings"
	"time"

	"api-app/config"
	"api-app/models"
	"api-app/token"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/steambap/captcha"
)

type Register_Input struct {
	Id           int    `json:"id,string,omitempty" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        int    `json:"phone,string,omitempty" binding:"required"`
	Captcha_ID   string `json:"captchaId" binding:"required"`
	Captcha_code string `json:"captchaCode" binding:"required"`
}

type login_input struct {
	Username     string `json:"username" binding:"required"`
	Pwd          string `json:"pwd" binding:"required"`
	Captcha_ID   string `json:"captchaId" binding:"required"`
	Captcha_code string `json:"captchaCode" binding:"required"`
}

type get_captcha_info struct {
	Captcha_ID string `json:"captchaId" binding:"required"`
}

type Reset_password struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

var db_connection *gorm.DB = config.Connect()
var captchaCode string
var captchaId string
var username_tkn string
var past time.Time

// TokenValidation godoc
// @Description Respond with user information such as username, SSN, Role, Email and phone.
// @Param BearerToken header string true "Get token string"
// @Accept json
// @Produce json
// @Success 200 "Token is valid"
// @Failure 401	"Invalid token"
// @Router /api/tokenValidation [POST]
func Token_validation(c *gin.Context) {

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	_, _ = id, role
	u := models.User{}
	u.Username = username
	username_tkn = username
	dbusername, Id, email, phone, role, err := config.Get_from_db(&u, db_connection)
	if err != nil || tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Username": dbusername,
			"Id":       Id,
			"Role":     role,
			"Email":    email,
			"Phone":    phone,
		})
	}
}

// GetCaptcha godoc
// @Description Get captcha PNG from server.
// @Param refId query string true "Get random captcha id"
// @Accept json
// @Produce json
// @Success 200 "Captcha image generated"
// @Failure 400	"Empty Body"
// @Failure 500
// @Router /api/getcaptcha [POST]
func GetCaptcha(c *gin.Context) {
	var input get_captcha_info

	captchastr, err := captcha.New(180, 80, func(o *captcha.Options) { o.TextLength = 5; o.Noise = 0.3 })
	captchaCode = captchastr.Text
	input.Captcha_ID = c.Query("refId")
	captchaId = c.Query("refId")
	past = time.Now()

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(captchaCode)
	captchastr.WriteImage(c.Writer)
}

// Login godoc
// @Description Respond with Token string as Json if login was succesfull.
// @Param BearerToken header string true "Get token string"
// @Param LoginInput body controllers.login_input true "Login input body"
// @Accept json
// @Produce json
// @Success 200 {object} string "Login successfull"
// @Failure 400	"Empty Body"
// @Failure 401	"username and password or captcha is not match"
// @Failure 408	"Captcha expired"
// @Failure 500
// @Router /api/login [POST]
func Login(c *gin.Context) {
	var input login_input
	var expired bool = false
	var captcha_result bool = false

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Pwd = input.Pwd
	login_captchaCode := input.Captcha_code
	login_captchaId := input.Captcha_ID

	now := time.Now()
	CaptchaExpired := fmt.Errorf("captcha expired")
	InternalServerError := fmt.Errorf("internal server error")

	if now.Minute()-past.Minute() > 1 {
		expired = true
		c.JSON(http.StatusRequestTimeout, gin.H{"message": CaptchaExpired.Error()})
	} else if now.Minute()-past.Minute() == 1 {
		if (now.Second() + (60 - past.Second())) > 60 {
			expired = true
			c.JSON(http.StatusRequestTimeout, gin.H{"message": CaptchaExpired.Error()})
		}
	}

	if !expired && (login_captchaId == captchaId) && (login_captchaCode == captchaCode) {
		captcha_result = true
	}

	res := config.Check_from_db(&u, db_connection)

	if res && captcha_result {
		usr, id, email, phone, role, db_err := config.Get_from_db(&u, db_connection)
		_, _ = email, phone
		tkn, err := token.Jwt_maker(usr, id, role, input.Pwd)
		if err != nil && db_err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": InternalServerError.Error()})
		} else {
			id, username, role, err := token.Jwt_Decoder(tkn)
			_, _ = id, role
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{})
			} else {
				username_tkn = username
				c.JSON(http.StatusOK, gin.H{"token": tkn})
			}
		}
	} else if captcha_result {
		UserAndPassNotMatch := fmt.Errorf("login fail")
		c.JSON(http.StatusUnauthorized, gin.H{"message": UserAndPassNotMatch.Error()})
	} else if !captcha_result {
		CaptchaCodeNotMatch := fmt.Errorf("captcha code not match")
		c.JSON(http.StatusUnauthorized, gin.H{"message": CaptchaCodeNotMatch.Error()})
	}
}

func Login2(c *gin.Context) {
	var input login_input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "test"})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Pwd = input.Pwd

	InternalServerError := fmt.Errorf("internal server error")
	res := config.Check_from_db(&u, db_connection)

	if res {
		usr, id, email, phone, role, db_err := config.Get_from_db(&u, db_connection)
		_, _ = email, phone
		tkn, err := token.Jwt_maker(usr, id, role, input.Pwd)
		if err != nil && db_err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": InternalServerError.Error()})
		} else {
			id, username, role, err := token.Jwt_Decoder(tkn)
			_, _ = id, role
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{})
			} else {
				username_tkn = username
				c.JSON(http.StatusOK, gin.H{"token": tkn})
			}
		}
	} else {
		UserAndPassNotMatch := fmt.Errorf("login fail")
		c.JSON(http.StatusUnauthorized, gin.H{"message": UserAndPassNotMatch.Error()})
	}
}

// Register godoc
// @Description user registration.
// @Param RegisterInput body controllers.Register_Input true "Register input body"
// @Accept json
// @Produce json
// @Success 200 "0: This Id is already registered , 1: This Username is already registered , 2: Register was successfull"
// @Failure 400	"Empty Body"
// @Failure 408	"Captcha expired"
// @Router /api/register [POST]
func Register(c *gin.Context) {
	var input Register_Input
	var captcha_result bool = false
	var expired bool = false

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Id = input.Id
	u.Username = input.Username
	u.Pwd = input.Password
	u.Email = input.Email
	u.Phone = input.Phone
	register_captchaCode := input.Captcha_code
	register_captchaId := input.Captcha_ID

	now := time.Now()
	CaptchaExpired := fmt.Errorf("captcha expired")
	if now.Minute()-past.Minute() > 1 {
		expired = true
		c.JSON(http.StatusRequestTimeout, gin.H{"message": CaptchaExpired.Error()})
	} else if now.Minute()-past.Minute() == 1 {
		if (now.Second() + (60 - past.Second())) > 60 {
			expired = true
			c.JSON(http.StatusRequestTimeout, gin.H{"message": CaptchaExpired.Error()})
		}
	}

	if !expired && (register_captchaId == captchaId) && (register_captchaCode == captchaCode) {
		captcha_result = true
	}

	if captcha_result {
		respond := config.Save_to_db(&u, db_connection)
		c.JSON(http.StatusOK, gin.H{"message": respond})
	}
}

// ResetPassword godoc
// @Description Changes the user's password.
// @Param BearerToken header string true "Get token string"
// @Param ResetPasswordInput body controllers.Reset_password true "Reset Password Body"
// @Accept json
// @Produce json
// @Success 200 "password was successfully changed"
// @Failure 400	"Empty Body"
// @Failure 401	"Null token"
// @Failure 401	"Same pass"
// @Failure 401	"Invalid password for user"
// @Failure 500	"Internal server error"
// @Router /api/resetpassword [POST]
func ResetPassword(c *gin.Context) {
	var input Reset_password

	if username_tkn == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "null token"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	oldPwd := input.OldPassword
	newPwd := input.NewPassword

	u_new := models.User{}
	u_old := models.User{}

	u_new.Pwd = newPwd
	u_new.Username = username_tkn

	u_old.Pwd = oldPwd
	u_old.Username = username_tkn

	if oldPwd == newPwd {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "same pass"})
		return
	}

	if (oldPwd != newPwd) && (config.Check_from_db(&u_old, db_connection)) {
		respond := config.Update_db(&u_new, db_connection)

		if respond == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid"})
	}
}
