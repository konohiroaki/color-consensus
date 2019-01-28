package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/dto"
	"github.com/twinj/uuid"
	"net/http"
	"time"
)

func Route() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", false)))
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	apiRouterGroup := router.Group("/api")
	{
		colorAPI(apiRouterGroup)
		userAPI(apiRouterGroup)
	}
	return router
}

func colorAPI(router *gin.RouterGroup) {
	router.POST("/colors/:lang/:color/:user", func(c *gin.Context) {
		// TODO: move session related logic to non-api endpoint.
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			c.AbortWithStatus(http.StatusForbidden)
		}

		var colors []string
		c.BindJSON(&colors) // TODO: error handling // if err := bind(); err != nil { handing... )
		lang := c.Param("lang")
		color := c.Param("color")
		vote := dto.ColorVote{Language: lang, Color: color, User: userID.(string), Date: time.Now(), Colors: colors}

		dto.Raw = append(dto.Raw, &vote)
		sum, found := findSum(lang, color)
		if found {
			sum.Vote += 1
			// TODO: distinct color list
			for _, votedColor := range vote.Colors {
				if _, found := sum.Colors[votedColor]; found {
					sum.Colors[votedColor] += 1
				} else {
					sum.Colors[votedColor] = 1
				}
			}
			c.JSON(200, sum)
		} else {
			c.AbortWithStatus(404)
		}
	})
	router.GET("/colors", func(c *gin.Context) {
		//TODO: provide a way to get only language list
		c.JSON(200, dto.Sum)
	})
	router.GET("/colors/:lang", func(c *gin.Context) {
		list := findSumListOfLang(c.Param("lang"))
		c.JSON(200, list)
	})
	router.GET("/colors/:lang/:color", func(c *gin.Context) {
		lang := c.Param("lang")
		color := c.Param("color")
		sum, found := findSum(lang, color)
		if found {
			c.JSON(200, sum)
		} else {
			c.AbortWithStatus(404)
		}
	})
	router.GET("/colors/:lang/:color/raw", func(c *gin.Context) {
		//TODO: do pagination
		lang := c.Param("lang")
		color := c.Param("color")
		list := findVoteList(lang, color)
		c.JSON(200, list)
	})
}

func userAPI(router *gin.RouterGroup) {
	router.POST("/users", func(c *gin.Context) {
		var user dto.User
		c.BindJSON(&user) // TODO: error handling // if err := bind(); err != nil { handing... )
		user.ID = uuid.NewV4().String()
		user.Date = time.Now()
		dto.Users = append(dto.Users, &user)
		// TODO: move session related logic to non-api endpoint.
		session := sessions.Default(c)
		session.Set("userID", user.ID)
		session.Save()
		c.JSON(200, user);
	})
	router.GET("/admin/users", func(c *gin.Context) {
		c.JSON(200, dto.Users)
	})
}

func findSumListOfLang(lang string) []dto.ColorConsensus {
	list := []dto.ColorConsensus{}
	for _, ele := range dto.Sum {
		if ele.Language == lang {
			list = append(list, *ele)
		}
	}
	return list
}
func findSum(lang, color string) (*dto.ColorConsensus, bool) {
	for _, ele := range dto.Sum {
		if ele.Language == lang && ele.Color == color {
			return ele, true
		}
	}
	return nil, false
}
func findVoteList(lang, color string) []dto.ColorVote {
	list := []dto.ColorVote{};
	for _, ele := range dto.Raw {
		if ele.Language == lang && ele.Color == color {
			list = append(list, *ele)
		}
	}
	return list
}
