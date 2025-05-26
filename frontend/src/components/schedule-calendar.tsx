import { Calendar, dateFnsLocalizer, Event } from "react-big-calendar";
import "react-big-calendar/lib/css/react-big-calendar.css";
import { parse, format, startOfWeek, getDay } from "date-fns";
import id from "date-fns/locale/id";
import { useMemo } from "react";
import { useSchedule } from "@/hooks/use-schedule";
import { parseISO } from "date-fns";

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
    <div className="p-6 w-full">
      <Calendar
        localizer={localizer}
        culture="id"
        events={events}
        startAccessor="start"
        endAccessor="end"
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
    </div>
  );
}
