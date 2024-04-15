package ports

import (
	"github.com/DuckyMomo20012/go-todo/internal/tasks/app"
	"github.com/gofiber/fiber/v2"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type HTTPServer struct {
	service *app.TaskService
}

func NewHTTPServer(service *app.TaskService) HTTPServer {
	return HTTPServer{
		service: service,
	}
}

func (h HTTPServer) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.service.GetAllTasks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := make([]Task, 0)

	for _, task := range tasks {
		response = append(response, Task{
			Id:          task.UUID,
			Title:       &task.Title,
			Description: &task.Description,
		})
	}

	err = c.Status(fiber.StatusOK).JSON(response)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

func (h HTTPServer) CreateTask(c *fiber.Ctx) error {
	var request CreateTaskJSONRequestBody
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.service.CreateTask(c.Context(), *request.Title, *request.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = c.SendStatus(fiber.StatusCreated)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

func (h HTTPServer) DeleteTask(c *fiber.Ctx, id openapi_types.UUID) error {
	err := h.service.DeleteTask(c.Context(), id.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = c.SendStatus(fiber.StatusNoContent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

func (h HTTPServer) GetOneTask(c *fiber.Ctx, id openapi_types.UUID) error {
	task, err := h.service.GetOneTask(c.Context(), id.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if task == nil {
		err := c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	err = c.Status(fiber.StatusOK).JSON(Task{
		Id:          task.UUID,
		Title:       &task.Title,
		Description: &task.Description,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

func (h HTTPServer) UpdateTask(c *fiber.Ctx, id openapi_types.UUID) error {
	var request UpdateTaskJSONRequestBody
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task := app.Task{
		UUID:        id.String(),
		Title:       *request.Title,
		Description: *request.Description,
	}

	err := h.service.UpdateTask(c.Context(), id.String(), &task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = c.SendStatus(fiber.StatusOK)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}
