import { chatToAI } from "@/api/schedule";
import { queryClient } from "@/lib/query-client";
import { useMutation } from "@tanstack/react-query";
import { useCancelSchedule, useCreateSchedule } from "./use-schedule";
import toast from "react-hot-toast";

export const useAiModel = () => {
	const createSchedule = useCreateSchedule();
	const cancelSchedule = useCancelSchedule();

	return useMutation({
		mutationKey: ["chat"],
		mutationFn: chatToAI,

		onSuccess: (data) => {
			const intent = data?.queryResult.intent.displayName;

			switch (intent) {
				case "create-schedule":
					createSchedule.mutate(
						{
							title: data?.queryResult.parameters.title || "Default Title",
							start_time:
								data?.queryResult.parameters["date-time"][0]?.date_time ||
								new Date().toISOString(),
							description:
								data?.queryResult.parameters.description?.[0] || undefined,
						},
						{
							onSuccess: () => {
								queryClient.invalidateQueries({ queryKey: ["schedule"] });
							},
						}
					);
					break;
				case "delete-schedule": {
					const scheduleId = data?.queryResult.parameters.id;

					if (scheduleId) {
						cancelSchedule.mutate(scheduleId, {
							onSuccess: () => {
								queryClient.invalidateQueries({ queryKey: ["schedule"] });
							},
						});
					} else {
						toast.error("No schedule ID provided by AI to cancel.");
					}
					break;
				}
			}
		},
	});
};
