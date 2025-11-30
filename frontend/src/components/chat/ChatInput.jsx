import React, { useState, useRef } from 'react';
import { Send, Paperclip, Mic } from 'lucide-react';
import Button from '../UI/Button';

const ChatInput = ({ onSendMessage, disabled = false }) => {
    const [message, setMessage] = useState('');
    const textareaRef = useRef(null);

    const handleSubmit = (e) => {
        e.preventDefault();
        if (message.trim() && !disabled) {
            onSendMessage(message.trim());
            setMessage('');
            if (textareaRef.current) {
                textareaRef.current.style.height = 'auto';
            }
        }
    };

    const handleKeyPress = (e) => {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            handleSubmit(e);
        }
    };

    const handleInputChange = (e) => {
        setMessage(e.target.value);

        // Auto-resize textarea
        if (textareaRef.current) {
            textareaRef.current.style.height = 'auto';
            textareaRef.current.style.height = Math.min(textareaRef.current.scrollHeight, 120) + 'px';
        }
    };

    return (
        <div className="border-t border-gray-200 bg-white p-4">
            <div className="max-w-3xl mx-auto">
                <form onSubmit={handleSubmit} className="relative">
                    <div className="flex items-end gap-2">
                        {/* Attachment button */}
                        <button
                            type="button"
                            disabled={disabled}
                            className="p-2 text-gray-500 hover:bg-gray-100 rounded-lg disabled:opacity-50 transition-colors"
                        >
                            <Paperclip size={20} />
                        </button>

                        {/* Textarea */}
                        <div className="flex-1 relative">
              <textarea
                  ref={textareaRef}
                  value={message}
                  onChange={handleInputChange}
                  onKeyPress={handleKeyPress}
                  placeholder="Message Assistant..."
                  disabled={disabled}
                  rows="1"
                  className="w-full px-4 py-3 pr-12 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed text-gray-900 placeholder-gray-500 resize-none max-h-32"
                  style={{ height: 'auto' }}
              />

                            {/* Voice input button */}
                            <button
                                type="button"
                                disabled={disabled}
                                className="absolute right-3 bottom-3 p-1 text-gray-500 hover:bg-gray-100 rounded disabled:opacity-50 transition-colors"
                            >
                                <Mic size={16} />
                            </button>
                        </div>

                        {/* Send button */}
                        <button
                            type="submit"
                            disabled={!message.trim() || disabled}
                            className="p-3 bg-black text-white rounded-lg hover:bg-gray-800 disabled:opacity-50 disabled:cursor-not-allowed transition-colors flex items-center justify-center"
                        >
                            <Send size={16} />
                        </button>
                    </div>
                </form>

                <div className="text-xs text-center text-gray-500 mt-3">
                    Assistant can make mistakes. Consider checking important information.
                </div>
            </div>
        </div>
    );
};

export default ChatInput;