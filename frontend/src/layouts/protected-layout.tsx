import { useCurrentUser } from "@/hooks/use-auth";
import { Navigate } from "react-router-dom";
import MainLayout from "./main-layout";
import { useAuthStore } from "@/stores/use-auth-store";

const ProtectedLayout = () => {
  const { isAuthenticated } = useAuthStore();
  const { isLoading } = useCurrentUser();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!isAuthenticated) {
    return <Navigate to="/auth/login" replace />;
  }

  return <MainLayout />;
};

export default ProtectedLayout;
