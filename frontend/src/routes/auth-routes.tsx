import { RouteObject } from "react-router-dom";
import LoginPage from "../pages/auth/login";
import RegisterPage from "../pages/auth/register";
import MainLayout from "../layouts/main-layout";

const authRoutes: RouteObject = {
  path: "/auth",
  element: <MainLayout />,
  children: [
    { path: "login", element: <LoginPage /> },
    { path: "register", element: <RegisterPage /> },
  ],
};

export default authRoutes;
