import { getCurrentUser, login, logout, oauthLogin } from "@/api/auth";
import { useAuthStore } from "@/stores/use-auth-store";
import { ApiResponse } from "@/types/api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

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
    } else if (query.isError) {
      logout();
    }
  }, [query.isSuccess, query.isError, query.data, setUser, logout]);

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

      queryClient.invalidateQueries({ queryKey: ["currentUser"] });

      navigate("/auth/login", { replace: true });
    },

    onError: (error: ApiResponse<null>) => {
      const errorMessage =
        error.message || "Failed to log out. Please try again.";
      toast.error(errorMessage, { id: "logoutToast" });
    },
  });
};
