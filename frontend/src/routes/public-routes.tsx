import { RouteObject } from "react-router-dom";
import MainLayout from "../layouts/main-layout";
import Home from "../pages/Home";

const publicRoutes: RouteObject = {
  path: "/",
  element: <MainLayout />,
  children: [{ index: true, element: <Home /> }],
};

export default publicRoutes;
