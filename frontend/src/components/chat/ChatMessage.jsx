import React from 'react';
import { User, Bot } from 'lucide-react';

const ChatMessage = ({ message, isUser, isTyping = false }) => {
    return (
        <div className={`group w-full py-4 ${isUser ? 'bg-white' : 'bg-gray-50'} border-b border-gray-100`}>
            <div className="max-w-3xl mx-auto px-4">
                <div className="flex gap-4">
                    {/* Avatar */}
                    <div className={`flex-shrink-0 w-8 h-8 rounded-sm flex items-center justify-center ${
                        isUser ? 'bg-green-500' : 'bg-gray-300'
                    }`}>
                        {isUser ? (
                            <User size={16} className="text-white" />
                        ) : (
                            <Bot size={16} className="text-white" />
                        )}
                    </div>

                    {/* Message Content */}
                    <div className="flex-1 min-w-0">
                        <div className="font-semibold text-gray-900 mb-1">
                            {isUser ? 'You' : 'Assistant'}
                        </div>

                        <div className="text-gray-800 whitespace-pre-wrap leading-relaxed">
                            {isTyping ? (
                                <div className="flex items-center text-gray-600">
                                    <span>Thinking</span>
                                    <div className="typing-dots">
                                        <span></span>
                                        <span></span>
                                        <span></span>
                                    </div>
                                </div>
                            ) : (
                                message
                            )}
                        </div>

                        {/* Message actions */}
                        <div className="flex items-center gap-2 mt-2 opacity-0 group-hover:opacity-100 transition-opacity">
                            <button className="p-1 hover:bg-gray-200 rounded text-gray-500">
                                <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
                                    <path d="M8 5H6C4.89543 5 4 5.89543 4 7V19C4 20.1046 4.89543 21 6 21H18C19.1046 21 20 20.1046 20 19V7C20 5.89543 19.1046 5 18 5H16M8 5C8 6.10457 8.89543 7 10 7H14C15.1046 7 16 6.10457 16 5M8 5C8 3.89543 8.89543 3 10 3H14C15.1046 3 16 3.89543 16 5" stroke="currentColor" strokeWidth="2"/>
                                </svg>
                            </button>
                            <button className="p-1 hover:bg-gray-200 rounded text-gray-500">
                                <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
                                    <path d="M8 16H6C4.89543 16 4 16.8954 4 18V19C4 20.1046 4.89543 21 6 21H18C19.1046 21 20 20.1046 20 19V18C20 16.8954 19.1046 16 18 16H16M8 16L12 12M12 12L16 16M12 12V21" stroke="currentColor" strokeWidth="2"/>
                                </svg>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ChatMessage;