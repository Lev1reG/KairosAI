import { RouteObject } from "react-router-dom";
import LoginPage from "../pages/auth/login";
import RegisterPage from "../pages/auth/register";
import MainLayout from "../layouts/main-layout";
import ForgotPasswordPage from "@/pages/auth/forgot-password";
import ResetPasswordPage from "@/pages/auth/reset-password";
import PasswordChangeSuccessPage from "@/pages/auth/password-change-success";
import PasswordChangeFailedPage from "@/pages/auth/password-change-failed";

const authRoutes: RouteObject = {
  path: "/auth",
  element: <MainLayout />,
  children: [
    { path: "login", element: <LoginPage /> },
    { path: "register", element: <RegisterPage /> },
    { path: "forgot-password", element: <ForgotPasswordPage /> },
    { path: "reset-password", element: <ResetPasswordPage /> },
    { path: "reset-password/success", element: <PasswordChangeSuccessPage /> },
    { path: "reset-password/failed", element: <PasswordChangeFailedPage />}
  ],
};

export default authRoutes;
