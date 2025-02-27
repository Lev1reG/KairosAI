import { useRoutes } from "react-router-dom";
import publicRoutes from "./public-routes";
import NotFound from "../components/notfound";
import authRoutes from "./auth-routes";

export default function AppRoutes() {
  return useRoutes([
    publicRoutes,
    authRoutes,
    { path: "*", element: <NotFound /> },
  ]);
}
