CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "role"  VARCHAR(30) NOT NULL,
    "email" VARCHAR(50),
    "password" VARCHAR(50),
    "phone_number" VARCHAR(20),
    "address" TEXT,
    "registration_date" DATE DEFAULT CURRENT_DATE,
    "updated_at" DATE DEFAULT CURRENT_DATE
);
