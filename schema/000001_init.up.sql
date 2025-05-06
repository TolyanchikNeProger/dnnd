CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL, -- Тип элемента: 'default', 'gm_room', 'played_room', 'custom'
    title VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL, -- URL или идентификатор
    icon VARCHAR(100),
    is_customizable BOOLEAN DEFAULT TRUE -- Можно ли убрать из меню
);

CREATE TABLE user_menu_settings (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    is_visible BOOLEAN DEFAULT TRUE,
    position INT, -- Порядок в меню
    PRIMARY KEY (user_id, item_id)
);

CREATE TABLE gm_rooms (
    id SERIAL PRIMARY KEY,
    gm_id INT REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_played_rooms (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    room_id INT REFERENCES gm_rooms(id) ON DELETE CASCADE,
    last_played TIMESTAMP DEFAULT NOW()
);