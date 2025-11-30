import React from 'react';

const ViewsChart = ({ data }) => {
    const maxViews = Math.max(...data.map(d => d.views));

    return (
        <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-200">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Динамика просмотров</h3>
            <div className="h-64">
                <div className="flex items-end justify-between h-48 gap-2">
                    {data.map((day, index) => (
                        <div key={index} className="flex-1 flex flex-col items-center">
                            <div
                                className="w-full bg-blue-500 rounded-t transition-all duration-300 hover:bg-blue-600"
                                style={{ height: `${(day.views / maxViews) * 100}%` }}
                            />
                            <div className="text-xs text-gray-500 mt-2 truncate">
                                {new Date(day.date).toLocaleDateString('ru-RU', { day: 'numeric', month: 'short' })}
                            </div>
                            <div className="text-sm font-medium text-gray-900">
                                {day.views > 1000 ? `${(day.views / 1000).toFixed(1)}K` : day.views}
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default ViewsChart;