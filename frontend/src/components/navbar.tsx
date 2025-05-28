import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
} from "./ui/navigation-menu";
import { Button } from "./ui/button";
import { Link, useLocation } from "react-router-dom";
import { useAuthStore } from "@/stores/use-auth-store";
import { useLogout } from "@/hooks/use-auth";

const Navbar = () => {
  const location = useLocation();
  const { isAuthenticated, user } = useAuthStore();

  const logoutMutation = useLogout();

  const resolveVariant = (path: string) => {
    if (location.pathname == path) {
      return "navbarActive";
    }
    return "navbar";
  };

  return (
    <nav className="z-10 w-full fixed bg-green-500 px-8 py-4">
      <div className="container mx-auto flex justify-end items-center">
        <NavigationMenu>
          <NavigationMenuList className="gap-8">
            <NavigationMenuItem>
              {isAuthenticated ? (
                <div className="flex items-center gap-4">
                  <p className="font-semibold text-lg">{user?.name || ""}</p>
                  <Button
                    onClick={() => logoutMutation.mutate()}
                    variant={resolveVariant("/")}
                    size="lg"
                  >
                    Logout
                  </Button>
                </div>
              ) : (
                <Link to="/auth/login">
                  <Button variant={resolveVariant("/auth/login")} size="lg">
                    Login
                  </Button>
                </Link>
              )}
            </NavigationMenuItem>
            {isAuthenticated ? null : (
              <NavigationMenuItem>
                <Link to="/auth/register">
                  <Button variant={resolveVariant("/auth/register")} size="lg">
                    Register
                  </Button>
                </Link>
              </NavigationMenuItem>
            )}
          </NavigationMenuList>
        </NavigationMenu>
      </div>
    </nav>
  );
};

export default Navbar;
