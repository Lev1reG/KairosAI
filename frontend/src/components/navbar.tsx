import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
} from "./ui/navigation-menu";
import { Button } from "./ui/button";
import { Link, useLocation } from "react-router-dom";

const Navbar = () => {
  const location = useLocation();

  const resolveVariant = (path: string) => {
    if (location.pathname == path) {
      return "navbarActive";
    }
    return "navbar";
  };

  return (
    <nav className="bg-green-500 px-8 py-4">
      <div className="container mx-auto flex justify-end items-center">
        <NavigationMenu>
          <NavigationMenuList className="gap-8">
            <NavigationMenuItem>
              <Link to="/auth/login">
                <Button variant={resolveVariant("/auth/login")} size="lg">
                  Login
                </Button>
              </Link>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <Link to="/auth/register">
                <Button variant={resolveVariant("/auth/register")} size="lg">
                  Register
                </Button>
              </Link>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>
      </div>
    </nav>
  );
};

export default Navbar;
