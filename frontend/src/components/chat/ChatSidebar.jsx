import React from 'react';
import { Plus, MessageSquare, Trash2, ChevronLeft } from 'lucide-react';
import Button from '../UI/Button';

const ChatSidebar = ({
                         chats,
                         activeChat,
                         onChatSelect,
                         onNewChat,
                         onClose
                     }) => {
    return (
        <div className="w-64 bg-gray-50 h-full flex flex-col border-r border-gray-200">
            {/* Header */}
            <div className="p-4 border-b border-gray-200">
                <Button
                    onClick={onNewChat}
                    variant="primary"
                    size="medium"
                    className="w-full flex items-center justify-center gap-2 bg-gray-800 hover:bg-gray-700"
                >
                    <Plus size={16} />
                    New chat
                </Button>

                <Button
                    onClick={onClose}
                    variant="ghost"
                    size="small"
                    className="w-full mt-2 md:hidden flex items-center justify-center gap-2 text-gray-600"
                >
                    <ChevronLeft size={16} />
                    Close sidebar
                </Button>
            </div>

            {/* Chat List */}
            <div className="flex-1 overflow-y-auto p-2">
                <div className="text-xs font-medium text-gray-500 uppercase tracking-wider px-3 py-2">
                    Today
                </div>

                {chats.map(chat => (
                    <div
                        key={chat.id}
                        onClick={() => onChatSelect(chat.id)}
                        className={`group px-3 py-2 rounded-lg cursor-pointer transition-colors duration-150 flex items-center gap-3 mb-1 ${
                            activeChat === chat.id
                                ? 'bg-gray-200 text-gray-900'
                                : 'hover:bg-gray-100 text-gray-700'
                        }`}
                    >
                        <MessageSquare size={16} className="flex-shrink-0" />
                        <span className="truncate text-sm flex-1">
              {chat.title || `Chat ${chat.id}`}
            </span>
                        <button
                            onClick={(e) => {
                                e.stopPropagation();
                                // TODO: Add delete functionality
                            }}
                            className="opacity-0 group-hover:opacity-100 transition-opacity p-1 hover:bg-gray-300 rounded"
                        >
                            <Trash2 size={14} className="text-gray-500" />
                        </button>
                    </div>
                ))}

                {chats.length === 0 && (
                    <div className="text-center text-gray-500 mt-8 px-4">
                        <MessageSquare size={32} className="mx-auto mb-3 opacity-50" />
                        <div className="text-sm">No chats yet</div>
                        <div className="text-xs text-gray-400 mt-1">Start a new conversation to begin</div>
                    </div>
                )}
            </div>

            {/* User section */}
            <div className="p-4 border-t border-gray-200">
                <div className="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-100 cursor-pointer">
                    <div className="w-8 h-8 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full flex items-center justify-center text-white text-sm font-medium">
                        U
                    </div>
                    <div className="flex-1 min-w-0">
                        <div className="text-sm font-medium text-gray-900 truncate">User</div>
                        <div className="text-xs text-gray-500 truncate">Free Plan</div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ChatSidebar;