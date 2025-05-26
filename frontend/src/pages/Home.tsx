import ScheduleCalendar from "@/components/schedule-calendar";
import Sidebar from "@/components/sidebar";

const Home = () => {
  return (
    <div className="w-full min-h-screen flex flex-row">
      <Sidebar />
      <div className="flex-1">
        <ScheduleCalendar />
      </div>
    </div>
  );
};

export default Home;
