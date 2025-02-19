import { Link, Outlet } from "react-router-dom";

const MainLayout = () => {
  return (
    <div>
      {/* Navigation */}
      <nav>
        <Link to="/">Home</Link> |<Link to="/about">About</Link> |
        <Link to="/dashboard">Dashboard</Link>
      </nav>

      {/* Page Content */}
      <main>
        <Outlet /> {/* This will render the child route component */}
      </main>

      {/* Footer */}
      <footer>
        <p>© 2025 KairosAI</p>
      </footer>
    </div>
  );
};

export default MainLayout;
