CREATE TABLE chats (
    id SERIAL PRIMARY KEY,    
    chat_num INTEGER, 
    is_user BOOLEAN,
    message TEXT
);

CREATE INDEX idx_chat_num ON chats(chat_num);