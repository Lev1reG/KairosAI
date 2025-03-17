import * as z from "zod";

import CardWrapper from "@/components/card/card-wrapper";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { ForgotPasswordSchema } from "@/schemas";
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

const ForgotPasswordPage = () => {
  const form = useForm<z.infer<typeof ForgotPasswordSchema>>({
    resolver: zodResolver(ForgotPasswordSchema),
    defaultValues: {
      email: "",
    },
  });

  const onSubmit = (data: z.infer<typeof ForgotPasswordSchema>) => {
    console.log(data);
  };

  return (
    <section className="w-full min-h-screen flex justify-center items-center bg-background-custom">
      <CardWrapper headerLabel="Forgot your password?">
        <div className="w-full flex justify-center">
          <img
            src="/src/assets/lock-icon.png"
            className="w-24 h-24 object-cover"
            draggable="false"
          />
        </div>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="mt-7 space-y-8"
          >
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
              <div className="flex flex-col space-y-2">
                <Button variant="submit" type="submit" className="w-full">
                  Send Reset Link
                </Button>
                <Link to="/auth/login">
                  <Button variant="ghost" className="w-full">
                    Back to Login
                  </Button>
                </Link>
              </div>
            </div>
          </form>
        </Form>
      </CardWrapper>
    </section>
  );
};

export default ForgotPasswordPage;
