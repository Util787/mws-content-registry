import React, { useState, useEffect } from 'react';
import { Filter, RefreshCw, Youtube } from 'lucide-react';
import AddVideoForm from '../Video/AddVideoForm';
import { useApi } from '../../hooks/useApi';
import { recordsAPI } from '../../services/api';

const VideoTable = () => {
    const [records, setRecords] = useState([]);
    const [filters, setFilters] = useState({
        pageNum: 1,
        pageSize: 10,
        sort: []
    });
    const { loading, error, callApi } = useApi();

    const loadRecords = async () => {
        try {
            const data = await callApi(recordsAPI.getRecords, filters);
            setRecords(data.records || []);
        } catch (err) {
            console.error('Failed to load records:', err);
        }
    };

    useEffect(() => {
        loadRecords();
    }, [filters]);

    const handleVideoAdded = () => {
        loadRecords();
    };

    return (
        <div className="p-6 bg-gray-50 min-h-screen">
            <div className="max-w-7xl mx-auto">
                <div className="flex justify-between items-center mb-6">
                    <div>
                        <h1 className="text-2xl font-bold text-gray-900">YouTube Videos</h1>
                        <p className="text-gray-600">Управление видео контентом</p>
                    </div>

                    <div className="flex items-center gap-3">
                        <button className="btn-secondary flex items-center gap-2">
                            <Filter size={16} />
                            Фильтры
                        </button>

                        <button
                            onClick={loadRecords}
                            className="btn-secondary flex items-center gap-2"
                            disabled={loading}
                        >
                            <RefreshCw size={16} className={loading ? 'animate-spin' : ''} />
                            Обновить
                        </button>

                        <AddVideoForm onVideoAdded={handleVideoAdded} />
                    </div>
                </div>

                <div className="card">
                    <div className="overflow-x-auto">
                        <table className="w-full">
                            <thead>
                            <tr className="border-b border-gray-200">
                                <th className="text-left p-4 font-semibold text-gray-900">ID</th>
                                <th className="text-left p-4 font-semibold text-gray-900">Видео</th>
                                <th className="text-left p-4 font-semibold text-gray-900">Просмотры</th>
                                <th className="text-left p-4 font-semibold text-gray-900">Лайки</th>
                                <th className="text-left p-4 font-semibold text-gray-900">Дата</th>
                                <th className="text-left p-4 font-semibold text-gray-900">Действия</th>
                            </tr>
                            </thead>
                            <tbody>
                            {records.map(record => (
                                <TableRow key={record.recordId} record={record} />
                            ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    );
};

const TableRow = ({ record }) => {
    return (
        <tr className="border-b border-gray-100 hover:bg-gray-50">
            <td className="p-4 text-gray-900">{record.recordId}</td>
            <td className="p-4">
                <div className="flex items-center gap-3">
                    <div className="w-16 h-12 bg-gray-200 rounded flex items-center justify-center">
                        <Youtube size={20} className="text-red-600" />
                    </div>
                    <div>
                        <div className="font-medium text-gray-900">
                            {record.fields?.title || 'Без названия'}
                        </div>
                        <div className="text-sm text-gray-500">
                            {record.fields?.author}
                        </div>
                    </div>
                </div>
            </td>
            <td className="p-4 text-gray-700">{record.fields?.views?.toLocaleString()}</td>
            <td className="p-4 text-gray-700">{record.fields?.likes?.toLocaleString()}</td>
            <td className="p-4 text-gray-700">
                {record.fields?.published_at ?
                    new Date(record.fields.published_at * 1000).toLocaleDateString('ru-RU') :
                    '-'
                }
            </td>
            <td className="p-4">
                <button className="text-blue-600 hover:text-blue-800 text-sm font-medium">
                    Анализировать
                </button>
            </td>
        </tr>
    );
};

export default VideoTable;