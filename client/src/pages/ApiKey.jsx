import { useState } from "react";
import BASE_URL from "../main";

const RegisterApiKey = ({ HasApiKey }) => {
  const [apiKey, setApiKey] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    setMessage("");

    try {
      const response = await fetch(BASE_URL + "/api/apikey", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ api_key: apiKey }),
      });

      if (response.ok) {
        setMessage("API key registered successfully!");
        setApiKey("");
      } else {
        const errorData = await response.json();
        setMessage(`Error: ${errorData.message || "Something went wrong."}`);
      }
    } catch (error) {
      setMessage("Error: Unable to register API key.");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white shadow-md rounded-lg p-8 w-full max-w-md">
        <h1 className="text-2xl font-semibold mb-6 text-center">
          Register API Key
        </h1>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label
              htmlFor="api_key"
              className="block text-gray-700 font-medium mb-2"
            >
              API Key
            </label>
            <input
              type="password"
              id="api_key"
              className="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={apiKey}
              onChange={(e) => setApiKey(e.target.value)}
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-500 text-white font-medium py-2 px-4 rounded-lg hover:bg-blue-600 transition"
          >
            Register
          </button>
        </form>
        {message && (
          <div
            className={`mt-4 text-center font-medium ${
              message.startsWith("Error") ? "text-red-500" : "text-green-500"
            }`}
          >
            {message}
          </div>
        )}
        {HasApiKey && (
          <div
            className="mt-4 text-center font-medium 
              text-green-500"
          >
            {"You already have API Key"}
          </div>
        )}
      </div>
    </div>
  );
};

export default RegisterApiKey;
