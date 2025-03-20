import { RouteObject } from "react-router-dom";
import LoginPage from "../pages/auth/login";
import RegisterPage from "../pages/auth/register";
import ForgotPasswordPage from "@/pages/auth/forgot-password";
import ResetPasswordPage from "@/pages/auth/reset-password";
import AuthLayout from "@/layouts/auth-layout";
import OAuthCallback from "@/pages/auth/oauth-callback";
import EmailVerificationSentPage from "@/pages/auth/email-verification-sent";
import VerifyEmailPage from "@/pages/auth/verify-email";
import EmailRequestResetPasswordSent from "@/pages/auth/email-request-reset-password-sent";
import PasswordChangeStatusPage from "@/pages/auth/password-change-status";

const authRoutes: RouteObject = {
  path: "/auth",
  element: <AuthLayout />,
  children: [
    { path: "login", element: <LoginPage /> },
    { path: "register", element: <RegisterPage /> },
    { path: "forgot-password", element: <ForgotPasswordPage /> },
    { path: "reset-password", element: <ResetPasswordPage /> },
    { path: "reset-password-status", element: <PasswordChangeStatusPage /> },
    { path: "reset-password-email-sent", element: <EmailRequestResetPasswordSent /> },
    { path: "email-verification", element: <EmailVerificationSentPage /> },
    { path: "verify-email", element: <VerifyEmailPage /> },
    { path: "google/callback", element: <OAuthCallback /> },
  ],
};

export default authRoutes;
