package routing

import (
	"github.com/gin-gonic/gin"
)

func startupWebServer() {
	router := gin.Default()
	registerAll(router)
	router.Run()
}

func registerAll(r *gin.Engine) {
	r.PUT("/:id", CreateServer)
	r.DELETE("/:id", RemoveServer)
	r.GET("/:id", GetServerInfo)
	r.POST("/:id", EditServer)

	r.POST("/:id/start", StartServer)
	r.POST("/:id/stop", StopServer)
	r.POST("/:id/restart", RestartServer)
	r.POST("/:id/kill", KillServer)
	r.POST("/:id/reinstall", ReinstallServer)

	r.GET("/:id/file/*filename", GetFile)
	r.PUT("/:id/file*filename", PutFile)
	r.DELETE("/:id/file/*filename", DeleteFile)

	r.POST("/:id/console", SendCommand)
	r.GET("/:id/console", GetLog)
}

//Creates a server.
//Requires an image id, a custom name, and all required options/environment variables for the selected image
func CreateServer(con *gin.Context) {

}

//Removes the server that matches with the id passed with the URL
//Server data is archived and kept in storage for 60 days
func RemoveServer(con *gin.Context) {

}

//Returns all information that belongs to a server, including status, ontime and custominfo
func GetServerInfo(con *gin.Context) {

}

//Updates certain options/environment variables for a server (i.e. change RAM)
func EditServer(con *gin.Context) {

}

func StartServer(con *gin.Context) {

}

func StopServer(con *gin.Context) {

}

func RestartServer(con *gin.Context) {

}

func KillServer(con *gin.Context) {

}

func ReinstallServer(con *gin.Context) {

}

func GetFile(con *gin.Context) {

}

func PutFile(con *gin.Context) {

}

func DeleteFile(con *gin.Context) {

}

func SendCommand(con *gin.Context) {

}

func GetLog(con *gin.Context) {

}