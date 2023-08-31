-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS chatroom WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create messages table
CREATE TABLE IF NOT EXISTS chatroom.messages (
message_id bigint PRIMARY KEY,
message text,
user text,
room_id int
);

-- Create Metadata table
CREATE TABLE IF NOT EXISTS chatroom.metadata (
room_id int PRIMARY KEY,
room_name text,
);
