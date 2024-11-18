import React, { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";

const AuthWrapper = ({ children, isAuthenticated }) => {
  const [isLoading, setIsLoading] = useState(true);
  const [authenticated, setAuthenticated] = useState(false);

  useEffect(() => {
    const checkAuth = async () => {
      const result = await isAuthenticated();
      setAuthenticated(result);
      setIsLoading(false);
    };
    checkAuth();
  }, []);

  if (isLoading) {
    return <div>Loading...</div>; // Optional: Show a loading indicator
  }

  return authenticated ? children : <Navigate to="/login" />;
};

export default AuthWrapper;