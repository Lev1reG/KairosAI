import { create } from "zustand";

interface ToogleSidebarState {
  sidebarState: "chat" | "calendar";
  toggle: () => void;
}

export const useToogleSidebarStore = create<ToogleSidebarState>((set) => ({
  sidebarState: "chat",
  toggle: () =>
    set((state) => ({
      sidebarState: state.sidebarState === "chat" ? "calendar" : "chat",
    })),
}));
