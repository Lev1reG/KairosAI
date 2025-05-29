import { ApiResponse } from "@/types/api";
import { api, handleApiError } from "./api";
import { Schedule } from "@/types/schedule";
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

export const chatToAI = async (data: AiServerRequest) => {
	const AI_SERVER_URL = import.meta.env.VITE_AI_SERVER_URL;
	try {
		const response = await api.post<AiServerResponse>(AI_SERVER_URL, data, {
			headers: {
				Authorization: `Bearer ${import.meta.env.VITE_AI_SERVER_TOKEN}`,
				"Content-Type": "application/json",
			},
		});
		return response.data;
	} catch (error: unknown) {
		handleApiError(error);
	}
};
