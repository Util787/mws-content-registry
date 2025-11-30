CREATE TABLE chats (
    id SERIAL PRIMARY KEY,    
    chat_id INTEGER, 
    is_user BOOLEAN,
    message TEXT,
    created_at INTEGER
);

CREATE INDEX idx_chat_id ON chats(chat_id);