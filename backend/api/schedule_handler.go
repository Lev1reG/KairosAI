package api

import (
	"encoding/json"
	"net/http"

	"github.com/Lev1reG/kairosai-backend/api/middlewares"
	"github.com/Lev1reG/kairosai-backend/internal/services"
	"github.com/Lev1reG/kairosai-backend/internal/validator"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
	"go.uber.org/zap"
)

type ScheduleResponse struct {
	ID          int64  `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ScheduleHandler struct {
	scheduleService *services.ScheduleService
}

func NewScheduleHandler(scheduleService *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: scheduleService,
	}
}

func (h *ScheduleHandler) CreateSchedules(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok || userID == "" {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var reqBody services.CreateScheduleInput
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	reqBody.UserID = userID

	if err := validator.Validate.Struct(reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Validation failed: "+validator.ValidationFailed(err))
		return
	}

	schedules, err := h.scheduleService.CreateSchedule(r.Context(), reqBody)
	if err != nil {
		logger.Log.Error("Failed to create schedule", zap.Error(err))
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create schedule: "+err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, "Schedule created successfully", schedules)
}

func (h *ScheduleHandler) GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok || userID == "" {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	limit := utils.ParseQueryInt(r, "limit", 10)
	offset := utils.ParseQueryInt(r, "offset", 0)

	schedules, err := h.scheduleService.GetSchedulesByUser(r.Context(), userID, int32(limit), int32(offset))
	if err != nil {
		logger.Log.Error("Failed to get schedules", zap.Error(err))
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to get schedules")
	}

	total, err := h.scheduleService.CountSchedulesByUser(r.Context(), userID)
	if err != nil {
		logger.Log.Error("Failed to count schedules", zap.Error(err))
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to get schedule count")
		return
	}

	meta := utils.PaginationMeta{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}

	utils.SuccessPaginatedResponse(w, http.StatusOK, "Schedules retrieved successfully", schedules, meta)
}
