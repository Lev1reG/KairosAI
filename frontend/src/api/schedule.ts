import { ApiResponse } from "@/types/api";
import { api, handleApiError } from "./api";
import { Schedule } from "@/types/schedule";

export const getSchedule = async (limit: number = 10, offset: number) => {
  try {
    const response = await api.get<ApiResponse<Schedule[]>>("/schedules", {
      params: {
        limit,
        offset,
      },
    });
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const getScheduleDetailed = async (id: string) => {
  try {
    const response = await api.get<ApiResponse<Schedule>>(`/schedules/${id}`);
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};
