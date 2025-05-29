import { chatToAI } from "@/api/schedule";
import { queryClient } from "@/lib/query-client";
import { useMutation } from "@tanstack/react-query";

export const useAiModel = () => {
	return useMutation({
		mutationKey: ["chat"],
		mutationFn: chatToAI,

		onSuccess: () => {
			queryClient.invalidateQueries();
		},
	});
};
