import { useRoutes } from "react-router-dom";
import publicRoutes from "./public-routes";
import NotFound from "../components/notfound";

export default function AppRoutes() {
  return useRoutes([publicRoutes, { path: "*", element: <NotFound /> }]);
}
