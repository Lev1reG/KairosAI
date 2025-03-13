import { Link } from "react-router-dom";
import CardWrapper from "../card/card-wrapper";

interface AuthCardProps {
  children: React.ReactNode;
  headerLabel: string;
  type: "login" | "register";
}

const AuthCard = ({ children, headerLabel, type }: AuthCardProps) => {
  const isLogin = type === "login";

  return (
    <CardWrapper
      headerLabel={headerLabel}
      footer={
        <>
          <div className="w-full font-semibold text-sm text-neutral-500 flex justify-center">
            {isLogin ? "Don’t have an account?" : "Already have an account?"}
            <Link to={isLogin ? "/auth/register" : "/auth/login"} className="ml-0.5">
              <span className="font-extrabold hover:underline cursor-pointer">
                {isLogin ? "Register" : "Login"}
              </span>
            </Link>
          </div>
        </>
      }
    >
      {children}
    </CardWrapper>
  );
};

export default AuthCard;
