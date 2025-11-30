import React from 'react';
import AnalyticsDashboard from '../components/Analytics/AnalyticsDashboard';

const AnalyticsPage = () => {
    return (
        <div className="flex-1 bg-gray-50 min-h-screen w-full">
            <div className="analytics-container px-8 py-8">
                <div className="mb-8">
                    <h1 className="text-3xl font-bold text-gray-900">Аналитика видео</h1>
                    <p className="text-gray-600 mt-2">
                        Статистика и метрики по всем видео в системе
                    </p>
                </div>

                <AnalyticsDashboard />
            </div>
        </div>
    );
};

export default AnalyticsPage;