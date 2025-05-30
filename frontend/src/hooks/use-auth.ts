import {
  forgotPassword,
  getCurrentUser,
  login,
  logout,
  oauthLogin,
  register,
  resetPassword,
  verifyEmail,
} from "@/api/auth";
import { useAuthStore } from "@/stores/use-auth-store";
import { useRegisterStore } from "@/stores/use-register-store";
import { useResetPasswordStore } from "@/stores/use-reset-password-store";
import { ApiResponse } from "@/types/api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

export const useResetPassword = () => {
  const { setSuccessResetPassword, setOpenResetPasswordModal } =
    useResetPasswordStore();
  const navigate = useNavigate();

  return useMutation({
    mutationKey: ["auth", "resetPassword"],
    mutationFn: ({
      token,
      newPassword,
    }: {
      token: string;
      newPassword: string;
    }) => resetPassword(token, newPassword),

    onMutate: () => {
      toast.loading("Resetting password...", { id: "resetPasswordToast" });
    },

    onSuccess: () => {
      toast.success("Password reset successfully!", {
        id: "resetPasswordToast",
      });
      setOpenResetPasswordModal(true);
      setSuccessResetPassword(true);
      navigate("/auth/reset-password-status", { replace: true });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message || "Failed to reset password. Please try again.";
      toast.error(errorMessage, { id: "resetPasswordToast" });
      setOpenResetPasswordModal(true);
      setSuccessResetPassword(false);
      navigate("/auth/reset-password-status", { replace: true });
    },
  });
};

export const useRequestPasswordReset = () => {
  const { setVerificationSent } = useResetPasswordStore();
  const navigate = useNavigate();

  return useMutation({
    mutationKey: ["auth", "requestPasswordReset"],
    mutationFn: ({ email }: { email: string }) => forgotPassword(email),

    onMutate: () => {
      toast.loading("Sending password reset email...", {
        id: "requestPasswordResetToast",
      });
    },

    onSuccess: () => {
      toast.success("Password reset email sent!", {
        id: "requestPasswordResetToast",
      });
      setVerificationSent(true);
      navigate("/auth/reset-password-email-sent", { replace: true });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message ||
        "Failed to send password reset email. Please try again.";
      toast.error(errorMessage, { id: "requestPasswordResetToast" });
    },
  });
};

export const useRegister = () => {
  const { setVerificationSent } = useRegisterStore();
  const navigate = useNavigate();

  return useMutation({
    mutationKey: ["auth", "register"],
    mutationFn: ({
      name,
      email,
      username,
      password,
    }: {
      name: string;
      email: string;
      username: string;
      password: string;
    }) => register(name, username, email, password),

    onMutate: () => {
      toast.loading("Registering...", { id: "registerToast" });
    },

    onSuccess: () => {
      toast.success("Registered successfully!", { id: "registerToast" });
      setVerificationSent(true);
      navigate("/auth/email-verification", { replace: true });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message || "Failed to register. Please try again.";
      toast.error(errorMessage, { id: "registerToast" });
    },
  });
};

export const useVerifyEmail = () => {
  return useMutation({
    mutationKey: ["auth", "verifyEmail"],
    mutationFn: ({ token }: { token: string }) => verifyEmail(token),

    onMutate: () => {
      toast.loading("Verifying email...", { id: "verifyEmailToast" });
    },

    onSuccess: () => {
      toast.success("Email verified successfully!", { id: "verifyEmailToast" });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message || "Failed to verify email. Please try again.";
      toast.error(errorMessage, { id: "verifyEmailToast" });
    },
  });
};

export const useLogin = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ["auth", "login"],
    mutationFn: ({ email, password }: { email: string; password: string }) =>
      login(email, password),

    onMutate: () => {
      toast.loading("Logging in...", { id: "loginToast" });
    },

    onSuccess: () => {
      toast.success("Logged in successfully!", { id: "loginToast" });
      queryClient.invalidateQueries({ queryKey: ["currentUser"] });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message || "Failed to log in. Please try again.";
      toast.error(errorMessage, { id: "loginToast" });
    },
  });
};

export const useOauthLogin = () => {
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  return useMutation({
    mutationKey: ["auth", "oauthLogin"],
    mutationFn: ({ code, state }: { code: string; state: string }) =>
      oauthLogin(code, state),

    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["currentUser"] });
      toast.success("Logged in successfully!", { id: "loginToast" });
      navigate("/", { replace: true });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message || "Failed to log in. Please try again.";
      toast.error(errorMessage, { id: "loginToast" });
    },
  });
};

export const useCurrentUser = () => {
  const { setUser, logout } = useAuthStore();

  const query = useQuery({
    queryKey: ["currentUser"],
    queryFn: getCurrentUser,
    retry: false,
  });

  useEffect(() => {
    if (query.isSuccess && query.data) {
      setUser(query.data?.data ?? null);
    }
  }, [query.isSuccess, query.data, setUser]);

  useEffect(() => {
    if (query.isError) {
      logout();
    }
  }, [query.isError, logout]);

  return query;
};

export const useLogout = () => {
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  return useMutation({
    mutationKey: ["auth", "logout"],
    mutationFn: logout,

    onMutate: () => {
      toast.loading("Logging out...", { id: "logoutToast" });
    },

    onSuccess: () => {
      toast.success("Logged out successfully!", { id: "logoutToast" });

      queryClient.removeQueries();

      navigate("/auth/login", { replace: true });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message || "Failed to log out. Please try again.";
      toast.error(errorMessage, { id: "logoutToast" });
    },
  });
};
