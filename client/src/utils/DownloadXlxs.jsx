import * as XLSX from "xlsx";

export const downloadAsExcel = (filename, data) => {
  const equipmentNames = Object.keys(data);

  // Create a workbook and a worksheet
  const workbook = XLSX.utils.book_new();
  const worksheetData = [];

  // Add headers
  const headers = ["Requirement"];
  equipmentNames.forEach((equipmentName) => {
    headers.push(
      `${equipmentName}_Answer`,
      `${equipmentName}_Source`,
      `${equipmentName}_Description`
    );
  });
  worksheetData.push(headers);

  // Collect all question keys
  const questionKeys = Object.keys(data[equipmentNames[0]].Map);

  // Add rows for each question
  questionKeys.forEach((key) => {
    const row = [];
    const question = data[equipmentNames[0]].Map[key]?.question || "N/A";
    row.push(question);

    equipmentNames.forEach((equipmentName) => {
      const productData = data[equipmentName]?.Map[key] || {};
      let answer = productData.answer || "N/A";

      // Add emoji directly to the answer based on its value
      if (answer === "Met") {
        answer = "✅";  // Replace with a green check
      } else if (answer === "Undefined") {
        answer = "❔";  // Replace with a question mark emoji
      } else if (answer === "Not Met") {
        answer = "❌";  // Replace with a red cross
      }

      row.push(answer, productData.source || "N/A", productData.description || "N/A");
    });

    worksheetData.push(row);
  });

  // Create the worksheet
  const worksheet = XLSX.utils.aoa_to_sheet(worksheetData);

  // Add the worksheet to the workbook
  XLSX.utils.book_append_sheet(workbook, worksheet, "RFP Results");

  // Generate the Excel file and trigger download
  XLSX.writeFile(workbook, filename);
};
