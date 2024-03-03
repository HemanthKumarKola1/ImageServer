CREATE TABLE "schedules" (
  "schedule_time" SERIAL PRIMARY KEY,
  "user_id" VARCHAR(255) NOT NULL,
  "message" VARCHAR(255) NOT NULL
);