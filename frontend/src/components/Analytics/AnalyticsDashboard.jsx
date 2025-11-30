import React, { useState, useEffect } from 'react';
import { useApi } from '../../hooks/useApi';
import { analyticsAPI } from '../../services/api';
import MetricsGrid from './MetricsGrid';
import VideoStats from './VideoStats';
import ViewsChart from './Charts/ViewsChart';
import EngagementChart from './Charts/EngagementChart';

const AnalyticsDashboard = () => {
    const [stats, setStats] = useState(null);
    const [timeRange, setTimeRange] = useState('7d');
    const { loading, error, callApi } = useApi();

    const loadAnalytics = async () => {
        try {
            // Используем mock данные, пока бэкенд не готов
            const data = await callApi(analyticsAPI.getVideoStats, timeRange);
            setStats(data);
        } catch (err) {
            console.error('Failed to load analytics:', err);
            // Если ошибка, используем fallback данные
            setStats({
                overview: {
                    totalVideos: 0,
                    totalViews: 0,
                    avgLikes: 0,
                    avgComments: 0
                },
                viewsData: [],
                engagementData: [],
                videos: []
            });
        }
    };

    useEffect(() => {
        loadAnalytics();
    }, [timeRange]);

    if (loading && !stats) {
        return (
            <div className="flex justify-center items-center h-64">
                <div className="text-gray-600">Загрузка аналитики...</div>
            </div>
        );
    }

    return (
        <div className="space-y-6">
            {/* Фильтры и контролы */}
            <div className="flex justify-between items-center">
                <div className="flex gap-2">
                    {['1d', '7d', '30d', '90d'].map(range => (
                        <button
                            key={range}
                            onClick={() => setTimeRange(range)}
                            className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                                timeRange === range
                                    ? 'bg-blue-600 text-white'
                                    : 'bg-white text-gray-700 hover:bg-gray-100 border border-gray-300'
                            }`}
                        >
                            {range === '1d' ? '1 день' :
                                range === '7d' ? '7 дней' :
                                    range === '30d' ? '30 дней' : '90 дней'}
                        </button>
                    ))}
                </div>

                <button
                    onClick={loadAnalytics}
                    className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
                >
                    Обновить данные
                </button>
            </div>

            {error && (
                <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
                    <div className="text-yellow-800">
                        ⚠️ Используются демо-данные. Бэкенд временно недоступен.
                    </div>
                </div>
            )}

            {/* Основные метрики */}
            {stats && <MetricsGrid stats={stats.overview} />}

            {/* Графики */}
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
                {stats && stats.viewsData.length > 0 && <ViewsChart data={stats.viewsData} />}
                {stats && stats.engagementData.length > 0 && <EngagementChart data={stats.engagementData} />}
            </div>

            {/* Детальная статистика по видео */}
            {stats && <VideoStats videos={stats.videos} />}
        </div>
    );
};

export default AnalyticsDashboard;