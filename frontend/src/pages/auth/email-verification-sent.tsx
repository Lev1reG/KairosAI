import StatusCard from "@/components/status-card";
import { useRegisterStore } from "@/stores/use-register-store";
import { Navigate } from "react-router-dom";

const EmailVerificationSentPage = () => {
  const { verificationSent } = useRegisterStore();

  if (!verificationSent) {
    return <Navigate to="/auth/register" replace />;
  }

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      <StatusCard
        type="email"
        headerLabel="Email Verification is Sent"
        message="A verification email has been sent to your email address."
      />
    </section>
  );
};

export default EmailVerificationSentPage;
