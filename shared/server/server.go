package server
import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port string 			//`json:"Port"`
	Mode string            	//`json:"Mode"`
	TemplateFolder string 	//`json:"TemplateFolder"`
}

func Run(r *gin.Engine, server Server) {
	switch server.Mode {
		case "Development":
			gin.SetMode(gin.DebugMode)
		case "Productioni":
			gin.SetMode(gin.ReleaseMode)
	}

	r.LoadHTMLGlob(server.TemplateFolder)
	r.Run(server.Port)
}