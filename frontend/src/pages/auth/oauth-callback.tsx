import StatusCard from "@/components/status-card";
import { useOauthLogin } from "@/hooks/use-auth";
import { useEffect } from "react";
import { useSearchParams } from "react-router-dom";

const OAuthCallback = () => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const state = searchParams.get("state");
  const { mutate: loginWithOauth, isPending, data } = useOauthLogin();

  useEffect(() => {
    if (code && state) {
      loginWithOauth({ code, state });
    }
  }, [code, state, loginWithOauth]);

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      {isPending ? (
        <p>Logging in...</p>
      ) : (
        <StatusCard
          type="error"
          headerLabel={code && state ? "Failed to log in" : "Forbidden request"}
          message={
            code && state
              ? data?.message || "Try again"
              : "You are not allowed to access this page"
          }
          buttonHref="/auth/login"
          buttonLabel="Back to login page"
        />
      )}
    </section>
  );
};

export default OAuthCallback;
