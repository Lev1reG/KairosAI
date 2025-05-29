import {
	createSchedule,
	getSchedule,
	getScheduleDetailed,
} from "@/api/schedule";
import { Schedule } from "@/types/schedule";
import { useInfiniteQuery, useMutation, useQuery } from "@tanstack/react-query";
import toast from "react-hot-toast";

const LIMIT = 5;

export const useSchedule = () => {
	return useInfiniteQuery({
		queryKey: ["schedule"],
		queryFn: async ({ pageParam = 0 }: { pageParam?: number }) => {
			const result = await getSchedule(LIMIT, pageParam);
			return result?.data ?? [];
		},
		initialPageParam: 0,
		getNextPageParam: (
			lastPage: Schedule[],
			allPages: Schedule[][]
		): number | undefined => {
			if (lastPage.length < LIMIT) return undefined;
			return allPages.length * LIMIT;
		},
	});
};

export const useScheduleDetail = (id?: string) => {
	return useQuery({
		queryKey: ["schedule", id],
		queryFn: () => getScheduleDetailed(id!),
		enabled: !!id,
	});
};

export const useCreateSchedule = () => {
	return useMutation({
		mutationKey: ["create-schedule"],
		mutationFn: createSchedule,
		onSuccess: () => {
			toast.success("Schedule created successfully!");
		},
		onError: () => {
			toast.error("You already have a schedule for this time!");
		},
	});
};
