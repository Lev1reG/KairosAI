import { CalendarCheck } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { format, parseISO } from "date-fns";
import { useMemo } from "react";
import { useSchedule } from "@/hooks/use-schedule"; // adjust to your file
import { Schedule } from "@/types/schedule";
import { useToogleSidebarStore } from "@/stores/toogle-sidebar-store";

export default function Sidebar() {
  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, isLoading } =
    useSchedule();

  const { sidebarState, toggle } = useToogleSidebarStore();

  const allSchedules = useMemo(() => {
    return data?.pages.flat() ?? [];
  }, [data]);

  const groupedSchedules = useMemo(() => {
    const map: Record<string, Schedule[]> = {};
    for (const schedule of allSchedules) {
      const date = format(parseISO(schedule.start_time), "yyyy-MM-dd");
      if (!map[date]) map[date] = [];
      map[date].push(schedule);
    }
    return map;
  }, [allSchedules]);

  return (
    <div className="fixed flex flex-col gap-10 px-4 py-25 w-[280px] bg-background border-r">
      <h1 className="text-2xl font-bold text-center">KairosAI</h1>

      <Button
        onClick={toggle}
        variant="navbarActive"
        className="font-semibold text-bold text-white cursor-pointer"
      >
        {sidebarState === "chat" ? "View Calendar" : "View Chat"}
      </Button>

      <div>
        <div className="flex items-center justify-between mb-2">
          <div className="flex items-center gap-2 font-semibold text-lg">
            <CalendarCheck size={18} />
            <span>Schedules</span>
          </div>
          <span className="text-lg">⌄</span>
        </div>

        <Card className="p-4 space-y-4 max-h-[400px] overflow-y-auto">
          {isLoading && (
            <p className="text-sm text-muted-foreground">Loading...</p>
          )}

          {Object.keys(groupedSchedules).length === 0 ? (
            <p className="text-sm text-muted-foreground italic">
              No schedules available
            </p>
          ) : (
            Object.entries(groupedSchedules).map(([date, items]) => (
              <div key={date}>
                <p className="text-sm text-muted-foreground font-semibold mb-1">
                  {format(new Date(date), "dd MMM yyyy")}
                </p>
                <ul className="space-y-1">
                  {items.map((item, i) => (
                    <li
                      key={item.id ?? i}
                      className="text-sm flex flex-col gap-0.5"
                    >
                      <div className="flex justify-between items-center">
                        <span className="truncate font-medium">
                          {item.title || "(untitled)"}
                        </span>
                        <span
                          className={`text-xs font-semibold rounded px-2 py-0.5 ${
                            item.status === "scheduled"
                              ? "bg-blue-100 text-blue-800"
                              : item.status === "completed"
                              ? "bg-green-100 text-green-800"
                              : "bg-red-100 text-red-800"
                          }`}
                        >
                          {item.status}
                        </span>
                      </div>
                      <p className="text-xs text-muted-foreground">
                        {format(parseISO(item.start_time), "HH:mm")} -{" "}
                        {format(parseISO(item.end_time), "HH:mm")}
                      </p>
                    </li>
                  ))}
                </ul>
              </div>
            ))
          )}

          {hasNextPage && (
            <Button
              size="sm"
              variant="ghost"
              onClick={() => fetchNextPage()}
              disabled={isFetchingNextPage}
              className="w-full text-xs text-muted-foreground mt-2"
            >
              {isFetchingNextPage ? "Loading more..." : "Load more"}
            </Button>
          )}
        </Card>
      </div>
    </div>
  );
}
