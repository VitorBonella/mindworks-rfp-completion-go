// Utility to escape CSV values
const escapeCSVValue = (value) => {
  if (value == null) return ""; // Handle null or undefined
  const stringValue = String(value); // Convert to string
  if (/[",;\n]/.test(stringValue)) {
    // If the value contains ", ; or newline, wrap in double quotes
    return `"${stringValue.replace(/"/g, '""')}"`;
  }
  return stringValue;
};

export const downloadAsCSV = (filename, data) => {
  const csvRows = [];
  const equipmentNames = Object.keys(data);

  // Add headers
  const headers = ["Requirement"];
  equipmentNames.forEach((equipmentName) => {
    headers.push(
      `${equipmentName}_Answer`,
      `${equipmentName}_Source`,
      `${equipmentName}_Description`
    );
  });
  csvRows.push(headers.map(escapeCSVValue).join(","));

  // Collect all question keys
  const questionKeys = Object.keys(data[equipmentNames[0]].Map);

  // Add rows for each question
  questionKeys.forEach((key) => {
    const row = [];
    const question = data[equipmentNames[0]].Map[key]?.question || "N/A";
    row.push(escapeCSVValue(question));

    equipmentNames.forEach((equipmentName) => {
      const productData = data[equipmentName]?.Map[key] || {};
      row.push(
        escapeCSVValue(productData.answer || ""),
        escapeCSVValue(productData.source || ""),
        escapeCSVValue(productData.description || "")
      );
    });

    csvRows.push(row.join(","));
  });

  // Create a blob from the CSV data
  const csvContent = csvRows.join("\n");
  const blob = new Blob([csvContent], { type: "text/csv" });
  const url = URL.createObjectURL(blob);

  // Create a temporary link element to trigger download
  const link = document.createElement("a");
  link.href = url;
  link.download = filename;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);

  // Clean up the object URL
  URL.revokeObjectURL(url);
};
