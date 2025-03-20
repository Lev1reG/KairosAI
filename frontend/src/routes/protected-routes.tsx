import { RouteObject } from "react-router-dom";
import Home from "../pages/Home";
import ProtectedLayout from "@/layouts/protected-layout";

const protectedRoutes: RouteObject = {
  path: "/",
  element: <ProtectedLayout />,
  children: [{ index: true, element: <Home /> }],
};

export default protectedRoutes;
