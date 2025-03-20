import StatusCard from "@/components/status-card";
import { useResetPasswordStore } from "@/stores/use-reset-password-store";
import { Navigate } from "react-router-dom";

const PasswordChangeStatusPage = () => {
  const { successResetPassword, openResetPasswordModal } =
    useResetPasswordStore();

  if (!openResetPasswordModal) {
    return <Navigate to="/auth/forgot-password" replace />;
  }

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      <StatusCard
        type={successResetPassword ? "success" : "error"}
        headerLabel={
          successResetPassword
            ? "Password Change Success"
            : "Password Change Failed"
        }
        message={
          successResetPassword
            ? "Your password has been successfully changed."
            : "Failed to change your password"
        }
        buttonLabel={successResetPassword ? "Go to login" : "Try again"}
        buttonHref={
          successResetPassword ? "/auth/login" : "/auth/reset-password"
        }
      />
    </section>
  );
};

export default PasswordChangeStatusPage;
