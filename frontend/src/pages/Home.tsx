import ChatPage from "@/components/chat-page";
import ScheduleCalendar from "@/components/schedule-calendar";
import Sidebar from "@/components/sidebar";
import { useToogleSidebarStore } from "@/stores/toogle-sidebar-store";

const Home = () => {
  const { sidebarState } = useToogleSidebarStore();

  return (
    <div className="w-full min-h-screen flex flex-row">
      <Sidebar />
      <div className="flex-1">
        {sidebarState === "calendar" ? <ScheduleCalendar /> : <ChatPage />}
      </div>
    </div>
  );
};

export default Home;
