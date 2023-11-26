CREATE TABLE State (
    userId VARCHAR(255) PRIMARY KEY,
    dataJson JSONB,
    updatedAt TIMESTAMP,
    createdAt TIMESTAMP 
);