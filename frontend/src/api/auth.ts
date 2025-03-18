import { ApiResponse } from "@/types/api";
import { api } from "./api";
import { User } from "@/types/user";

export const login = async (email: string, password: string) => {
  const response = await api.post("/auth/login", { email, password });
  return response.data;
};

export const register = async (
  name: string,
  username: string,
  email: string,
  password: string,
) => {
  const response = await api.post("/auth/register", {
    name,
    username,
    email,
    password,
  });
  return response.data;
};

export const getCurrentUser = async () => {
  const response = await api.get<ApiResponse<User>>("/auth/me");
  return response.data;
};

export const logout = async () => {
  const response = await api.post<ApiResponse<null>>("/auth/logout");
  return response.data;
};

export const verifyEmail = async (token: string) => {
  const response = await api.post<ApiResponse<null>>(`/auth/verify-email?token=${token}`);
  return response.data;
};

export const resendVerificationEmail = async (email: string) => {
  const response = await api.post<ApiResponse<null>>("/auth/resend-verification", { email });
  return response.data;
};

export const forgotPassword = async (email: string) => {
  const response = await api.post<ApiResponse<null>>("/auth/forgot-password", { email });
  return response.data;
};

export const resetPassword = async (token: string, new_password: string) => {
  const response = await api.post<ApiResponse<null>>(`/auth/reset-password?token=${token}`, {
    new_password,
  });
  return response.data;
};
