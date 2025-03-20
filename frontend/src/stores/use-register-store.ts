import { create } from "zustand";

interface RegisterState {
  verificationSent: boolean;
  setVerificationSent: (value: boolean) => void;
  processVerification: boolean;
  setProcessVerification: (value: boolean) => void;
}

export const useRegisterStore = create<RegisterState>((set) => ({
  verificationSent: false,
  setVerificationSent: (value) => set({ verificationSent: value }),
  processVerification: false,
  setProcessVerification: (value) => set({ processVerification: value }),
}));
