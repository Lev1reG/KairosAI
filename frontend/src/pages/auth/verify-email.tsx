import StatusCard from "@/components/status-card";
import { useVerifyEmail } from "@/hooks/use-auth";
import { useEffect } from "react";
import { useSearchParams } from "react-router-dom";

const VerifyEmailPage = () => {
  const [searchParams] = useSearchParams();
  const token = searchParams.get("token");
  const { mutate: verifyEmail, isPending, error, isSuccess } = useVerifyEmail();

  useEffect(() => {
    if (token) {
      verifyEmail({ token });
    }
  }, [token, verifyEmail]);

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      {isPending ? (
        <p>Verifying...</p>
      ) : (
        <StatusCard
          type={isSuccess ? "success" : "error"}
          headerLabel={
            token && isSuccess
              ? "Successfully verified your email"
              : token && !isSuccess
                ? "Failed to verify your email"
                : "Forbidden request"
          }
          message={
            token && isSuccess
              ? "You can now login to your account"
              : token && !isSuccess
                ? error?.message || "Try again"
                : "You are not allowed to access this page"
          }
          buttonHref="/auth/login"
          buttonLabel="Back to login page"
        />
      )}
    </section>
  );
};

export default VerifyEmailPage;
