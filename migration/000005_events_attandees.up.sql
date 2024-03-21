CREATE TABLE "event_attendees" (
    "id" INT REFERENCES events(id),
    "student_id" INT REFERENCES users(id),
    PRIMARY KEY (id, student_id)
);
