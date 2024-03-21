CREATE TABLE "events" (
    "id" SERIAL PRIMARY KEY,
    "event_name" VARCHAR(100) NOT NULL,
    "description" TEXT,
    "start_date" TIMESTAMP NOT NULL,
    "end_date" TIMESTAMP NOT NULL,
    "location" VARCHAR(100) NOT NULL
);