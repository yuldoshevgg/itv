package router

import (
	"itv-go/internal/db"
	"itv-go/internal/middleware"
	"itv-go/internal/movie"

	swaggerFiles "github.com/swaggo/files"

	_ "itv-go/api/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func NewRouter(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()

	dbConn, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	if err := dbConn.AutoMigrate(&movie.Movie{}); err != nil {
		panic(err)
	}

	movieRepo := movie.NewRepository(dbConn)
	movieService := movie.NewService(movieRepo)
	movieHandler := movie.NewHandler(movieService)

	api := r.Group("/api")
	{
		api.POST("/login", fakeLoginHandler)
		api.Use(middleware.AuthMiddleware())
		api.POST("/movies", movieHandler.CreateMovie)
		api.GET("/movies", movieHandler.GetAllMovies)
		api.GET("/movies/:id", movieHandler.GetMovieByID)
		api.PUT("/movies/:id", movieHandler.UpdateMovie)
		api.DELETE("/movies/:id", movieHandler.DeleteMovie)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DocExpansion("none"),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	return r
}
