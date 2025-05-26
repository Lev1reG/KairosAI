import { getSchedule } from "@/api/schedule";
import { Schedule } from "@/types/schedule";
import { useInfiniteQuery } from "@tanstack/react-query";

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
