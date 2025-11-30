import React, { useState, useRef, useEffect } from 'react';
import { Plus, MessageCircle, User, Bot, Send, Sparkles, Paperclip, Mic } from 'lucide-react';
import { useChat } from '../../hooks/useChat';

const ChatInterface = () => {
    const [inputMessage, setInputMessage] = useState('');
    const messagesEndRef = useRef(null);

    const {
        chats,
        activeChat,
        messages,
        loading,
        error,
        loadChatHistory,
        createNewChat,
        sendMessage,
    } = useChat();

    const scrollToBottom = () => {
        messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
    };

    useEffect(() => {
        scrollToBottom();
    }, [messages, loading]);

    const handleNewChat = () => {
        createNewChat();
    };

    const handleChatSelect = (chatId) => {
        loadChatHistory(chatId);
    };

    const handleSendMessage = async (e) => {
        e.preventDefault();
        if (inputMessage.trim()) {
            await sendMessage(inputMessage.trim());
            setInputMessage('');
        }
    };

    const handleKeyPress = (e) => {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            handleSendMessage(e);
        }
    };

    return (
        <div className="flex-1 flex bg-gray-50 h-full">
            {/* Левая панель */}
            <div className="w-80 bg-white border-r border-gray-200 flex flex-col h-full shadow-soft">
                {/* Заголовок и кнопка нового чата */}
                <div className="p-6 border-b border-gray-100">
                    <div className="flex items-center gap-3 mb-4">
                        <div className="w-10 h-10 gradient-primary rounded-xl flex items-center justify-center">
                            <Sparkles size={20} className="text-white" />
                        </div>
                        <div>
                            <h1 className="font-bold text-gray-900">AI Assistant</h1>
                        </div>
                    </div>

                    <button
                        onClick={handleNewChat}
                        className="btn-primary w-full flex items-center justify-center gap-2"
                    >
                        <Plus size={20} />
                        Новый чат
                    </button>
                </div>

                {/* Список чатов */}
                <div className="flex-1 overflow-y-auto p-4 left-block">
                    <div className="space-y-2">
                        {chats.map(chat => (
                            <div
                                key={chat.id}
                                onClick={() => handleChatSelect(chat.id)}
                                className={`p-4 rounded-xl cursor-pointer transition-all duration-200 flex items-center gap-3 card-hover ${
                                    activeChat === chat.id
                                        ? 'bg-blue-50 border border-blue-200 shadow-soft'
                                        : 'bg-white border border-gray-100 hover:border-gray-200'
                                }`}
                            >
                                <div className={`p-2 rounded-lg ${
                                    activeChat === chat.id ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600'
                                }`}>
                                    <MessageCircle size={16} />
                                </div>
                                <div className="flex-1 min-w-0">
                                    <div className={`text-sm font-semibold truncate ${
                                        activeChat === chat.id ? 'text-blue-900' : 'text-gray-900'
                                    }`}>
                                        {chat.title || `Чат ${chat.id}`}
                                    </div>
                                    <div className="text-xs text-gray-500 mt-1">
                                        {new Date(chat.createdAt).toLocaleDateString('ru-RU')}
                                    </div>
                                </div>
                            </div>
                        ))}

                        {chats.length === 0 && (
                            <div className="text-center text-gray-500 py-12 item-chat">
                                <MessageCircle size={48} className="mx-auto mb-3 opacity-40" />
                                <div className="text-sm font-medium">Нет чатов</div>
                                <div className="text-xs mt-1">Создайте новый чат</div>
                            </div>
                        )}
                    </div>
                </div>
            </div>

            {/* Основная область - чат */}
            <div className="flex-1 flex flex-col h-full bg-gray-50">
                {/* Область сообщений */}
                <div className="flex-1 overflow-y-auto w-full">
                    <div className="chat-main-container">
                        {messages.length === 0 && !loading ? (
                            <div className="flex-1 flex items-center justify-center w-full h-full px-8">
                                <div className="text-center max-w-2xl">
                                    <div className="w-20 h-20 gradient-primary rounded-2xl flex items-center justify-center mx-auto mb-6 shadow-medium">
                                        <Bot size={32} className="text-white" />
                                    </div>
                                    <h2 className="text-3xl font-bold text-gray-900 mb-4">
                                        Чем могу помочь?
                                    </h2>
                                    <p className="text-gray-600 text-lg mb-8 max-w-md mx-auto">
                                        Задайте любой вопрос, и я постараюсь помочь вам
                                    </p>
                                </div>
                            </div>
                        ) : (
                            <div className="w-full px-8 py-6">
                                <div className="space-y-6">
                                    {messages.map((message) => (
                                        <div key={message.id} className={`flex gap-4 fade-in-up ${message.is_user ? 'justify-end' : 'justify-start'}`}>
                                            {!message.is_user && (
                                                <div className="w-10 h-10 gradient-secondary rounded-full flex items-center justify-center flex-shrink-0 shadow-soft">
                                                    <Bot size={18} className="text-white" />
                                                </div>
                                            )}

                                            <div className={message.is_user ? 'message-user' : 'message-assistant'}>
                                                <div className="whitespace-pre-wrap leading-relaxed">
                                                    {message.message}
                                                </div>
                                            </div>

                                            {message.is_user && (
                                                <div className="w-10 h-10 gradient-primary rounded-full flex items-center justify-center flex-shrink-0 shadow-soft">
                                                    <User size={18} className="text-white" />
                                                </div>
                                            )}
                                        </div>
                                    ))}

                                    {loading && (
                                        <div className="flex gap-4 justify-start">
                                            <div className="w-10 h-10 gradient-secondary rounded-full flex items-center justify-center flex-shrink-0 shadow-soft">
                                                <Bot size={18} className="text-white" />
                                            </div>
                                            <div className="message-assistant">
                                                <div className="flex items-center text-gray-600">
                                                    Печатает
                                                    <div className="typing-dots">
                                                        <span></span>
                                                        <span></span>
                                                        <span></span>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    )}

                                    <div ref={messagesEndRef} />
                                </div>
                            </div>
                        )}
                    </div>
                </div>

                {/* Область ввода сообщения */}
                <div className="border-t border-gray-200 bg-white p-6 w-full shadow-soft chat-input_style">
                    <div className="chat-main-container">
                        {error && (
                            <div className="mb-4 p-4 bg-red-50 border border-red-200 rounded-xl">
                                <div className="text-red-800 text-sm flex items-center gap-2">
                                    <span>⚠️</span>
                                    <span>{error}</span>
                                </div>
                            </div>
                        )}

                        <form onSubmit={handleSendMessage} className="flex gap-3 w-full items-end">
                            <div className="flex-1 relative">
                                <div className="input-with-icon">
                                    <input
                                        type="text"
                                        value={inputMessage}
                                        onChange={(e) => setInputMessage(e.target.value)}
                                        onKeyPress={handleKeyPress}
                                        placeholder="Напишите сообщение..."
                                        disabled={loading}
                                        className="input-modern animated-placeholder pr-20 input-search_style"
                                    />
                                    <div className="input-icon">
                                        <MessageCircle size={20} />
                                    </div>
                                </div>

                                {/* Кнопки действий */}
                                {/*<div className="absolute right-3 bottom-3 flex gap-1">*/}
                                {/*    <button*/}
                                {/*        type="button"*/}
                                {/*        className="p-2 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"*/}
                                {/*        title="Прикрепить файл"*/}
                                {/*    >*/}
                                {/*        <Paperclip size={18} />*/}
                                {/*    </button>*/}
                                {/*    <button*/}
                                {/*        type="button"*/}
                                {/*        className="p-2 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"*/}
                                {/*        title="Голосовой ввод"*/}
                                {/*    >*/}
                                {/*        <Mic size={18} />*/}
                                {/*    </button>*/}
                                {/*</div>*/}
                            </div>

                            <button
                                type="submit"
                                disabled={!inputMessage.trim() || loading}
                                className="btn-primary px-8 h-[56px] flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none"
                            >
                                <Send size={18} />
                                Отправить
                            </button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ChatInterface;