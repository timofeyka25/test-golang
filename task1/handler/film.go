package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"testGolang"
)

// @Summary AddFilm
// @Description  Add new film
// @Tags film
// @Accept json
// @Produce json
// @Param input body testGolang.Film true "Film"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router /film [post]
func (h *Handler) addFilm(c *gin.Context) {
	var input testGolang.Film
	if err := c.Bind(&input); err != nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"status": "bodyInvalid",
		})
		return
	}

	err := h.service.AddFilm(input)
	if err != nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"status": "bodyInvalid",
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"status": "success",
	})
}

// @Summary      Get Film
// @Description  Get Film by ID
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        ID   path      int  true  "Film ID"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router       /film/{ID} [get]
func (h *Handler) getFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"status": "error",
		})
		return
	}
	film, err := h.service.GetFilmById(id)
	if err != nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"status": "error",
		})
		return
	}
	if film == nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"status": "empty",
		})
		return
	}

	/*c.JSON(200, map[string]interface{}{
		"status": "success",
		"film":   film,
	})*/

	c.JSON(200, map[string]interface{}{
		"id":          film.Id,
		"title":       film.Title,
		"releaseDate": film.ReleaseDate,
		"status":      "success",
	})
}

// @Summary      Get Films
// @Description  Get all films
// @Tags         film
// @Produce      json
// @Param sort query string false "sort order" Enums(releaseDate, -releaseDate)
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router       /film [get]
func (h *Handler) getFilms(c *gin.Context) {
	sort := c.Query("sort")
	films, err := h.service.GetFilms(sort)
	if err != nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"status": "error",
		})
		return
	}
	if films == nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"status": "empty",
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   films,
	})
}
