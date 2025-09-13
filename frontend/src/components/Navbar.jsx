import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <nav className="bg-blue-600 text-white p-4 flex gap-6">
      <Link to="/">Products</Link>
      <Link to="/stock-in">Stock In</Link>
      <Link to="/stock-out">Stock Out</Link>
      <Link to="/inventory">Inventory</Link>
    </nav>
  );
}