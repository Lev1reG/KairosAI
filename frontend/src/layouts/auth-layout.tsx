import { useCurrentUser } from "@/hooks/use-auth";
import { Navigate } from "react-router-dom";
import MainLayout from "./main-layout";

const AuthLayout = () => {
  const { data: response, isLoading } = useCurrentUser();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  const user = response?.data;

  if (!!user) {
    return <Navigate to="/" replace />;
  }

  return <MainLayout />;
};

export default AuthLayout;
