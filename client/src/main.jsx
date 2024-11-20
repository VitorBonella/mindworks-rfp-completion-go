import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.jsx";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Login from "./pages/Login";
import AuthWrapper from "./components/AuthWrapper.jsx";
import Rfps from "./pages/RFP";
import Equipment from "./pages/Equipment";
import NewRfp from "./pages/NewRFP";

//const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:8000" : "";
const BASE_URL = "https://orange-waddle-67v949p953x6pg-5174.app.github.dev"
export default BASE_URL;

const isAuthenticated = async () => {
  const response = await fetch(BASE_URL + "/api/user", {
    headers: { "Content-Type": "application/json" },
    credentials: "include",
  });

  if (response.status !== 200) {
    return false
  }

  const content = await response.json();

  if (content.Name === "") {
    return false;
  }
  return true;
};

// Define routes with conditional rendering for authentication
const router = createBrowserRouter([
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/",
    element: (
      <AuthWrapper isAuthenticated={isAuthenticated}>
        <App />
      </AuthWrapper>
    ),
    children: [
      {
        path: "/rfps", // Matches the "/rfps" path
        element: <Rfps />,
      },
      {
        path: "/equipment", // Matches the "/equipment" path
        element: <Equipment />,
      },
      {
        path: "/new_rfp", // Matches the "/new_rfp" path
        element: <NewRfp />,
      },
      // Add a fallback route if needed
      {
        path: "/*", // Catch-all route for any unmatched paths
        element: <div>Page not found</div>,
      },
    ],
  },
]);
// Main render logic
createRoot(document.getElementById("root")).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);
