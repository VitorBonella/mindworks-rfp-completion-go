import './App.css';
import Nav from "./components/Nav";
import React, { useEffect, useState } from 'react';
import { Routes, Route } from 'react-router-dom';
import BASE_URL from "./main";
import Rfps from "./pages/RFP";
import Equipment from "./pages/Equipment";
import NewRfp from "./pages/NewRFP";

function App() {
  const [name, setName] = useState('');

  useEffect(() => {
    (async () => {
      const response = await fetch(BASE_URL + '/api/user', {
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      });

      const content = await response.json();
      setName(content.name);
    })();
  }, []); // Ensure the effect runs only once on component mount

  return (
    <div>
      <Nav name={name} setName={setName} />
      <Routes>
        <Route path="/rfps" element={<Rfps />} />
        <Route path="/equipment" element={<Equipment />} />
        <Route path="/new_rfp" element={<NewRfp />} />
        <Route path="/*" element={<h1> PAGINA INICIAL</h1>} />
        {/* Add other routes here */}
      </Routes>
    </div>
  );
}

export default App;