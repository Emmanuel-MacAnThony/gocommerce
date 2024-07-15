package handlers

import (
	"net/http"

	"github.com/Emmanuel-MacAnThony/gocommerce/internal/api/rest"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/dto"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/repository"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	svc := service.UserService{Repo: repository.NewUserRepository(rh.DB)}
	handler := UserHandler{svc: svc}
	app := rh.App

	//public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	//private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)
	app.Get("/cart", handler.AddToCart)
	app.Post("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrder)
	app.Get("/order/:id", handler.GetOrders)
	app.Post("/become-seller", handler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Please provide valid inputs"})
	}
	token, err := h.svc.SignUp(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "error on signup"})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": token})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Please provide valid inputs"})
	}
	token, err := h.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "error on login, Invalid email or password"})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Login Successful", "token": token})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "verify"})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "verification-code"})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "profile"})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "profile"})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "cart"})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "cart"})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "order"})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "order"})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "seller"})
}
