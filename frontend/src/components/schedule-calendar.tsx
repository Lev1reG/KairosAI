import { Calendar, dateFnsLocalizer, Event } from "react-big-calendar";
import "react-big-calendar/lib/css/react-big-calendar.css";
import { parse, format, startOfWeek, getDay } from "date-fns";
import id from "date-fns/locale/id";
import { useMemo, useState } from "react";
import { useSchedule, useScheduleDetail } from "@/hooks/use-schedule";
import { parseISO } from "date-fns";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

export type CalendarScheduleEvent = Event & {
  id: string;
  title: string;
  start: Date;
  end: Date;
  status: "scheduled" | "completed" | "canceled";
};

const locales = { id };

const localizer = dateFnsLocalizer({
  format,
  parse,
  startOfWeek: () => startOfWeek(new Date(), { weekStartsOn: 1 }),
  getDay,
  locales,
});

export default function ScheduleCalendar() {
  const { data, isLoading } = useSchedule();

  const [selectedId, setSelectedId] = useState<string | null>(null);
  const { data: scheduleDetail, isLoading: isDetailLoading } =
    useScheduleDetail(selectedId || undefined);

  const events: CalendarScheduleEvent[] = useMemo(() => {
    const schedules = data?.pages.flat() ?? [];
    return schedules.map((item) => ({
      id: item.id,
      title: item.title || "(untitled)",
      start: parseISO(item.start_time),
      end: parseISO(item.end_time),
      status: item.status,
    }));
  }, [data]);

  if (isLoading) return <div className="p-4">Loading calendar...</div>;

  return (
    <div className="pl-[304px] pr-6 pt-25 pb-6">
      <Calendar
        localizer={localizer}
        culture="id"
        events={events}
        startAccessor="start"
        endAccessor="end"
        onSelectEvent={(event: CalendarScheduleEvent) => {
          setSelectedId(event.id);
        }}
        components={{
          event: ({ event }: { event: CalendarScheduleEvent }) => (
            <div>
              <div className="font-medium">
                {format(event.start, "HH:mm")} - {format(event.end, "HH:mm")}
              </div>
              <div>{event.title}</div>
            </div>
          ),
        }}
        style={{ height: "calc(100vh - 48px)" }}
      />
      <Dialog open={!!selectedId} onOpenChange={() => setSelectedId(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>
              {isDetailLoading
                ? "Loading..."
                : scheduleDetail?.data?.title || "No Title"}
            </DialogTitle>
          </DialogHeader>
          {!isDetailLoading && scheduleDetail && (
            <div className="text-sm space-y-2">
              <p>
                <strong>Time:</strong>{" "}
                {format(
                  parseISO(scheduleDetail?.data?.start_time || ""),
                  "dd MMM yyyy HH:mm"
                )}{" "}
                -{" "}
                {format(
                  parseISO(scheduleDetail?.data?.end_time || ""),
                  "dd MMM yyyy HH:mm"
                )}
              </p>
              <p>
                <strong>Description:</strong>{" "}
                {scheduleDetail?.data?.description || "-"}
              </p>
            </div>
          )}
        </DialogContent>
      </Dialog>
    </div>
  );
}
