CREATE TABLE IF NOT EXISTS users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name varchar NOT NULL,
	age int32 NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT(now())
);