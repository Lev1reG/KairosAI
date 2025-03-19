import * as z from "zod";

import AuthCard from "@/components/auth/auth-card";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { LoginSchema } from "@/schemas";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
import { useLogin } from "@/hooks/use-auth";
import GoogleLoginButton from "@/components/auth/google-login-button";

const LoginPage = () => {
  const loginMutation = useLogin();

  const form = useForm<z.infer<typeof LoginSchema>>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = (data: z.infer<typeof LoginSchema>) => {
    loginMutation.mutate(data);
  };

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      <div className="flex flex-col items-center gap-10">
        <h1 className="text-4xl font-bold text-black">Hi, Welcome Back!</h1>
        <AuthCard headerLabel="Login to your account" type="login">
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
              <div className="space-y-4">
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Email</FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          type="email"
                          placeholder="Enter your email"
                          className="bg-neutral-100"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Password</FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          type="password"
                          placeholder="Enter your password"
                          className="bg-neutral-100"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <Link to="/auth/forgot-password">
                  <span className="font-semibold text-sm text-neutral-500 hover:underline cursor-pointer">
                    Forgot password?
                  </span>
                </Link>
              </div>
              <div className="flex flex-col space-y-2">
                <Button
                  variant="submit"
                  type="submit"
                  className="w-full"
                  disabled={loginMutation.isPending}
                >
                  {loginMutation.isPending ? "Loading..." : "Login"}
                </Button>
                <GoogleLoginButton />   
              </div>
            </form>
          </Form>
        </AuthCard>
      </div>
    </section>
  );
};

export default LoginPage;
