import Sidebar from "@/components/sidebar";
import { useCurrentUser } from "@/hooks/use-auth";

const Home = () => {
  const { data: user, isPending } = useCurrentUser();

  if (isPending) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <Sidebar />
      <pre>{JSON.stringify(user, null, 2)}</pre>
    </div>
  );
};

export default Home;
