package app

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func (u Handler) GetAllSourceProduct(c echo.Context) error {
	data, err := u.Source.GetProductSource()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed process request",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success fetch data",
		"data":    data,
	})
}

func (u Handler) UpdateDestinationProduct(c echo.Context) error {
	data, err := u.Source.GetProductSource()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed process request",
			"data":    err.Error(),
		})
	}
	go func() {
		// delete time.sleep if you don't want to pause the execution
		time.Sleep(1 * time.Second)
		log.Printf("start background job in %v\n", time.Now())
		err := u.Destination.UpdateProductDestinationTx(data)
		if err != nil {
			log.Printf("error where update data %v\n", time.Now())
			return
		}
		log.Printf("finish background job in %v\n", time.Now())
	}()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data is in progress",
	})
}

func (u Handler) GetAllDestinationProduct(c echo.Context) error {
	data, err := u.Destination.GetProductDestination()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed process request",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success fetch data",
		"data":    data,
	})
}
