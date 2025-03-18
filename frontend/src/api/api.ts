import { ApiResponse } from "@/types/api";
import axios from "axios";

const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL || "http://localhost:8080/api";

export const api = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

export const handleApiError = (error: unknown): never => {
  if (axios.isAxiosError(error)) {
    throw (
      (error.response?.data as ApiResponse<null>) ?? {
        status: "error",
        message: "An unknown error occurred.",
        code: 500,
      }
    );
  }
  throw {
    status: "error",
    message: "Something went wrong.",
    code: 500,
  } as ApiResponse<null>;
};
