import React from 'react';

const EngagementChart = ({ data }) => {
    return (
        <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-200">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">Вовлеченность</h3>
            <div className="h-64 space-y-4">
                {data.map((metric, index) => (
                    <div key={index} className="space-y-2">
                        <div className="flex justify-between text-sm">
                            <span className="text-gray-700">{metric.label}</span>
                            <span className="font-medium text-gray-900">{metric.value}%</span>
                        </div>
                        <div className="w-full bg-gray-200 rounded-full h-2">
                            <div
                                className="bg-green-500 h-2 rounded-full transition-all duration-300"
                                style={{ width: `${metric.value}%` }}
                            />
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default EngagementChart;