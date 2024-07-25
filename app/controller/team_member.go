package controller

import (
	"net/http"
	"strconv"
	"strings"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-skeleton-fiber/app/configs"
	"github.com/adamnasrudin03/go-skeleton-fiber/app/dto"
	"github.com/adamnasrudin03/go-skeleton-fiber/app/middlewares"
	"github.com/adamnasrudin03/go-skeleton-fiber/app/service"
	"github.com/adamnasrudin03/go-skeleton-fiber/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TeamMemberController interface {
	Mount(group fiber.Router)
	Create(c *fiber.Ctx) error
	GetDetail(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	GetList(c *fiber.Ctx) error
}

type TeamMemberHandler struct {
	Service  service.TeamMemberService
	Cfg      *configs.Configs
	Logger   *logrus.Logger
	Validate *validator.Validate
}

func NewTeamMemberDelivery(
	srv service.TeamMemberService,
	cfg *configs.Configs,
	logger *logrus.Logger,
	validator *validator.Validate,
) TeamMemberController {
	return &TeamMemberHandler{
		Service:  srv,
		Cfg:      cfg,
		Logger:   logger,
		Validate: validator,
	}
}
func (h *TeamMemberHandler) Mount(group fiber.Router) {
	group.Post("", middlewares.BasicAuth(h.Cfg.App.BasicUsername, h.Cfg.App.BasicPassword), h.Create)
	group.Get("", h.GetList)
	group.Get("/:id", h.GetDetail)
	group.Delete("/:id", middlewares.BasicAuth(h.Cfg.App.BasicUsername, h.Cfg.App.BasicPassword), h.Delete)
	group.Put("/:id", middlewares.BasicAuth(h.Cfg.App.BasicUsername, h.Cfg.App.BasicPassword), h.Update)
}

func (h *TeamMemberHandler) getParamID(c *fiber.Ctx) (uint64, error) {
	idParam := strings.TrimSpace(c.Params("id"))
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		h.Logger.Errorf("TeamMemberController-getParamID error parse param: %v ", err)
		return 0, response_mapper.ErrInvalid("ID Anggota team", "Team Member ID")
	}
	return id, nil
}

func (h *TeamMemberHandler) Create(c *fiber.Ctx) error {
	var (
		opName = "TeamMemberController-Create"
		ctx    = c.Context()
		input  dto.TeamMemberCreateReq
		err    error
	)

	err = c.BodyParser(&input)
	if err != nil {
		h.Logger.Errorf("%v error bind json: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrGetRequest())
	}

	// validation input user
	err = h.Validate.Struct(input)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, response_mapper.FormatValidationError(err))
	}

	res, err := h.Service.Create(ctx, input)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(response_mapper.RenderStruct(http.StatusCreated, res))
}

func (h *TeamMemberHandler) GetDetail(c *fiber.Ctx) error {
	var (
		opName = "TeamMemberController-GetDetail"
		ctx    = c.Context()
		err    error
	)

	id, err := h.getParamID(c)
	if err != nil {
		return utils.HttpError(c, err)
	}

	res, err := h.Service.GetByID(ctx, id)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, err)
	}

	return c.Status(http.StatusOK).JSON(response_mapper.RenderStruct(http.StatusOK, res))
}

func (h *TeamMemberHandler) Delete(c *fiber.Ctx) error {
	var (
		opName = "TeamMemberController-Delete"
		ctx    = c.Context()
		err    error
	)

	id, err := h.getParamID(c)
	if err != nil {
		return utils.HttpError(c, err)
	}

	err = h.Service.DeleteByID(ctx, id)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, err)
	}

	return c.Status(http.StatusOK).JSON(response_mapper.RenderStruct(http.StatusOK, response_mapper.MultiLanguages{
		ID: "Anggota Tim Berhasil Dihapus",
		EN: "Team Member Deleted Successfully",
	}))
}

func (h *TeamMemberHandler) Update(c *fiber.Ctx) error {
	var (
		opName = "TeamMemberController-Update"
		ctx    = c.Context()
		input  dto.TeamMemberUpdateReq
		err    error
	)

	id, err := h.getParamID(c)
	if err != nil {
		return utils.HttpError(c, err)
	}

	err = c.BodyParser(&input)
	if err != nil {
		h.Logger.Errorf("%v error bind json: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrGetRequest())
	}
	input.ID = id

	// validation input user
	err = h.Validate.Struct(input)
	if err != nil {
		return utils.HttpError(c, response_mapper.FormatValidationError(err))
	}

	err = h.Service.Update(ctx, input)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, err)
	}

	return c.Status(http.StatusOK).JSON(response_mapper.RenderStruct(http.StatusOK, response_mapper.MultiLanguages{
		ID: "Anggota Tim Berhasil Diperbarui",
		EN: "Team Member Updated Successfully",
	}))
}

func (h *TeamMemberHandler) GetList(c *fiber.Ctx) error {
	var (
		opName = "TeamMemberController-GetList"
		ctx    = c.Context()
		input  dto.TeamMemberListReq
		err    error
	)

	err = c.QueryParser(&input)
	if err != nil {
		h.Logger.Errorf("%v error bind json: %v ", opName, err)
		return utils.HttpError(c, response_mapper.ErrGetRequest())
	}

	res, err := h.Service.GetList(ctx, input)
	if err != nil {
		h.Logger.Errorf("%v error: %v ", opName, err)
		return utils.HttpError(c, err)
	}

	return c.Status(http.StatusOK).JSON(response_mapper.RenderStruct(http.StatusOK, res))
}
