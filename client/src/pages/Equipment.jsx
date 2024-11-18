import React, { useState, useEffect } from "react";
import BASE_URL from "../main";

// Utility function to validate URLs
const isValidURL = (str) => {
  const pattern = new RegExp(
    "^(https?://)?[a-zA-Z0-9-]+(.[a-zA-Z0-9-]+)+(/.*)?$",
    "i"
  );
  return pattern.test(str);
};

const EquipmentForm = () => {
  const [name, setName] = useState("");
  const [downloadLink, setDownloadLink] = useState("");
  const [equipments, setEquipments] = useState([]);
  const [selectedEquipment, setSelectedEquipment] = useState(null); // Track selected equipment for showing the link

  // Fetch the list of equipment when the component mounts
  useEffect(() => {
    const fetchEquipments = async () => {
      try {
        const response = await fetch(BASE_URL + "/api/equipments", {
          credentials: "include",
        });
        if (!response.ok) {
          throw new Error("Failed to fetch equipments");
        }
        const data = await response.json();
        setEquipments(data);
      } catch (error) {
        console.error("Error fetching equipments:", error);
      }
    };

    fetchEquipments();
  }, []); // Runs only once when the component mounts

  // Handle form submission to create a new equipment
  const handleSubmit = async (e) => {
    e.preventDefault();

    //if (!isValidURL(downloadLink)) {
     // alert("Please enter a valid URL.");
     // return;
    //}

    try {
      const response = await fetch(BASE_URL + "/api/equipment", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name, download_link: downloadLink }),
      });

      if (!response.ok) {
        const content = await response.json();
        alert(content.message)
        throw new Error("Failed to create equipment");
      }

      // Update the equipment list by fetching the latest data
      const newEquipment = await response.json();
      setEquipments((prevEquipments) => [...prevEquipments, newEquipment]);
      
      setName("");
      setDownloadLink("");
    } catch (error) {
      console.error("Error creating equipment:", error);
    }
  };

  // Handle equipment deletion
  const handleDelete = async (id) => {
    try {
      const response = await fetch(BASE_URL + "/api/equipment", {
        method: "DELETE",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ id }),
      });

      if (!response.ok) {
        throw new Error("Failed to delete equipment");
      }

      // Update the equipment list by removing the deleted equipment
      setEquipments((prevEquipments) =>
        prevEquipments.filter((equipment) => equipment.id !== id)
      );
    } catch (error) {
      console.error("Error deleting equipment:", error);
    }
  };

  const toggleLinkVisibility = (index) => {
    setSelectedEquipment(selectedEquipment === index ? null : index); // Toggle the visibility of the link
  };

  return (
    <div className="max-w-4xl mx-auto p-6">
      <h2 className="text-2xl font-bold mb-6">Equipment Registration</h2>

      <form onSubmit={handleSubmit} className="mb-6">
        <div className="mb-4">
          <label htmlFor="name" className="block text-gray-700">
            Equipment Name
          </label>
          <input
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="mt-2 p-2 w-full border border-gray-300 rounded-md"
            placeholder="Enter equipment name"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="downloadLink" className="block text-gray-700">
            Equipment Download Link
          </label>
          <input
            id="downloadLink"
            type="text"
            value={downloadLink}
            onChange={(e) => setDownloadLink(e.target.value)}
            className="mt-2 p-2 w-full border border-gray-300 rounded-md"
            placeholder="Enter equipment download link"
            required
          />
        </div>

        <button
          type="submit"
          className="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600"
        >
          Register Equipment
        </button>
      </form>

      <h3 className="text-xl font-semibold mb-4">Equipment List</h3>

      <ul className="space-y-4">
        {equipments.map((equipment, index) => (
          <li
            key={equipment.id}
            className="flex items-center justify-between p-4 border border-gray-200 rounded-md"
          >
            <div className="flex flex-col">
              <span className="text-lg font-medium">{equipment.name}</span>
              <button
                onClick={() => toggleLinkVisibility(index)}
                className="text-blue-500 hover:text-blue-700 mt-2"
              >
                View Link
              </button>

              {selectedEquipment === index && (
                <div className="mt-2 text-gray-500 text-sm">
                  <p className="italic">Link:</p>
                  <p className="">{equipment.download_link}</p>
                </div>
              )}
            </div>

            <button
              onClick={() => handleDelete(equipment.id)}
              className="text-red-500 hover:text-red-700"
            >
              Exclude
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default EquipmentForm;
