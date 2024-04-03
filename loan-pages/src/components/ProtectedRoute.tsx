import { useEffect, useState } from "react";
import { Navigate, useLocation } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [loading, setLoading] = useState(true);
  const location = useLocation();

  useEffect(() => {
    const checkLoginStatus = async () => {
      try {
        const response = await fetch("http://localhost:8000/api/user", {
          method: "GET",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
        });
        setIsLoggedIn(response.ok);
      } catch (error) {
        console.error("Error checking login status:", error);
        setIsLoggedIn(false);
      } finally {
        setLoading(false);
      }
    };

    checkLoginStatus();
  }, []);

  // Jika masih dalam proses memeriksa status login, tampilkan indikator loading
  if (loading) {
    return <div>Loading...</div>;
  }

  // Jika pengguna belum login dan bukan di halaman login, arahkan ke halaman login
  if (!isLoggedIn && location.pathname !== "/login") {
    return <Navigate to="/login" />;
  }

  // Jika pengguna sudah login, arahkan langsung ke halaman rumah
  if (isLoggedIn && location.pathname === "/login") {
    return <Navigate to="/" />;
  }

  // Jika pengguna sudah login atau sedang di halaman login, tampilkan konten children
  return children;
};

export default ProtectedRoute;
