import { ReactNode } from "react";
interface NavbarProps {
    children?: ReactNode;
}

const Navbar: React.FC<NavbarProps>= ({children}) => {
    return (
        <nav className="w-full h-16 flex justify-end items-center px-8" style={{backgroundColor: "#A3C6C4"}}>
            {children}
        </nav>
    );
};

export default Navbar;