import React, { useState } from 'react';
import { Plus, Youtube, Link, Loader } from 'lucide-react';
import { useApi } from '../../hooks/useApi';
import { videoAPI } from '../../services/api';

const AddVideoForm = ({ onVideoAdded }) => {
    const [isOpen, setIsOpen] = useState(false);
    const [url, setUrl] = useState('');
    const { loading, error, callApi } = useApi();

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!url.trim()) return;

        try {
            await callApi(videoAPI.addVideo, url);
            setUrl('');
            setIsOpen(false);
            if (onVideoAdded) onVideoAdded();
        } catch (err) {
            // Ошибка уже обработана в useApi
        }
    };

    const isValidYouTubeUrl = (url) => {
        const regex = /(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})/;
        return regex.test(url);
    };

    return (
        <div className="relative">
            <button
                onClick={() => setIsOpen(!isOpen)}
                className="btn-primary flex items-center gap-2"
            >
                <Plus size={18} />
                Добавить видео
            </button>

            {isOpen && (
                <>
                    <div
                        className="fixed inset-0 z-40"
                        onClick={() => setIsOpen(false)}
                    />

                    <div className="absolute top-full right-0 mt-2 w-96 bg-white rounded-xl shadow-soft border border-gray-200 p-4 z-50">
                        <div className="flex items-center gap-2 mb-3">
                            <Youtube size={20} className="text-red-600" />
                            <h3 className="font-semibold text-gray-900">Добавить YouTube видео</h3>
                        </div>

                        <form onSubmit={handleSubmit} className="space-y-3">
                            <div className="relative">
                                <input
                                    type="text"
                                    value={url}
                                    onChange={(e) => setUrl(e.target.value)}
                                    placeholder="URL YouTube видео"
                                    className="input-modern w-full pr-10"
                                />
                                <div className="absolute right-3 top-1/2 transform -translate-y-1/2">
                                    <Link size={16} className="text-gray-400" />
                                </div>
                            </div>

                            {url && !isValidYouTubeUrl(url) && (
                                <div className="text-sm text-red-600 bg-red-50 p-2 rounded-lg">
                                    ⚠️ Введите корректный URL YouTube видео
                                </div>
                            )}

                            {error && (
                                <div className="text-sm text-red-600 bg-red-50 p-2 rounded-lg">
                                    ❌ {error}
                                </div>
                            )}

                            <div className="flex gap-2 pt-2">
                                <button
                                    type="button"
                                    onClick={() => setIsOpen(false)}
                                    className="btn-secondary flex-1"
                                    disabled={loading}
                                >
                                    Отмена
                                </button>
                                <button
                                    type="submit"
                                    disabled={!url.trim() || !isValidYouTubeUrl(url) || loading}
                                    className="btn-primary flex-1 flex items-center justify-center gap-2 disabled:opacity-50"
                                >
                                    {loading ? (
                                        <>
                                            <Loader size={16} className="animate-spin" />
                                            Добавляем...
                                        </>
                                    ) : (
                                        <>
                                            <Plus size={16} />
                                            Добавить
                                        </>
                                    )}
                                </button>
                            </div>
                        </form>

                        <div className="mt-3 p-3 bg-blue-50 rounded-lg">
                            <p className="text-xs text-blue-700">
                                💡 Пример: https://www.youtube.com/watch?v=VIDEO_ID
                            </p>
                        </div>
                    </div>
                </>
            )}
        </div>
    );
};

export default AddVideoForm;