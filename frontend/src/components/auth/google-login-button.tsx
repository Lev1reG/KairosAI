import { API_BASE_URL } from "@/api/api";
import { Button } from "../ui/button";
import { FcGoogle } from "react-icons/fc";

const GoogleLoginButton = () => {
  const handleGoogleLogin = () => {
    window.location.href = `${API_BASE_URL}/auth/oauth/google/login`;
  };

  return (
    <Button
      onClick={handleGoogleLogin}
      variant="submit"
      type="button"
      className="w-full"
    >
      <FcGoogle className="w-8 h-8" />
      Login with Google
    </Button>
  );
};

export default GoogleLoginButton;
