package handlers

import (
	"net/http"
	"strconv"

	"github.com/AndresOsorio0710/BackendGoCiCd/internals/application/services"
	"github.com/AndresOsorio0710/BackendGoCiCd/internals/core/entities"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service services.ITaskService
}

func NewTaskHandler(service services.ITaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) RegisterRoutes(router *gin.Engine) {
	tasks := router.Group("/api/tasks")
	{
		tasks.POST("/", h.CreateTask)
		tasks.GET("/", h.GetAllTasks)
		tasks.GET("/:id", h.GetTaskByID)
	}
}

// CreateTask godoc
// @Summary Crear una nueva tarea
// @Description Crea una nueva tarea y la guarda en la base de datos
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body entities.Task true "Datos de la tarea"
// @Success 201 {object} entities.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/ [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetAllTasks godoc
// @Summary Obtener todas las tareas
// @Description Retorna todas las tareas registradas
// @Tags tasks
// @Produce json
// @Success 200 {array} entities.Task
// @Failure 500 {object} map[string]string
// @Router /tasks/ [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID godoc
// @Summary Obtener una tarea por ID
// @Description Retorna una tarea específica dado su ID
// @Tags tasks
// @Produce json
// @Param id path int true "ID de la tarea"
// @Success 200 {object} entities.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	task, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}
