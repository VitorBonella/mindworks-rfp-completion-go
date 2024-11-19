import './App.css';
import Nav from "./components/Nav";
import React, { useEffect, useState } from 'react';
import { Routes, Route } from 'react-router-dom';
import BASE_URL from "./main";
import Rfps from "./pages/RFP";
import Equipment from "./pages/Equipment";
import NewRfp from "./pages/NewRFP";
import RFPResult from './pages/RFPResult';
import ApiKey from './pages/ApiKey';
function App() {
  const [name, setName] = useState('');
  const [hasApiKey, setHasApiKey] = useState(false);

  useEffect(() => {
    (async () => {
      const response = await fetch(BASE_URL + '/api/user', {
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      });

      const content = await response.json();
      setName(content.name);
      setHasApiKey(content.HasApiKey);
    })();
  }, []); // Ensure the effect runs only once on component mount

  return (
    <div>
      <Nav name={name} setName={setName} />
      <Routes>
        <Route path="/rfps" element={<Rfps />} />
        <Route path="/rfp_detail/:id" element={<RFPResult />} /> {/* New Route */}
        <Route path="/equipment" element={<Equipment />} />
        <Route path="/new_rfp" element={<NewRfp />} />
        <Route path="/*" element={<ApiKey HasApiKey={hasApiKey}/>} />
        {/* Add other routes here */}
      </Routes>
    </div>
  );
}

export default App;