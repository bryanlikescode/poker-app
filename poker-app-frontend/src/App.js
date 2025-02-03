// mimport './App.css';
// src/App.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './pages/LoginPage';

function App() {
  const [message, setMessage] = useState('hello');

  useEffect(() => {
    axios.get('http://localhost:8080/')
      .then(response => {
        setMessage(response.data.message);
      })
      .catch(error => {
        console.error('There was an error fetching the data!', error);
      });
  }, []);

  return (
    <Router>
      <Routes>
        {/* Make login the default route */}
        <Route path="/" element={<LoginPage />} />
        
        {/* Add other routes here later */}
        {/* <Route path="/game" element={<GamePage />} /> */}
      </Routes>
    </Router>
  );
}

export default App;
