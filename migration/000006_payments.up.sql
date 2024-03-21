CREATE TABLE Payments (
    "id" SERIAL PRIMARY KEY,
    "student_id" INT REFERENCES users(id),
    "amount" DECIMAL(10, 2) NOT NULL,
    "payment_date" DATE DEFAULT CURRENT_DATE,
    "payment_method" VARCHAR(50)
);
