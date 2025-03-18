import { getCurrentUser, login } from "@/api/auth";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";

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

    onError: () => {
      toast.error("Failed to log in. Please try again.", { id: "loginToast" });
    },
  });
};

export const useCurrentUser = () => {
  return useQuery({
    queryKey: ["currentUser"],
    queryFn: getCurrentUser,
  });
};
