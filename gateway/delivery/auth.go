package delivery

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tgkzz/ecom/gateway/service"
	innerpb "github.com/tgkzz/ecom/innerpb/auth"
)

type authRouter struct {
	engine  *echo.Echo
	logger  *slog.Logger
	service *service.Service
}

func NewAuthRouter(e *echo.Echo, l *slog.Logger, service *service.Service) Handler {
	return &authRouter{engine: e, logger: l, service: service}
}

func (a *authRouter) RegisterRoutes() error {
	auth := a.engine.Group("/auth")

	a.registerAuthRoutes(auth)

	return nil
}

func (a *authRouter) registerAuthRoutes(auth *echo.Group) {
	user := auth.Group("/user")
	user.POST("/register", a.register)
	user.POST("/login", a.login)
}

func (a *authRouter) login(c echo.Context) error {
	const op = "authRouter.login"

	log := a.logger.WithGroup(op)

}

func (a *authRouter) register(c echo.Context) error {
	const op = "authRouter.register"

	log := a.logger.WithGroup(op)

	var req registerRequest
	if err := c.Bind(&req); err != nil {
		log.Error("error while binding request", "err", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := a.service.Auth.Register(c.Request().Context(), &innerpb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Error("error while register", "err", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, registerResponse{
		OTP: resp.Otp,
	})
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	OTP string `json:"otp"`
}
