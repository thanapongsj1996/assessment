CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    title TEXT,
    amount FLOAT,
    note TEXT,
    tags TEXT[]
);

INSERT INTO "public"."expenses"("title","amount","note","tags")
VALUES
    (E'strawberry smoothie',79,E'night market promotion discount 10 bath',E'{food,beverage}');