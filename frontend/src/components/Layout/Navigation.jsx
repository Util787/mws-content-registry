import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { MessageCircle, BarChart3 } from 'lucide-react';

const Navigation = () => {
    const location = useLocation();

    return (
        <div className="w-20 bg-gray-900 text-white flex flex-col items-center py-6">
            <Link
                to="/chat"
                className={`p-3 rounded-lg mb-4 transition-colors ${
                    location.pathname === '/chat' ? 'bg-blue-600' : 'hover:bg-gray-800'
                }`}
            >
                <MessageCircle size={24} />
            </Link>

            <Link
                to="/analytics"
                className={`p-3 rounded-lg transition-colors ${
                    location.pathname === '/analytics' ? 'bg-blue-600' : 'hover:bg-gray-800'
                }`}
            >
                <BarChart3 size={24} />
            </Link>
        </div>
    );
};

export default Navigation;