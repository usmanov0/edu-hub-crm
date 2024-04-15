CREATE TABLE "courses"(
    "id" SERIAL PRIMARY KEY,
    "course_name" VARCHAR(100) NOT NULL,
    "description" TEXT,
    "user_id" INT REFERENCES users(id),
    "capacity" INT,
    "updated_at" DATE NOT NULL,
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL
);
