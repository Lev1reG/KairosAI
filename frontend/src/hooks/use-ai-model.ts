import { chatToAI } from "@/api/schedule";
import { queryClient } from "@/lib/query-client";
import { useMutation } from "@tanstack/react-query";
import { useCreateSchedule } from "./use-schedule";

export const useAiModel = () => {
	const createSchedule = useCreateSchedule();

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
			}
		},
	});
};
