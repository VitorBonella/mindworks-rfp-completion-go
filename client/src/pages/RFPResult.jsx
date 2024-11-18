import React, { useState, useEffect } from "react";
import BASE_URL from "../main";
import { useParams } from "react-router-dom";

function RFPResult() {
  const [rfpResults, setRfpResults] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [viewTable, setViewTable] = useState(false); // State for toggling the view
  const { id } = useParams();

  useEffect(() => {
    // Fetch data from the API
    fetch(BASE_URL + `/api/rfp/result?id=${id}`, {
      credentials: "include",
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch RFP results.");
        }
        return response.json();
      })
      .then((data) => {
        setRfpResults(data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [id]);

  const handleToggleView = () => {
    setViewTable((prevView) => !prevView);
  };

  if (loading) {
    return <div className="text-center mt-8 text-gray-500">Loading...</div>;
  }

  if (error) {
    return <div className="text-center mt-8 text-red-500">Error: {error}</div>;
  }

  // Helper function to render the checkmark or cross emoji
  const getStatusEmoji = (met) => {
    // Check for a valid boolean or met status (you can adjust this based on how your data is structured)
    return met === "Met" ? "✅" : "❌";
  };

  // Prepare data for table view
  const questions = [];
  const equipmentNames = Object.keys(rfpResults);

  // Extract questions from the data
  equipmentNames.forEach((productName) => {
    Object.entries(rfpResults[productName].Map).forEach(([questionKey, questionData]) => {
      if (!questions.find((q) => q.key === questionKey)) {
        questions.push({ key: questionKey, question: questionData.question });
      }
    });
  });

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">RFP Results</h1>

      {/* Toggle Button */}
      <button
        onClick={handleToggleView}
        className="mb-4 bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600"
      >
        {viewTable ? "Switch to Question View" : "Switch to Table View"}
      </button>

      {/* View Based on Toggle */}
      {viewTable ? (
        // Table View
        <div className="overflow-x-auto">
          <table className="min-w-full table-auto border-collapse border border-gray-200">
            <thead>
              <tr>
                <th className="border px-4 py-2">Question</th>
                {equipmentNames.map((equipmentName) => (
                  <th key={equipmentName} className="border px-4 py-2">
                    {equipmentName}
                  </th>
                ))}
              </tr>
            </thead>
            <tbody>
              {questions.map(({ key, question }) => (
                <tr key={key}>
                  <td className="border px-4 py-2">{question}</td>
                  {equipmentNames.map((equipmentName) => {
                    const questionData = rfpResults[equipmentName].Map[key];
                    // Ensure 'met' is a valid boolean indicating the requirement status
                    const met = questionData?.answer;
                    return (
                      <td key={equipmentName} className="border px-4 py-2 text-center">
                        {getStatusEmoji(met)}
                      </td>
                    );
                  })}
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        // Normal View
        <div>
          {Object.entries(rfpResults).map(([productName, productData]) => (
            <div key={productName} className="border rounded-lg p-4 mb-4 shadow-md">
              <h2 className="text-xl font-semibold mb-2">{productName}</h2>
              {Object.entries(productData.Map).map(([questionKey, questionData]) => (
                <div key={questionKey} className="bg-gray-50 p-4 mb-3 rounded-lg shadow-sm">
                  <p className="text-gray-700 font-medium">
                    <span className="font-semibold">Question:</span> {questionData.question}
                  </p>
                  <p className="text-green-600 font-medium">
                    <span className="font-semibold">Answer:</span> {questionData.answer}
                  </p>
                  <p className="text-gray-700">
                    <span className="font-semibold">Source:</span> {questionData.source}
                  </p>
                  <p className="text-gray-700">
                    <span className="font-semibold">Description:</span> {questionData.description}
                  </p>
                </div>
              ))}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default RFPResult;
