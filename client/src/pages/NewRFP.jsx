import { useEffect, useState } from "react";
import BASE_URL from "../main";

const NewRFP = () => {
  const [equipments, setEquipments] = useState([]);
  const [selectedEquipment, setSelectedEquipment] = useState([]);
  const [requirements, setRequirements] = useState("");
  const [name, setName] = useState("");

  useEffect(() => {
    // Fetch equipment list with credentials included
    fetch(`${BASE_URL}/api/equipments`, {
      method: "GET",
      credentials: "include", // This ensures cookies and other credentials are sent
    })
      .then((response) => response.json())
      .then((data) => {
        // Check if the data is in the expected format
        setEquipments(data || []);  // Ensure it's an array
      })
      .catch((error) => console.error("Error fetching equipments:", error));
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();

    // Prepare the selected equipment objects in the correct format
    const selectedEquipments = equipments.filter((equipment) =>
      selectedEquipment.includes(equipment.id)
    ).map((equipment) => ({
      id: equipment.id,
      name: equipment.name,
      download_link: equipment.download_link
    }));

    const rfpData = {
      name,
      requirements: requirements.split("\n").map((item) => item.trim()).filter(Boolean),
      equipments: selectedEquipments,  // Use the selected equipment objects
    };

    // POST request to submit the RFP with credentials
    fetch(`${BASE_URL}/api/rfp`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include", // Ensure credentials are sent with the request
      body: JSON.stringify(rfpData),
    })
      .then((response) => {
        if (!response.ok) {
          // If the response is not "ok" (i.e., status is not 200), throw an error
          throw new Error(`Error: ${response.statusText}`);
        }
        return response.json(); // Only parse JSON if response is OK
      })
      .then((data) => {
        if (data.message && data.message !== "") {
          throw new Error(`Error: ${data.message}`);
        }
        alert("RFP created successfully!");
      })
      .catch((error) => {
        // Handle errors, including non-200 responses
        console.error("Error submitting RFP:", error);
        alert(`An error occurred while submitting the RFP: ${error.message}`);
      });
  };

  return (
    <div className="max-w-2xl mx-auto p-6 bg-white rounded-lg shadow-lg">
      <h1 className="text-2xl font-semibold mb-4">New RFP</h1>
      <form onSubmit={handleSubmit}>
        {/* RFP Name */}
        <div className="mb-4">
          <label htmlFor="name" className="block text-sm font-medium text-gray-700">
            RFP Name
          </label>
          <input
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>

        {/* Requirements */}
        <div className="mb-4">
          <label htmlFor="requirements" className="block text-sm font-medium text-gray-700">
            Requirements (one per line)
          </label>
          <textarea
            id="requirements"
            value={requirements}
            onChange={(e) => setRequirements(e.target.value)}
            rows="4"
            className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>

        {/* Equipment Selection */}
        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700">Select Equipment</label>
          <div className="mt-2 space-y-2">
            {equipments.length > 0 ? (
              equipments.map((equipment) => (
                <div key={equipment.id} className="flex items-center">
                  <input
                    type="checkbox"
                    id={`equipment-${equipment.id}`}
                    value={equipment.id}
                    onChange={(e) => {
                      const checked = e.target.checked;
                      setSelectedEquipment((prev) =>
                        checked
                          ? [...prev, equipment.id]
                          : prev.filter((id) => id !== equipment.id)
                      );
                    }}
                    className="mr-2"
                  />
                  <label htmlFor={`equipment-${equipment.id}`} className="text-sm text-gray-600">
                    {equipment.name}
                    <br />
                    <a href={equipment.download_link} className="text-blue-600 underline">
                      Download PDF
                    </a>
                  </label>
                </div>
              ))
            ) : (
              <p className="text-sm text-gray-500">No equipment available</p>
            )}
          </div>
        </div>

        {/* Requirements Preview Table */}
        <div className="mt-8 mb-4">
          <h2 className="text-lg font-semibold mb-2">Requirements Preview</h2>
          <div className="overflow-y-auto max-h-40"> {/* Set a max height and make the table scrollable */}
            <table className="min-w-full table-auto border-collapse border border-gray-300">
              <thead>
                <tr>
                  <th className="px-4 py-2 border text-left">#</th>
                  <th className="px-4 py-2 border text-left">Requirement</th>
                </tr>
              </thead>
              <tbody>
                {requirements.split("\n").map((req, index) => {
                  if (req.trim() === "") return null; // Skip empty lines
                  return (
                    <tr key={index} className="transition-all duration-300 ease-in-out transform hover:bg-gray-100">
                      <td className="px-4 py-2 border">{index + 1}</td>
                      <td className="px-4 py-2 border">{req.trim()}</td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        </div>

        {/* Submit Button */}
        <div className="flex justify-end">
          <button
            type="submit"
            className="px-6 py-2 text-white bg-indigo-600 rounded-lg shadow hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          >
            Create RFP
          </button>
        </div>
      </form>
    </div>
  );
};

export default NewRFP;
