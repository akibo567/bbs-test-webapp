CREATE TABLE IF NOT EXISTS threads (
		id SERIAL NOT NULL,
		title varchar NOT NULL,
		created_at TIMESTAMP,
        updated_at TIMESTAMP,
		PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS thread_messages (
		id SERIAL NOT NULL,
		name varchar,
		message_text varchar,
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		PRIMARY KEY (id)
);