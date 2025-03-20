import * as z from "zod";

import CardWrapper from "@/components/card/card-wrapper";
import { useForm } from "react-hook-form";
import { ResetPasswordSchema } from "@/schemas";
import { zodResolver } from "@hookform/resolvers/zod";
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
import { useResetPassword } from "@/hooks/use-auth";
import { useSearchParams } from "react-router-dom";
import StatusCard from "@/components/status-card";

const ResetPasswordPage = () => {
  const resetPassword = useResetPassword();

  const [searchParams] = useSearchParams();
  const token = searchParams.get("token");

  const form = useForm<z.infer<typeof ResetPasswordSchema>>({
    resolver: zodResolver(ResetPasswordSchema),
    defaultValues: {
      password: "",
      confirmPassword: "",
    },
  });

  const onSubmit = (data: z.infer<typeof ResetPasswordSchema>) => {
    resetPassword.mutate({
      token: token ?? "",
      newPassword: data.password,
    });
  };

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      {token ? (
        <CardWrapper headerLabel="Change your password">
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
              <div className="space-y-4">
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>New Password</FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          type="password"
                          placeholder="Enter your new password"
                          className="bg-neutral-100"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="confirmPassword"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Confirm Password</FormLabel>
                      <FormControl>
                        <Input
                          {...field}
                          type="password"
                          placeholder="Confirm your password"
                          className="bg-neutral-100"
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <Button variant="submit" type="submit" className="w-full">
                Reset password
              </Button>
            </form>
          </Form>
        </CardWrapper>
      ) : (
        <StatusCard
          type="error"
          headerLabel="Forbidden Request"
          message="You are not authorized to reset your password."
          buttonLabel="Go back"
          buttonHref="/auth/login"
        />
      )}
    </section>
  );
};

export default ResetPasswordPage;
