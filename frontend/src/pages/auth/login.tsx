import InputField from "../../components/inputField";
import googleLogo from "../../assets/Google.svg";
import { User, Lock } from "lucide-react";
import { Link } from "react-router-dom";


const LoginPage = () => {
  return (
    <div className="bg-gray-50">
      <h1 className=" text-2xl font-bold text-center mb-2 py-10">Hi, Welcome Back!</h1>
      <div className="flex items-center justify-center min-h-screen py-12">
        <div className="bg-white p-20 rounded-2xl shadow-lg w-full max-w-2xl h-150 space-y-6 py-25">
        <h2 className="text-center text-2xl font-semibold mb-6">Login to your account</h2>
        <form className="space-y-4 ">
          <InputField icon={<User size={18} />} placeholder="Username" />
          <InputField icon={<Lock size={18} />} placeholder="Password" type="password" />
          <a href="#" className="text-gray-500 hover:uderline px-3">Forgot Password?</a>
          <button type="submit" className="w-full py-2 rounded-full text-black font-medium cursor-pointer" style={{backgroundColor: "#A3C6C4"}}>Login</button>
          <button type="button" className="w-full py-2 bg-transparent border text-black font-medium rounded-full flex items-center justify-center space-x-2 hover:bg-gray-100 transition duration-200"
          style={{borderColor: "#A3C6C4", borderWidth: "1px"}}>
            <img src={googleLogo}
                alt="Google"
                className="w-5 h-5"/>
            <span>Sign in with Google</span>
          </button>
        </form>    

        <p className="text-center mt-6 text-sm">
            Don't have an account?{" "}
            <Link to="/auth/register" className="font-semibold hover:underline">
              Register
            </Link>
          </p>  

        </div>
      </div>
      
    </div>
  );
};

export default LoginPage;
