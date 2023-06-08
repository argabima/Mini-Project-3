package account

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AccountControllerInterface interface {
	//EncryptPassword(password string) (string, error)
	Create(req *CreateRequest) (*AccountItemResponse, error)
	//ComparePassword(hashedPassword string, password string) error
	ReadByUsername(username string) (*ReadByUsernameResponse, error)
	Login(req *LoginRequest) (*LoginResponse, error)
	// Update(req *CreateRequest, id any) (*CreateResponse, error)
	// Delete(id any) (*CreateResponse, error)
	// FindAll() (*GetAll, error)
}
type accountController struct {
	usecase AccountUsecaseInterface
}

func NewAccountController(usecase AccountUsecaseInterface) AccountControllerInterface {
	return accountController{
		usecase: usecase,
	}
}

type AccountItemResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role_id  int    `json:"role_id"`
	Verified string
	Active   string
}

type JwtResponse struct {
	Data string `json:"token"`
}

type CreateResponse struct {
	Message string              `json:"message"`
	Data    AccountItemResponse `json:"data"`
}

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (c accountController) Create(req *CreateRequest) (*AccountItemResponse, error) {
	password, err := EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	fmt.Println(password)

	account := Actors{
		Username: req.Username,
		Password: password,
		Role_id:  req.Role_id,
		Verified: req.Verified,
		Active:   req.Active,
	}

	err = c.usecase.Create(&account)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: AccountItemResponse{
			ID:       int(account.ID),
			Username: account.Username,
			Password: account.Password,
			Role_id:  int(account.Role_id),
			Verified: account.Verified,
			Active:   account.Active,
		},
	}
	return &res.Data, nil
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

type LoginResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ReadByUsernameResponse struct {
	Message string              `json:"message"`
	Data    AccountItemResponse `json:"data"`
}

func (c accountController) ReadByUsername(username string) (*ReadByUsernameResponse, error) {
	//fmt.Println(username)
	account, err := c.usecase.getByUsername(username)
	if err != nil {
		return nil, err
	}

	res := &ReadByUsernameResponse{
		Message: "Success",
		Data: AccountItemResponse{
			ID:       int(account.ID),
			Username: account.Username,
			Password: account.Password,
			Role_id:  int(account.Role_id),
			Verified: account.Verified,
			Active:   account.Active,
		},
	}
	return res, nil
}

func (c accountController) Login(req *LoginRequest) (*LoginResponse, error) {
	account, err := c.usecase.getByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = ComparePassword(account.Password, req.Password)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"sub":  account.Role_id,
		"name": account.Username,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	}

	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return nil, err
	}

	// Gunakan signedToken seperti yang Anda butuhkan

	res := &LoginResponse{
		Message: "Success",
		Data:    signedToken,
	}
	return res, nil
}
/*
func (c accountController) Update(req *CreateRequest, id any) (*CreateResponse, error) {
	
	data, err := c.usecase.FindByID(id)

	if err != nil {
		return nil, err
	}
	
	account := Actors{
		Username: data.Username,
		Password: data.Password,
		Role_id:  data.Role_id,
		Verified: data.Verified,
		Active:   data.Active,
	}

	if req.Active != "" {
		account.Active = req.Active
	}
	if req.Verified != "" {
		account.Verified = req.Verified
	}

	createdAt := data.CreatedAt
	account.CreatedAt = createdAt

	err = c.usecase.Update(&account)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: AccountItemResponse{
			ID:       int(account.ID),
			Username: account.Username,
			Password: account.Password,
			Role_id:  int(account.Role_id),
			Verified: account.Verified,
			Active:   account.Active,
		},
	}
	return res, nil
}

func (c accountController) Delete(id any) (*CreateResponse, error) {
	data, err := c.usecase.FindByID(id)

	if err != nil {
		return nil, err
	}

	account := Actors{
		Username: data.Username,
		Password: data.Password,
		Role_id:  data.Role_id,
		Verified: data.Verified,
		Active:   data.Active,
	}

	createdAt := data.CreatedAt
	account.CreatedAt = createdAt

	err = c.usecase.Delete(&account)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: AccountItemResponse{
			ID:       int(account.ID),
			Username: account.Username,
			Password: account.Password,
			Role_id:  int(account.Role_id),
			Verified: account.Verified,
			Active:   account.Active,
		},
	}
	return res, nil
}

type GetAll struct {
	Message string              `json:"message"`
	Data    []AccountItemResponse `json:"data"`
}

func (c accountController) FindAll() (*GetAll, error) {
	data, err := c.usecase.FindAll()
	if err != nil {
		return nil, err
	}

	res := &GetAll{}
	for _, v := range data {
		res.Data = append(res.Data, AccountItemResponse{
			ID:       int(v.ID),
			Username: v.Username,
			Password: v.Password,
			Role_id:  int(v.Role_id),
			Verified: v.Verified,
			Active:   v.Active,
		})
	}

	return res, nil
}
*/