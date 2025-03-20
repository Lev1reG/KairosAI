import StatusCard from "@/components/status-card";
import { useResetPasswordStore } from "@/stores/use-reset-password-store";
import { Navigate } from "react-router-dom";

const EmailRequestResetPasswordSent = () => {
  const { verificationSent } = useResetPasswordStore();

  if (!verificationSent) {
    return <Navigate to="/auth/login" replace />;
  }

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      <StatusCard
        type="email"
        headerLabel="Reset Password Email Sent"
        message="A reset password email has been sent to your email address."
      />
    </section>
  );
};

export default EmailRequestResetPasswordSent;
