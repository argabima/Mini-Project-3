package account

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateRequest struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
	Role_id  uint
	Verified string `json:"verified"`
	Active   string `json:"active"`
}

type RequestHandler struct {
	ctrl AccountControllerInterface
}

type AccountHandlerInterface interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	ReadByUsername(c *gin.Context)
	// Update(c *gin.Context)
}

func NewAccountRequestHandler(ctrl AccountControllerInterface) AccountHandlerInterface {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) AccountHandlerInterface {
	return NewAccountRequestHandler(
		NewAccountController(
			NewAccountUsecase(
				NewAccountRepository(db),
			),
		),
	)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h RequestHandler) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h RequestHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.Header("Authorization", fmt.Sprint(res.Data))
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) ReadByUsername(c *gin.Context) {

	username := c.Query("username")

	res, err := h.ctrl.ReadByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// func (h RequestHandler) Update(c *gin.Context) {
// 	var req CreateRequest
// 	id := c.Param("id")

// 	role, isErr := c.Get("role")
// 	if isErr != true {
// 		c.JSON(http.StatusBadRequest, ErrorResponse{"Role tidak ditemukan"})
// 		fmt.Println("role : ", isErr, role)
// 		return
// 	}

// 	fmt.Println("role   :   ", role)
// 	if role != "2" {
// 		c.JSON(http.StatusNonAuthoritativeInfo, ErrorResponse{"You aren't not super admin"})
// 		return
// 	}

// 	if id == "" {
// 		c.JSON(http.StatusNotFound, ErrorResponse{"Param not found"})
// 		return
// 	}
	
// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	res, err := h.ctrl.Update(&req, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// func (h RequestHandler) FindAll(c *gin.Context) {
// 	res, err := h.ctrl.FindAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

