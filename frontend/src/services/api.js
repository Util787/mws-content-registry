import axios from 'axios';

const API_BASE_URL = 'http://localhost:8000/api/v1';

const api = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

// Добавляем обработку CORS ошибок
api.interceptors.response.use(
    (response) => {
        console.log('✅ API Response:', response.status, response.config.url);
        return response;
    },
    (error) => {
        if (error.code === 'ERR_NETWORK') {
            console.error('❌ CORS/Network Error:', 'Проверьте настройки CORS на бэкенде');
        }
        return Promise.reject(error);
    }
);

// Chat API - используем только существующие endpoint'ы из документации
export const chatAPI = {
    // Отправить сообщение в чат
    sendMessage: (chatId, message) =>
        api.post('/ai-chat/send-message', {
            chat_id: parseInt(chatId),
            message: message
        }),

    // Получить историю чата
    getChatHistory: (chatId) =>
        api.get(`/ai-chat/${chatId}`),
};

// Video API - используем существующие endpoint'ы
export const videoAPI = {
    addVideo: (url) =>
        api.post('/add-yt-video', { url }),

    addRecentVideos: () =>
        api.post('/add-yt-videos/recent'),

    analyzeContent: (recordId) =>
        api.post(`/add-llm-analyze/${recordId}`),
};

// Records API - используем существующий endpoint
export const recordsAPI = {
    getRecords: (params) =>
        api.get('/records', { params }),
};

// Analytics API - временно отключаем, пока нет endpoint'а
export const analyticsAPI = {
    getVideoStats: (timeRange = '7d') =>
        api.get(`/analytics/videos?range=${timeRange}`),
};

// ВРЕМЕННО: Mock для аналитики, пока бэкенд не готов
analyticsAPI.getVideoStats = async () => {
    await new Promise(resolve => setTimeout(resolve, 1000));
    return {
        data: {
            overview: {
                totalVideos: 24,
                totalViews: 1542078,
                avgLikes: 12543,
                avgComments: 892
            },
            viewsData: [
                { date: '2024-01-01', views: 12000 },
                { date: '2024-01-02', views: 18000 },
                { date: '2024-01-03', views: 15000 },
                { date: '2024-01-04', views: 22000 },
                { date: '2024-01-05', views: 19000 },
                { date: '2024-01-06', views: 25000 },
                { date: '2024-01-07', views: 21000 }
            ],
            engagementData: [
                { label: 'Лайки/Просмотры', value: 4.2 },
                { label: 'Комментарии/Просмотры', value: 0.8 },
                { label: 'Удержание', value: 68 },
                { label: 'CTR', value: 12 }
            ],
            videos: [
                {
                    id: 1,
                    video_id: 'abc123',
                    title: 'Как работает искусственный интеллект',
                    author: 'Tech Channel',
                    views: 250000,
                    likes: 12000,
                    comments_count: 850,
                    published_at: 1704067200
                }
            ]
        }
    };
};

export default api;