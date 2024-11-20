import { useState, useEffect } from "react";
import BASE_URL from "../main";
import { useParams } from "react-router-dom";
import { downloadAsCSV } from "../utils/DownloadCsv"; // Import the utility
import { downloadAsExcel } from "../utils/DownloadXlxs";

function RFPResult() {
  const [rfpResults, setRfpResults] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [viewTable, setViewTable] = useState(false);
  const [productToggles, setProductToggles] = useState({});
  const { id } = useParams();

  useEffect(() => {
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
        const initialToggles = Object.keys(data).reduce((acc, productName) => {
          acc[productName] = false;
          return acc;
        }, {});
        setProductToggles(initialToggles);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [id]);

  const handleGlobalToggleView = () => {
    setViewTable((prevView) => !prevView);
  };

  const handleProductToggle = (productName) => {
    setProductToggles((prevToggles) => ({
      ...prevToggles,
      [productName]: !prevToggles[productName],
    }));
  };

  const handleDownload = () => {
    downloadAsCSV("rfp_results.csv", rfpResults);
  };

  const handleDownloadExcel = () => {
    downloadAsExcel("rfp_results.xlsx", rfpResults);
  };

  const getRowBackgroundColor = (answer) => {
    if (answer === "Met") return "bg-green-100";
    if (answer === "Undefined") return "bg-yellow-100";
    return "bg-red-100";
  };

  if (loading) {
    return <div className="text-center mt-8 text-gray-500">Loading...</div>;
  }

  if (error) {
    return <div className="text-center mt-8 text-red-500">Error: {error}</div>;
  }

  const equipmentNames = Object.keys(rfpResults);

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">RFP Results</h1>
      {/* Global Toggle Button */}
      <button
        onClick={handleGlobalToggleView}
        className="mb-4 bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600"
      >
        {viewTable ? "Switch to Question View" : "Switch to Table View"}
      </button>
      {/* Download Button */}
      <button
        onClick={handleDownload}
        className="mb-4 ml-4 bg-green-500 text-white py-2 px-4 rounded-md hover:bg-green-600"
      >
        Download as CSV
      </button>
      <button
        onClick={handleDownloadExcel}
        className="mb-4 bg-green-500 text-white py-2 px-4 rounded-md hover:bg-green-600"
      >
        Download as XLSX
      </button>
      {viewTable ? (
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
              {Object.entries(rfpResults[equipmentNames[0]]?.Map || {}).map(
                ([key, questionData]) => (
                  <tr key={key}>
                    <td className="border px-4 py-2">
                      {questionData.question}
                    </td>
                    {equipmentNames.map((equipmentName) => {
                      const answer =
                        rfpResults[equipmentName]?.Map[key]?.answer;
                      return (
                        <td
                          key={equipmentName}
                          className={`border px-4 py-2 text-center ${getRowBackgroundColor(
                            answer
                          )}`}
                        >
                          {answer === "Met"
                            ? "✅"
                            : answer === "Undefined"
                            ? "❓"
                            : "❌"}
                        </td>
                      );
                    })}
                  </tr>
                )
              )}
            </tbody>
          </table>
        </div>
      ) : (
        <div>
          {Object.entries(rfpResults).map(([productName, productData]) => (
            <div
              key={productName}
              className="border rounded-lg p-4 mb-4 shadow-md"
            >
              <div className="flex justify-between items-center">
                <h2 className="text-xl font-semibold mb-2">{productName}</h2>
                <button
                  onClick={() => handleProductToggle(productName)}
                  className="bg-gray-200 py-1 px-3 rounded-md hover:bg-gray-300"
                >
                  {productToggles[productName] ? "Collapse" : "Expand"}
                </button>
              </div>
              {productToggles[productName] && (
                <div className="mt-4">
                  {Object.entries(productData.Map).map(
                    ([questionKey, questionData]) => (
                      <div
                        key={questionKey}
                        className={`p-4 mb-3 rounded-lg shadow-sm ${getRowBackgroundColor(
                          questionData.answer
                        )}`}
                      >
                        <p className="text-gray-700 font-medium">
                          <span className="font-semibold">Question:</span>{" "}
                          {questionData.question}
                        </p>
                        <p className="text-gray-700 font-medium">
                          <span className="font-semibold">Answer:</span>{" "}
                          {questionData.answer}
                        </p>
                        <p className="text-gray-700">
                          <span className="font-semibold">Source:</span>{" "}
                          {questionData.source}
                        </p>
                        <p className="text-gray-700">
                          <span className="font-semibold">Description:</span>{" "}
                          {questionData.description}
                        </p>
                      </div>
                    )
                  )}
                </div>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default RFPResult;
