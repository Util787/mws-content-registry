import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navigation from './components/Layout/Navigation';
import ChatPage from './pages/ChatPage';
import AnalyticsPage from './pages/AnalyticsPage';
import './styles/globals.css';
import VideoTable from './components/Tables/VideoTable';

function App() {
    return (
        <Router>
            <div className="app-container">
                <Navigation />

                <Routes>
                    <Route path="/" element={<ChatPage />} />
                    <Route path="/chat" element={<ChatPage />} />
                    <Route path="/analytics" element={<AnalyticsPage />} />
                    <Route path="/videos" element={<VideoTable />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;