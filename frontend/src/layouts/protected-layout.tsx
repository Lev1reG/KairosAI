import { useCurrentUser } from "@/hooks/use-auth";
import { Navigate } from "react-router-dom";
import MainLayout from "./main-layout";

const ProtectedLayout = () => {
  const { data, isLoading, isError } = useCurrentUser();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError || !data) {
    return <Navigate to="/auth/login" replace />;
  }

  return <MainLayout />;
};

export default ProtectedLayout;
