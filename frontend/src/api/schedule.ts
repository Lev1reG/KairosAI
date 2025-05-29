import { ApiResponse } from "@/types/api";
import { api, handleApiError } from "./api";
import { CreateScheduleInput, Schedule } from "@/types/schedule";
import { AiServerRequest, AiServerResponse } from "@/types/ai";

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

export const createSchedule = async (data: CreateScheduleInput) => {
	try {
		const response = await api.post<ApiResponse<Schedule>>("/schedules", data);
		return response.data;
	} catch (error: unknown) {
		handleApiError(error);
	}
};

export const chatToAI = async (data: AiServerRequest) => {
	try {
		const response = await api.post<AiServerResponse>("/chat", data);
		return response.data;
	} catch (error: unknown) {
		handleApiError(error);
	}
};

export const cancelSchedule = async (id: string) => {
	try {
		const response = await api.delete<ApiResponse<null>>(
			`/schedules/${id}/cancel`
		);
		return response.data;
	} catch (error: unknown) {
		handleApiError(error);
	}
};
