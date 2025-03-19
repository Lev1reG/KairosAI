import { ApiResponse } from "@/types/api";
import { api, handleApiError } from "./api";
import { User } from "@/types/user";
import { Token } from "@/types/token";

export const login = async (email: string, password: string) => {
  try {
    const response = await api.post<ApiResponse<User>>("/auth/login", {
      email,
      password,
    });
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const oauthLogin = async (code: string, state: string) => {
  try {
    const response = await api.get<ApiResponse<Token>>(
      `/auth/oauth/google/callback?state=${state}&code=${code}`,
    );
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const register = async (
  name: string,
  username: string,
  email: string,
  password: string,
) => {
  try {
    const response = await api.post<ApiResponse<User>>("/auth/register", {
      name,
      username,
      email,
      password,
    });
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const getCurrentUser = async () => {
  try {
    const response = await api.get<ApiResponse<User>>("/auth/me");
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const logout = async () => {
  try {
    const response = await api.post<ApiResponse<null>>("/auth/logout");
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const verifyEmail = async (token: string) => {
  try {
    const response = await api.post<ApiResponse<null>>(
      `/auth/verify-email?token=${token}`,
    );
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const resendVerificationEmail = async (email: string) => {
  try {
    const response = await api.post<ApiResponse<null>>(
      "/auth/resend-verification-email",
      {
        email,
      },
    );
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const forgotPassword = async (email: string) => {
  try {
    const response = await api.post<ApiResponse<null>>(
      "/auth/forgot-password",
      {
        email,
      },
    );
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};

export const resetPassword = async (token: string, new_password: string) => {
  try {
    const response = await api.post<ApiResponse<null>>(
      `/auth/reset-password?token=${token}`,
      {
        new_password,
      },
    );
    return response.data;
  } catch (error: unknown) {
    handleApiError(error);
  }
};
