import React from 'react';
import { Eye, ThumbsUp, MessageCircle, TrendingUp } from 'lucide-react';

const MetricsGrid = ({ stats }) => {
    const metrics = [
        {
            label: 'Всего видео',
            value: stats.totalVideos,
            icon: <Eye className="w-6 h-6" />,
            color: 'blue'
        },
        {
            label: 'Всего просмотров',
            value: stats.totalViews.toLocaleString(),
            icon: <TrendingUp className="w-6 h-6" />,
            color: 'green'
        },
        {
            label: 'Средние лайки',
            value: Math.round(stats.avgLikes),
            icon: <ThumbsUp className="w-6 h-6" />,
            color: 'purple'
        },
        {
            label: 'Средние комментарии',
            value: Math.round(stats.avgComments),
            icon: <MessageCircle className="w-6 h-6" />,
            color: 'orange'
        }
    ];

    const colorClasses = {
        blue: 'bg-blue-500',
        green: 'bg-green-500',
        purple: 'bg-purple-500',
        orange: 'bg-orange-500'
    };

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            {metrics.map((metric, index) => (
                <div key={index} className="bg-white rounded-xl p-6 shadow-sm border border-gray-200">
                    <div className="flex items-center justify-between">
                        <div>
                            <p className="text-gray-600 text-sm font-medium">{metric.label}</p>
                            <p className="text-2xl font-bold text-gray-900 mt-1">{metric.value}</p>
                        </div>
                        <div className={`p-3 rounded-lg ${colorClasses[metric.color]} text-white`}>
                            {metric.icon}
                        </div>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default MetricsGrid;