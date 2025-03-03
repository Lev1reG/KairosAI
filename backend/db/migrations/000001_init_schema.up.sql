-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT, -- Nullable for OAuth users
    oauth_provider TEXT DEFAULT 'local', -- Values: 'github', 'google', 'local'
    oauth_id TEXT UNIQUE, -- Stores OAuth user ID if applicable
    avatar_url TEXT, -- Stores profile picture from OAuth
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create schedules table
CREATE TABLE schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'scheduled', -- Enum: 'scheduled', 'canceled', 'completed'
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create schedule_participants table
CREATE TABLE schedule_participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Nullable for external participants
    schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL,
    response_status VARCHAR(20) NOT NULL DEFAULT 'pending', -- Options: 'accepted', 'declined', 'pending'
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create notifications table
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    sent_at TIMESTAMP DEFAULT NOW()
);

-- Create chat_logs table
CREATE TABLE chat_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    response TEXT,
    timestamp TIMESTAMP DEFAULT NOW()
);

