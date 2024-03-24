CREATE TABLE "roles"(
    "role_id" SERIAL PRIMARY KEY,
    "role_name" VARCHAR(20) NOT NULL
);

CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "role_id" INT REFERENCES roles(role_id),
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "email" VARCHAR(100),
    "phone_number" VARCHAR(20),
    "address" TEXT,
    "registration_date" DATE DEFAULT CURRENT_DATE,
    "updated_at" DATE DEFAULT CURRENT_DATE
);
