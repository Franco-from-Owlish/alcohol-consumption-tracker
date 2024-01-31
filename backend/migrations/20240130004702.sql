-- Modify "recipes" table
ALTER TABLE "public"."recipes" ADD COLUMN "created_at" timestamptz NULL, ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "deleted_at" timestamptz NULL;
-- Create index "idx_recipes_deleted_at" to table: "recipes"
CREATE INDEX "idx_recipes_deleted_at" ON "public"."recipes" ("deleted_at");
