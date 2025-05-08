package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Lev1reG/kairosai-backend/api/middlewares"
	"github.com/Lev1reG/kairosai-backend/internal/services"
	"github.com/Lev1reG/kairosai-backend/internal/validator"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
	"github.com/go-chi/chi/v5"
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

func (h *ScheduleHandler) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok || userID == "" {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	scheduleID := chi.URLParam(r, "id")
	if scheduleID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Schedule ID is required")
		return
	}

	var input services.UpdateScheduleInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validator.Validate.Struct(input); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Validation failed: "+validator.ValidationFailed(err))
		return
	}

	err := h.scheduleService.UpdateScheduleByID(r.Context(), userID, scheduleID, input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			utils.ErrorResponse(w, http.StatusNotFound, "Schedule not found")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update schedule: "+err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Schedule updated successfully", nil)
}

func (h *ScheduleHandler) CancelSchedule(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok || userID == "" {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	scheduleID := chi.URLParam(r, "id")
	if scheduleID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Schedule ID is required")
		return
	}

	err := h.scheduleService.CancelScheduleByID(r.Context(), userID, scheduleID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			utils.ErrorResponse(w, http.StatusNotFound, "Schedule not found or already canceled")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to cancel schedule")
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Schedule canceled successfully", nil)
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

func (h *ScheduleHandler) GetScheduleDetail(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok || userID == "" {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	scheduleID := chi.URLParam(r, "id")
	if scheduleID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Schedule ID is required")
		return
	}

	schedule, err := h.scheduleService.GetScheduleDetail(r.Context(), userID, scheduleID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to get schedule detail: s"+err.Error())
		return
	}

	if schedule == nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Schedule not found")
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Schedule retrieved successfully", schedule)
}
