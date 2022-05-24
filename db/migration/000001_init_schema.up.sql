CREATE TABLE "accounts" (
  "id" SERIAL PRIMARY KEY,
  "owner" varchar(255) NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" int PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" int PRIMARY KEY,
  "from_account_id" bigint,
  "to_account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX "accounts_index_0" ON "accounts" ("owner");

CREATE INDEX "entries_index_1" ON "entries" ("account_id");

CREATE INDEX "transfers_index_2" ON "transfers" ("from_account_id");

CREATE INDEX "transfers_index_3" ON "transfers" ("to_account_id");

CREATE INDEX "transfers_index_4" ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'it must be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");