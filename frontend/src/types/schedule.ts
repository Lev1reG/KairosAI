export interface Schedule {
	id: string;
	user_id: string;
	title: string;
	description: string;
	start_time: string;
	end_time: string;
	status: "scheduled" | "canceled" | "completed";
	created_at: string;
	updated_at: string;
}
export interface CreateScheduleInput {
	title: string;
	description?: string;
	start_time: string;
	end_time?: string;
}
