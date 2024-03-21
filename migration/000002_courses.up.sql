CREATE TABLE "courses"(
    "id" SERIAL PRIMARY KEY,
    "course_name" VARCHAR(100) NOT NULL,
    "description" TEXT,
    "instructor_id" INT REFERENCES users(id),
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL,
    "capacity" INT
);
