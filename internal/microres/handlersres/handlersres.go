package handlersres

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
)

func CrearAsignatura(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		err := repo.CrearAsignatura(asig)

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return
		} else {
			c.JSON(http.StatusCreated, gin.H{"Location": "resenias/asignatura/"+asig})
			return
		} 
	}
}

func GetResenias(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		resenias, err := repo.GetResenias(asig)


		if err != nil{
			if err.Error() == "Algo salió mal con la reseña:  la asignatura no está registrada"{
				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		} else {
			c.JSON(http.StatusOK, gin.H{"resenias": resenias})
			return
		} 
	}
}

func GetResenia(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		id, err1 := strconv.Atoi( c.Param("id") )

		if err1 != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Identificador no es un entero"})
			return
		}
		resenia, err := repo.GetResenia(asig, id)


		if err != nil{
			if msg := err.Error(); (msg == "Algo salió mal con la reseña:  la asignatura no está registrada" ||
				msg == "Algo salió mal con la reseña:  la reseña no contiene un identificador válido"){

				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			} 
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		} else {

			c.JSON(http.StatusOK, gin.H{"resenia": resenia})
			return
		} 
	}
}

type CreateReseniaInput struct {
	Opinion string `json:"opinion" binding:"required"`
}

func Opinar(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		var input CreateReseniaInput
		
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se han enviado los campos requeridos"})
			return
		}

		err := repo.Opinar(asig, input.Opinion)

		if err != nil{
			if err.Error() == "Algo salió mal con la reseña:  la asignatura no está registrada"{
				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		} else {
			res, _ := repo.GetResenias(asig)
			dir := "resenias/asignatura/" + asig + "/" + strconv.Itoa(len(res)-1)
			c.JSON(http.StatusCreated, gin.H{"Location": dir})
			return
		} 
	}
}

func GustaResenia(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {

		asig := c.Param("asig")
		id,err := strconv.Atoi( c.Param("id") )

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Identificador no es un entero"})
			return
		}


		err = repo.GustaResenia(asig, id)
		if err != nil{
			if msg := err.Error(); (msg == "Algo salió mal con la reseña:  la asignatura no está registrada" ||
				msg == "Algo salió mal con la reseña:  la reseña no contiene in identificador válido"){

				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			} 
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		}
		c.JSON(http.StatusCreated, gin.H{"Mensaje": "creada correctamente"}) 
		
	}
}

func NoGustaResenia(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {

		asig := c.Param("asig")
		id,err := strconv.Atoi( c.Param("id") )

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Identificador no es un entero"})
			return
		}


		err = repo.NoGustaResenia(asig, id)
		if err != nil{
			if msg := err.Error(); (msg == "Algo salió mal con la reseña:  la asignatura no está registrada" ||
				msg == "Algo salió mal con la reseña:  la reseña no contiene in identificador válido"){

				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			} 
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		}
		c.JSON(http.StatusCreated, gin.H{"Mensaje": "creada correctamente"}) 
		
	}
}