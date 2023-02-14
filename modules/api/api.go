package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/Mrpye/cisco-webex-bot/docs"
	"github.com/Mrpye/cisco-webex-bot/modules/body_types"
	"github.com/Mrpye/cisco-webex-bot/modules/webex"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var web_port string
var web_ip string

// @Summary Send  a message to Webex text format
// @ID post-send-message-text
// @Produce json
// @Param request body body_types.SendMessage.request true true "query params"
// @Success 200 {string}  string "message sent"
// @Failure 404 {string}  string "error"
// @Router /send_message [post]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func postSendMessage(c *gin.Context) {
	var importRequest body_types.SendMessage

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}

	if importRequest.Message == "" {
		c.IndentedJSON(http.StatusBadRequest, "No Message")
		return
	}

	//*******************
	//Perform the install
	//*******************
	web_webex_token := ""
	if val, ok := c.Request.Header["Authorization"]; ok {
		web_webex_token = strings.ReplaceAll(val[0], "Bearer ", "")
	}
	if web_webex_token == "" {
		c.IndentedJSON(http.StatusForbidden, "No Message")
		return
	}

	web_ex := webex.CreateK8(web_webex_token)
	err := web_ex.SendMessage(importRequest.Message, false)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	} else {
		c.IndentedJSON(http.StatusCreated, "Message Sent")
	}

}

// @Summary Send  a message to Webex markdown format
// @ID post-send-message-markdown
// @Produce json
// @Param request body body_types.SendMessage.request true true "query params"
// @Success 200 {string}  string "message sent"
// @Failure 404 {string}  string "error"
// @Router /send_message_md [post]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func postSendMessageMD(c *gin.Context) {
	var importRequest body_types.SendMessage

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}

	if importRequest.Message == "" {
		c.IndentedJSON(http.StatusBadRequest, "No Message")
		return
	}

	//*******************
	//Perform the install
	//*******************
	web_webex_token := ""
	if val, ok := c.Request.Header["Authorization"]; ok {
		web_webex_token = strings.ReplaceAll(val[0], "Bearer ", "")
	}
	if web_webex_token == "" {
		c.IndentedJSON(http.StatusForbidden, "No Message")
		return
	}

	web_ex := webex.CreateK8(web_webex_token)
	err := web_ex.SendMessage(importRequest.Message, true)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	} else {
		c.IndentedJSON(http.StatusCreated, "Message Sent")
	}

}

// @Summary Check API Endpoint
// @ID check-api-endpoint
// @Produce json
// @Success 200 {string}  string "ok"
// @Router / [get]
func getOK(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "OK")
}

// Function to start web server
// ip - IP address to listen on
// port - Port to listen on
func StartWebServer(ip string, port string) {
	//****************
	//Set the variable
	//****************
	web_ip = ip
	web_port = port

	//*****************
	//Set up the server
	//*****************
	fmt.Println("Starting Web-Server")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//*****************
	//Set up the routes
	//*****************
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/send_message", postSendMessage)
	router.POST("/send_message_md", postSendMessageMD)
	router.GET("/", getOK)

	//**********************************
	//Set up the environmental variables
	//**********************************
	if os.Getenv("WEB_IP") != "" {
		web_ip = os.Getenv("WEB_IP")
	}
	if os.Getenv("WEB_PORT") != "" {
		web_port = os.Getenv("WEB_PORT")
	}

	//****************
	//Start the server
	//****************
	fmt.Printf("Started Web-Server %s:%s", web_ip, web_port)
	router.Run(fmt.Sprintf("%s:%s", web_ip, web_port))

}
