package controllers

import (
	"net/http"
	"strings"
	"time"

	"api-app/config"
	"api-app/models"
	"api-app/token"

	"github.com/gin-gonic/gin"
)

type Ticketing_input struct {
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type MessageForTicketing_input struct {
	TicketId int    `json:"ticketid,string,omitempty" binding:"required"`
	Text     string `json:"text" binding:"required"`
}

type GetTicketInfo_input struct {
	TicketId int `json:"ticketid,string,omitempty" binding:"required"`
}

// struct for Documentation
type GetTicketsList_U struct {
	HasbeenRead   int
	TicketSubject string
	TicketMessage string
	TicketID      int
}

// struct for Documentation
type GetTicketInfo struct {
	TicketSubject  string
	TicketMessage  string
	TextsForTicket string
	TextTime       time.Time
}

// struct for Documentation
type GetTicketsList_A struct {
	HasbeenRead    int
	TicketsSubject string
	TicketsMessage string
	TicketId       int
}

// Ticketing godoc
// @Description Get user's Message.
// @Accept json
// @Produce json
// @Success 200 "Message was successfully received"
// @Failure 400	"Empty Body"
// @Failure 401	"Null token"
// @Failure 500	"Problem in saving message"
// @Router /api/ticketing [POST]
func Ticketing(c *gin.Context) {
	var input Ticketing_input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	_, _ = username, role
	if tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}

	m := models.Tickets{}
	m.UserId = id
	m.Subject = input.Subject
	m.Message = input.Message
	m.ReadByAdmin = 0
	m.ReadByUser = 1

	res := config.Save_Tickets(&m, db_connection)
	if res {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}

// Get users tickets list godoc
// @Description Respond with users tickets list.
// @Param BearerToken header string true "Get token string"
// @Accept json
// @Produce json
// @Success 200 {object} controllers.GetTicketsList_U
// @Failure 401	"Token is not valid"
// @Failure 500
// @Router /api/getTicketsListU [GET]
func GetTicketsList_User(c *gin.Context) {
	var hasbeenRead = []int{}
	var ticketSubject = []string{}
	var ticketMessage = []string{}
	var ticketId = []int{}

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	_, _ = username, role
	if tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	m := models.Tickets{}
	m.UserId = id

	res, err := config.Get_Tickets(&m, db_connection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	} else {
		for _, res := range res {
			hasbeenRead = append(hasbeenRead, res.ReadByAdmin)
			ticketSubject = append(ticketSubject, res.Subject)
			ticketMessage = append(ticketMessage, res.Message)
			ticketId = append(ticketId, res.Id)
		}
		c.JSON(http.StatusOK, gin.H{
			"has been Read":   hasbeenRead,
			"tickets Subject": ticketSubject,
			"tickets Message": ticketMessage,
			"ticketId":        ticketId,
		})

	}
}

// Get Ticket information godoc
// @Description Respond with User ticket info.
// @Param BearerToken header string true "Get token string"
// @Param GetTicketInfo body controllers.GetTicketInfo_input true "GetTicketInfo input body"
// @Accept json
// @Produce json
// @Success 200 {object} controllers.GetTicketInfo
// @Failure 400	"Empty Body"
// @Failure 401	"Token is not valid"
// @Failure 500
// @Router /api/GetTicketInfoU [GET]
func GetTicketInfo_User(c *gin.Context) {
	var ticketSubject = []string{}
	var ticketMessage = []string{}
	var texts_for_ticket = []string{}
	var time = []time.Time{}
	var input GetTicketInfo_input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	_, _, _ = username, id, role
	if tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	m := models.Message{}
	m.TicketId = input.TicketId

	t := models.Tickets{}
	t.Id = input.TicketId
	t.ReadByUser = 1

	messages, getMessage_err := config.Get_messages(&m, db_connection)
	update_ticket := config.Update_Ticket(&t, db_connection)
	getTicket, getTicket_err := config.Get_Ticket(&t, db_connection)

	if (getMessage_err != nil) || (getTicket_err != nil) || !update_ticket {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		for _, message := range messages {
			texts_for_ticket = append(texts_for_ticket, message.Text)
			time = append(time, message.CreatedAt)
		}
		for _, getTicket := range getTicket {
			ticketSubject = append(ticketSubject, getTicket.Subject)
			ticketMessage = append(ticketMessage, getTicket.Message)

		}
		c.JSON(http.StatusOK, gin.H{
			"ticket subject":   ticketSubject,
			"ticket message":   ticketMessage,
			"texts for ticket": texts_for_ticket,
			"text time":        time,
		})

	}
}

// Message for ticket godoc
// @Description Respond with Json body.
// @Param BearerToken header string true "Get token string"
// @Param MessageForTicketing body controllers.MessageForTicketing_input true "MessageForTicketing input body"
// @Accept json
// @Produce json
// @Success 200 "Succesfully message saved"
// @Failure 400	"Empty Body"
// @Failure 401	"Token is not valid"
// @Failure 500
// @Router /api/messageForTicketU [POST]
func MessageForTicket_User(c *gin.Context) {
	var input MessageForTicketing_input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	_, _, _ = username, id, role
	if tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	m := models.Message{}
	m.TicketId = input.TicketId
	m.Text = input.Text

	res := config.Save_Messages(&m, db_connection)
	if res {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}

// Get admins Tickets list godoc
// @Description Respond with admins tickets list.
// @Param BearerToken header string true "Get token string"
// @Accept json
// @Produce json
// @Success 200 {object} controllers.GetTicketsList_A
// @Failure 401	"token in not valid or role is not admin"
// @Failure 500
// @Router /api/getTicketsListA [GET]
func GetTicketsList_Admin(c *gin.Context) {
	var hasbeenRead = []int{}
	var ticketSubject = []string{}
	var ticketMessage = []string{}
	var ticketId = []int{}

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	_, _ = username, id
	if tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	res, err := config.Get_all_tickets(db_connection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	} else {
		for _, res := range res {
			hasbeenRead = append(hasbeenRead, res.ReadByUser)
			ticketSubject = append(ticketSubject, res.Subject)
			ticketMessage = append(ticketMessage, res.Message)
			ticketId = append(ticketId, res.Id)
		}
		c.JSON(http.StatusOK, gin.H{
			"has been Read":   hasbeenRead,
			"tickets Subject": ticketSubject,
			"tickets Message": ticketMessage,
			"ticketId":        ticketId,
		})

	}
}

// Get ticket info for admin godoc
// @Description Respond with Ticket info.
// @Param BearerToken header string true "Get token string"
// @Param GetTicketInfo body controllers.GetTicketInfo_input true "GetTicketInfo input body"
// @Accept json
// @Produce json
// @Success 200 {object} controllers.GetTicketInfo
// @Failure 400	"Empty Body"
// @Failure 401	"token in not valid or role is not admin"
// @Failure 500
// @Router /api/GetTicketInfoA [POST]
func GetTicketInfo_Admin(c *gin.Context) {
	var ticketSubject = []string{}
	var ticketMessage = []string{}
	var texts_for_ticket = []string{}
	var time = []time.Time{}
	var input GetTicketInfo_input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	_, _, _ = username, id, role
	if tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	m := models.Message{}
	m.TicketId = input.TicketId

	t := models.Tickets{}
	t.Id = input.TicketId
	t.ReadByAdmin = 1
	t.ReadByUser = 0

	messages, getMessage_err := config.Get_messages(&m, db_connection)
	update_ticket := config.Update_Ticket(&t, db_connection)
	getTicket, getTicket_err := config.Get_Ticket(&t, db_connection)

	if (getMessage_err != nil) || (getTicket_err != nil) || !update_ticket {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		for _, message := range messages {
			texts_for_ticket = append(texts_for_ticket, message.Text)
			time = append(time, message.CreatedAt)
		}
		for _, getTicket := range getTicket {
			ticketSubject = append(ticketSubject, getTicket.Subject)
			ticketMessage = append(ticketMessage, getTicket.Message)

		}
		c.JSON(http.StatusOK, gin.H{
			"ticket subject":   ticketSubject,
			"ticket message":   ticketMessage,
			"texts for ticket": texts_for_ticket,
			"text time":        time,
		})

	}
}

// Message for ticket from admin godoc
// @Description Respond with Json body.
// @Param BearerToken header string true "Get token string"
// @Param MessageForTicketing body controllers.MessageForTicketing_input true "MessageForTicketing input body"
// @Accept json
// @Produce json
// @Success 200 "Succesfully message saved"
// @Failure 400	"Empty Body"
// @Failure 401	"token in not valid or role is not admin"
// @Failure 500
// @Router /api/messageForTicketU [POST]
func MessageForTicket_Admin(c *gin.Context) {
	var input MessageForTicketing_input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tknStr := c.Request.Header.Get("Authorization")
	tknStr = strings.Split(tknStr, " ")[1]
	id, username, role, tknerr := token.Jwt_Decoder(tknStr)
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	_, _ = username, role
	if tknerr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	m := models.Message{}
	m.TicketId = input.TicketId
	m.Text = input.Text
	m.AdminId = id

	res := config.Save_Messages(&m, db_connection)
	if res {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}
