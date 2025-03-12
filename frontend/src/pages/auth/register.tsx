import InputField from "../../components/inputField";
import googleLogo from "../../assets/Google.svg";
import { Mail, User, Lock } from "lucide-react";
import { Link } from "react-router-dom";

const RegisterPage = () => {
  return (
    <>
      <div className="flex items-center justify-center min-h-screen bg-gray-50 py-12">
        <div className="bg-white p-20 rounded-2xl shadow-lg w-full max-w-2xl h-150 space-y-6 py-25">
          <h2 className="text-center text-2xl font-semibold mb-6">
            Create your account
          </h2>

          <form className="space-y-4 ">
            <InputField icon={<Mail size={18} />} placeholder="Email" />
            <InputField icon={<User size={18} />} placeholder="Username" />
            <InputField
              icon={<Lock size={18} />}
              placeholder="Password"
              type="password"
            />
            <InputField
              icon={<Lock size={18} />}
              placeholder="Confirm Password"
              type="password"
            />

            <button
              className="w-full py-2 rounded-full text-black font-medium cursor-pointer"
              style={{ backgroundColor: "#A3C6C4" }}
            >
              Register
            </button>
            <button
              type="button"
              className="w-full py-2 bg-transparent border text-black font-medium rounded-full flex items-center justify-center space-x-2 hover:bg-gray-100 transition duration-200"
              style={{ borderColor: "#A3C6C4", borderWidth: "1px" }}
            >
              <img src={googleLogo} alt="Google" className="w-5 h-5" />
              <span>Register with Google</span>
            </button>
          </form>

          <p className="text-center my-4">
            Already have an account?{" "}
            <Link
              to="/auth/login"
              className="font-semibold cursor-pointer text-gray-700 hover:underline"
            >
              Login
            </Link>
          </p>
        </div>
      </div>
    </>
  );
};

export default RegisterPage;
