package web

import (
	"go-react-blog/db"
	"go-react-blog/model"
	"log"
	"net/http"
	"strconv" // Add this import

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	d   db.DB
	e   *echo.Echo
}

func NewApp(d db.DB, cors bool) App {
	e := echo.New()
	app := App{
		d: d,
		e: e,
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	if cors {
		e.Use(middleware.CORS())
	}

	// Routes
	e.GET("/api/technologies", app.GetTechnologies)
	e.GET("/api/blogs", app.GetBlogs)
	e.POST("/api/blogs", app.CreateBlog)
	e.GET("/api/blogs/:id", app.GetBlog)
	e.PUT("/api/blogs/:id", app.UpdateBlog)
	e.DELETE("/api/blogs/:id", app.DeleteBlog)
	
	// Serve static files
	e.Static("/", "webapp")
	
	return app
}

func (a *App) Serve() error {
	return a.e.Start(":8080")
}

func (a *App) GetTechnologies(c echo.Context) error {
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, technologies)
}

func (a *App) GetBlogs(c echo.Context) error {
    log.Println("Fetching blogs...")
    blogs, err := a.d.GetBlogs()
    if err != nil {
        log.Printf("Error fetching blogs: %v\n", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    log.Printf("Found %d blogs\n", len(blogs))
    return c.JSON(http.StatusOK, blogs)
}

func (a *App) CreateBlog(c echo.Context) error {
	blog := new(model.Blog)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := a.d.CreateBlog(blog); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, blog)
}

func (a *App) GetBlog(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	blog, err := a.d.GetBlog(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Blog not found"})
	}
	return c.JSON(http.StatusOK, blog)
}

func (a *App) UpdateBlog(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	blog := new(model.Blog)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := a.d.UpdateBlog(id, blog); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, blog)
}

func (a *App) DeleteBlog(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	if err := a.d.DeleteBlog(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
