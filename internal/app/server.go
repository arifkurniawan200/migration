package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"migration/internal/repository"
	"net/http"
	"time"
)

type Handler struct {
	Destination repository.DestinationProduct
	Source      repository.SourceProduct
}

func Run(d repository.DestinationProduct, s repository.SourceProduct) {
	e := echo.New()

	handler := Handler{
		Destination: d,
		Source:      s,
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Rate Limiter Configuration
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	e.Use(middleware.RateLimiterWithConfig(config))

	src := e.Group("/source")
	{
		src.GET("/", handler.GetAllSourceProduct)
	}

	dst := e.Group("/destination")
	{
		dst.GET("/", handler.GetAllDestinationProduct)
		dst.PUT("/", handler.UpdateDestinationProduct)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
