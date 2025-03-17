import { z } from "zod";

const passwordSchema = z
  .string()
  .min(8, { message: "Password must be at least 8 characters long" })
  .refine((password) => /[A-Z]/.test(password), {
    message: "Password must contain at least one uppercase letter",
  })
  .refine((password) => /[a-z]/.test(password), {
    message: "Password must contain at least one lowercase letter",
  })
  .refine((password) => /[0-9]/.test(password), {
    message: "Password must contain at least one number",
  });

export const usernameSchema = z
  .string()
  .min(3, { message: "Username must be at least 3 characters long" })
  .max(20, { message: "Username must be no more than 20 characters long" })
  .regex(/^[a-zA-Z0-9_.]+$/, {
    message:
      "Username can only contain letters, numbers, underscores, and dots",
  })
  .refine(
    (username) => !username.startsWith(".") && !username.startsWith("_"),
    {
      message: "Username cannot start with a dot (.) or underscore (_)",
    },
  )
  .refine((username) => !username.endsWith("."), {
    message: "Username cannot end with a dot (.)",
  });

export const LoginSchema = z.object({
  email: z.string().email({ message: "Invalid email format" }),
  password: passwordSchema,
});

export const RegisterSchema = z
  .object({
    email: z.string().email({ message: "Invalid email format" }),
    username: usernameSchema,
    password: passwordSchema,
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });
