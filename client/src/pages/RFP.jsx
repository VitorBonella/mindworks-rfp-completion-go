import React, { useState, useEffect } from "react";
import BASE_URL from "../main";

const RFPStatusCreated = "Created";
const RFPStatusProcessing = "Processing";
const RFPStatusFinished = "Finished";
const RFPStatusFinishedWithError = "Finished With Error";

const Rfp = () => {
  const [rfps, setRfps] = useState([]);

  // Fetch RFPs from the API
  useEffect(() => {
    fetch(BASE_URL + "/api/rfps", {
      credentials: "include",
    })
      .then((response) => response.json())
      .then((data) => setRfps(data))
      .catch((error) => console.error("Error fetching RFPs:", error));
  }, []);

  const reprocessRFP = (id) => {
    fetch(BASE_URL + "/api/rfp/reprocess", {
      credentials: "include",
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id }), // Sending the ID to the reprocess endpoint
    })
      .then((response) => {
        if (response.ok) {
          // Optionally, refetch the RFPs or update the status locally
          setRfps((prevRfps) =>
            prevRfps.map((rfp) =>
              rfp.id === id ? { ...rfp, status: RFPStatusProcessing } : rfp
            )
          );
        }
      })
      .catch((error) => console.error("Error reprocessing RFP:", error));
  };

  // Helper function to format date
  const formatDate = (date) => {
    if (!date) return "-";
    return new Date(date).toLocaleString(); // This will include both date and time
  };

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-2xl font-bold mb-4">RFP List</h1>
      <table className="table-auto w-full border-collapse text-left">
        <thead>
          <tr>
            <th className="px-4 py-2 border-b">Name</th>
            <th className="px-4 py-2 border-b">Creation Date</th>
            <th className="px-4 py-2 border-b">End Date</th>
            <th className="px-4 py-2 border-b">Status</th>
            <th className="px-4 py-2 border-b">Actions</th>
          </tr>
        </thead>
        <tbody>
          {rfps.map((rfp) => (
            <tr key={rfp.id}>
              <td className="px-4 py-2 border-b">{rfp.name}</td>
              <td className="px-4 py-2 border-b">
                {formatDate(rfp.creation_date)}
              </td>
              <td className="px-4 py-2 border-b">{formatDate(rfp.end_date)}</td>
              <td className="px-4 py-2 border-b">
                <span
                  className={`px-2 py-1 rounded-full ${
                    rfp.status === RFPStatusCreated
                      ? "bg-blue-100 text-blue-800"
                      : rfp.status === RFPStatusProcessing
                      ? "bg-yellow-100 text-yellow-800"
                      : rfp.status === RFPStatusFinished
                      ? "bg-green-100 text-green-800"
                      : rfp.status === RFPStatusFinishedWithError
                      ? "bg-red-100 text-red-800"
                      : ""
                  }`}
                >
                  {rfp.status}
                </span>
              </td>
              <td className="px-4 py-2 border-b">
                {(rfp.status === RFPStatusCreated ||
                  rfp.status === RFPStatusFinishedWithError ||
                  rfp.status === RFPStatusFinished) && (
                  <button
                    onClick={() => reprocessRFP(rfp.id)}
                    className="bg-blue-500 text-white py-1 px-3 rounded-md hover:bg-blue-600"
                  >
                    Reprocess
                  </button>
                )}
                {!(
                  rfp.status === RFPStatusCreated ||
                  rfp.status === RFPStatusFinishedWithError ||
                  rfp.status === RFPStatusFinished
                ) && <span className="text-gray-500">N/A</span>}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Rfp;
