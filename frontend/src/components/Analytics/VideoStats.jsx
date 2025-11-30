import React, { useState } from 'react';
import { ExternalLink, Eye, ThumbsUp, MessageCircle } from 'lucide-react';

const VideoStats = ({ videos }) => {
    const [sortBy, setSortBy] = useState('views');

    const sortedVideos = [...videos].sort((a, b) => b[sortBy] - a[sortBy]);

    const formatNumber = (num) => {
        if (num >= 1000000) {
            return (num / 1000000).toFixed(1) + 'M';
        }
        if (num >= 1000) {
            return (num / 1000).toFixed(1) + 'K';
        }
        return num;
    };

    return (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200">
            <div className="p-6 border-b border-gray-200">
                <h2 className="text-xl font-semibold text-gray-900">Статистика по видео</h2>
                <p className="text-gray-600 mt-1">Детальная аналитика каждого видео</p>
            </div>

            {/* Сортировка */}
            <div className="p-4 border-b border-gray-200">
                <div className="flex gap-2">
                    {[
                        { key: 'views', label: 'Просмотры' },
                        { key: 'likes', label: 'Лайки' },
                        { key: 'comments_count', label: 'Комментарии' },
                        { key: 'published_at', label: 'Дата' }
                    ].map(option => (
                        <button
                            key={option.key}
                            onClick={() => setSortBy(option.key)}
                            className={`px-3 py-1 rounded-lg text-sm font-medium transition-colors ${
                                sortBy === option.key
                                    ? 'bg-blue-600 text-white'
                                    : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
                            }`}
                        >
                            {option.label}
                        </button>
                    ))}
                </div>
            </div>

            {/* Таблица видео */}
            <div className="overflow-x-auto">
                <table className="w-full">
                    <thead>
                    <tr className="border-b border-gray-200">
                        <th className="text-left p-4 font-medium text-gray-900">Видео</th>
                        <th className="text-left p-4 font-medium text-gray-900">
                            <div className="flex items-center gap-1">
                                <Eye size={16} />
                                Просмотры
                            </div>
                        </th>
                        <th className="text-left p-4 font-medium text-gray-900">
                            <div className="flex items-center gap-1">
                                <ThumbsUp size={16} />
                                Лайки
                            </div>
                        </th>
                        <th className="text-left p-4 font-medium text-gray-900">
                            <div className="flex items-center gap-1">
                                <MessageCircle size={16} />
                                Комментарии
                            </div>
                        </th>
                        <th className="text-left p-4 font-medium text-gray-900">Дата</th>
                    </tr>
                    </thead>
                    <tbody>
                    {sortedVideos.map((video, index) => (
                        <tr key={video.id} className="border-b border-gray-100 hover:bg-gray-50">
                            <td className="p-4">
                                <div className="flex items-center gap-3">
                                    <div className="w-12 h-9 bg-gray-200 rounded flex items-center justify-center">
                                        <img
                                            src={`https://img.youtube.com/vi/${video.video_id}/default.jpg`}
                                            alt=""
                                            className="w-12 h-9 rounded object-cover"
                                        />
                                    </div>
                                    <div className="flex-1 min-w-0">
                                        <div className="font-medium text-gray-900 truncate max-w-xs">
                                            {video.title || `Video ${video.video_id}`}
                                        </div>
                                        <div className="text-sm text-gray-500 truncate max-w-xs">
                                            {video.author}
                                        </div>
                                    </div>
                                    <a
                                        href={`https://youtube.com/watch?v=${video.video_id}`}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="p-1 hover:bg-gray-200 rounded transition-colors"
                                    >
                                        <ExternalLink size={16} className="text-gray-400" />
                                    </a>
                                </div>
                            </td>
                            <td className="p-4 font-medium text-gray-900">
                                {formatNumber(video.views)}
                            </td>
                            <td className="p-4 text-gray-700">
                                {formatNumber(video.likes)}
                            </td>
                            <td className="p-4 text-gray-700">
                                {formatNumber(video.comments_count)}
                            </td>
                            <td className="p-4 text-gray-700">
                                {new Date(video.published_at * 1000).toLocaleDateString('ru-RU')}
                            </td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default VideoStats;