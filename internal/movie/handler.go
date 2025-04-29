package movie

import (
	"itv-go/internal/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

type Response struct {
	Message string `json:"Message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// FakeLoginHandler godoc
// @Summary Generate a fake JWT token
// @Description Generates a fake JWT token for a dummy user for testing purposes
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func FakeLoginHandler(c *gin.Context) {
	fakeUsername := "username"

	token, err := auth.GenerateJWT(fakeUsername)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, Response{
		Message: token,
	})
}

// CreateMovie godoc
// @Router /movies [POST]
// @Summary Create a new movie
// @Description Add a new movie
// @Tags movies
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param movie body Movie true "Movie to create"
// @Success 201 {object} Movie
// @Failure 400 {object} ErrorResponse
// @Router /movies [post]
func (h *Handler) CreateMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.service.Create(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// GetAllMovies godoc
// @Summary Get all movies
// @Description Retrieve all movies
// @Tags movies
// @Security BearerAuth
// @Produce json
// @Success 200 {array} Movie
// @Failure 500 {object} ErrorResponse
// @Router /movies [get]
func (h *Handler) GetAllMovies(c *gin.Context) {
	movies, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// GetMovieByID godoc
// @Summary Get a movie by ID
// @Description Retrieve a movie by ID
// @Tags movies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} Movie
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /movies/{id} [get]
func (h *Handler) GetMovieByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid movie ID"})
		return
	}

	movie, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// UpdateMovie godoc
// @Summary Update a movie
// @Description Update an existing movie by ID
// @Tags movies
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param movie body Movie true "Movie data to update"
// @Success 200 {object} Movie
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /movies/{id} [put]
func (h *Handler) UpdateMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid movie ID"})
		return
	}

	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if err := h.service.Update(uint(id), &movie); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "movie not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "movie updated successfully"})
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie by ID
// @Tags movies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /movies/{id} [delete]
func (h *Handler) DeleteMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid movie ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "movie not found"})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "movie deleted successfully"})
}
