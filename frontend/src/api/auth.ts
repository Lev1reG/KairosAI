import { api } from "./api";

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
  const response = await api.get("/auth/me");
  return response.data;
};

export const logout = async () => {
  const response = await api.post("/auth/logout");
  return response.data;
};

export const verifyEmail = async (token: string) => {
  const response = await api.post(`/auth/verify-email?token=${token}`);
  return response.data;
};

export const resendVerificationEmail = async (email: string) => {
  const response = await api.post("/auth/resend-verification", { email });
  return response.data;
};

export const forgotPassword = async (email: string) => {
  const response = await api.post("/auth/forgot-password", { email });
  return response.data;
};

export const resetPassword = async (token: string, new_password: string) => {
  const response = await api.post(`/auth/reset-password?token=${token}`, {
    new_password,
  });
  return response.data;
};
