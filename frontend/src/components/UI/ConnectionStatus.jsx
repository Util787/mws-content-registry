import React from 'react';
import { Wifi, WifiOff, Server } from 'lucide-react';

const ConnectionStatus = ({ isConnected, backendUrl }) => {
    if (!isConnected) {
        return (
            <div className="fixed top-4 right-4 z-50 flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium bg-red-100 text-red-800 border border-red-200">
                <WifiOff size={16} />
                <span>Нет подключения к бэкенду</span>
                <div className="flex items-center gap-1 text-xs opacity-70">
                    <Server size={12} />
                    {backendUrl}
                </div>
            </div>
        );
    }

    return null; // Не показываем когда подключено
};

export default ConnectionStatus;