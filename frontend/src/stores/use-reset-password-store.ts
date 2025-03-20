import { create } from "zustand";

interface ResetPasswordState {
  verificationSent: boolean;
  setVerificationSent: (value: boolean) => void;
  successResetPassword: boolean;
  setSuccessResetPassword: (value: boolean) => void;
  openResetPasswordModal: boolean;
  setOpenResetPasswordModal: (value: boolean) => void;
}

export const useResetPasswordStore = create<ResetPasswordState>((set) => ({
  verificationSent: false,
  setVerificationSent: (value) => set({ verificationSent: value }),
  successResetPassword: false,
  setSuccessResetPassword: (value) => set({ successResetPassword: value }),
  openResetPasswordModal: false,
  setOpenResetPasswordModal: (value) => set({ openResetPasswordModal: value }),
}));
