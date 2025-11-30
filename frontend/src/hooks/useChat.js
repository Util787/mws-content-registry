import { useState, useEffect } from 'react';
import { useApi } from './useApi';
import { chatAPI } from '../services/api';

export const useChat = () => {
    const [chats, setChats] = useState([]);
    const [activeChat, setActiveChat] = useState(null);
    const [messages, setMessages] = useState([]);
    const { loading, error, callApi } = useApi();

    // Загрузить историю чата
    const loadChatHistory = async (chatId) => {
        try {
            console.log('🔄 Loading chat history for:', chatId);
            const data = await callApi(chatAPI.getChatHistory, chatId);

            // Преобразуем данные от бэкенда в наш формат
            const formattedMessages = data.chatHistory?.map(msg => ({
                id: msg.id || msg.chat_id,
                chat_id: msg.chat_id,
                message: msg.message,
                is_user: msg.is_user,
                created_at: msg.created_at
            })) || [];

            setMessages(formattedMessages);
            setActiveChat(chatId);
            console.log('✅ Chat history loaded:', formattedMessages.length, 'messages');
        } catch (err) {
            console.error('❌ Failed to load chat history:', err);
            // Если чата нет, создаем новый
            if (err.response?.status === 404) {
                console.log('Chat not found, creating new one');
                createNewChat();
            }
        }
    };

    // Создать новый чат
    const createNewChat = async () => {
        try {
            console.log('🔄 Creating new chat...');
            // Если есть endpoint для создания чата, используем его
            // const data = await callApi(chatAPI.createChat);
            // const newChatId = data.chat_id;

            // Или создаем локально (временное решение)
            const newChatId = Date.now();
            const newChat = {
                id: newChatId,
                title: 'Новый чат',
                createdAt: new Date().toISOString(),
            };

            setChats(prev => [newChat, ...prev]);
            setActiveChat(newChatId);
            setMessages([]);
            console.log('✅ New chat created:', newChatId);
            return newChatId;
        } catch (err) {
            console.error('❌ Failed to create chat:', err);
            // Fallback - создаем локально
            const newChatId = Date.now();
            const newChat = {
                id: newChatId,
                title: 'Новый чат',
                createdAt: new Date().toISOString(),
            };

            setChats(prev => [newChat, ...prev]);
            setActiveChat(newChatId);
            setMessages([]);
            return newChatId;
        }
    };

    // Отправить сообщение
    const sendMessage = async (message) => {
        if (!activeChat) {
            const newChatId = await createNewChat();
            setActiveChat(newChatId);
        }

        // Создаем сообщение пользователя
        const userMessage = {
            id: Date.now(),
            chat_id: activeChat,
            message: message,
            is_user: true,
            created_at: Math.floor(Date.now() / 1000),
        };

        // Добавляем сообщение пользователя в историю
        setMessages(prev => [...prev, userMessage]);
        console.log('👤 User message sent:', message);

        try {
            // Отправляем сообщение на бэкенд
            console.log('🔄 Sending to backend, chatId:', activeChat, 'message:', message);
            const data = await callApi(chatAPI.sendMessage, activeChat, message);
            console.log('✅ Backend response:', data);

            // Создаем сообщение ассистента из ответа бэкенда
            const botMessage = {
                id: Date.now() + 1,
                chat_id: activeChat,
                message: data.answer?.message || data.message || 'Извините, не удалось получить ответ',
                is_user: false,
                created_at: data.answer?.created_at || Math.floor(Date.now() / 1000),
            };

            // Добавляем ответ ассистента в историю
            setMessages(prev => [...prev, botMessage]);
            console.log('🤖 Assistant response received:', botMessage.message);

        } catch (err) {
            console.error('❌ Failed to send message:', err);

            // Показываем сообщение об ошибке
            const errorMessage = {
                id: Date.now() + 1,
                chat_id: activeChat,
                message: '❌ Ошибка подключения к серверу. Проверьте, запущен ли бэкенд на localhost:8000',
                is_user: false,
                created_at: Math.floor(Date.now() / 1000),
            };

            setMessages(prev => [...prev, errorMessage]);
        }
    };

    // Загружаем список чатов при инициализации
    useEffect(() => {
        // Пока используем локальные чаты, можно добавить endpoint для получения списка чатов
        const savedChats = localStorage.getItem('chat_sessions');
        if (savedChats) {
            try {
                setChats(JSON.parse(savedChats));
            } catch (e) {
                console.error('Error loading saved chats:', e);
            }
        }
    }, []);

    // Сохраняем чаты в localStorage при изменении
    useEffect(() => {
        if (chats.length > 0) {
            localStorage.setItem('chat_sessions', JSON.stringify(chats));
        }
    }, [chats]);

    return {
        chats,
        activeChat,
        messages,
        loading,
        error,
        loadChatHistory,
        createNewChat,
        sendMessage,
    };
};