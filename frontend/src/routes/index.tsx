import { useRoutes } from "react-router-dom";
import NotFound from "../components/notfound";
import authRoutes from "./auth-routes";
import protectedRoutes from "./protected-routes";

export default function AppRoutes() {
  return useRoutes([
    protectedRoutes,
    authRoutes,
    { path: "*", element: <NotFound /> },
  ]);
}
