CREATE TABLE "enrollments" (
    "id" SERIAL PRIMARY KEY,
    "student_id" INT REFERENCES users(id),
    "course_id" INT REFERENCES courses(id),
    "enrollment_date" DATE DEFAULT CURRENT_DATE,
    "completion_status" VARCHAR(20) DEFAULT 'Enrolled',
    CONSTRAINT valid_completion_status CHECK (completion_status IN ('Completed', 'Enrolled', 'In Progress', 'Withdrawn', 'Failed', 'On Hold', 'Cancelled'))
);
