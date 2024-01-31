-- Create "cocktails" table
CREATE TABLE "public"."cocktails" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(48) NULL,
  "total_alcohol" numeric NULL DEFAULT 0,
  "on_menu" boolean NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- Create index "cocktails_name_key" to table: "cocktails"
CREATE UNIQUE INDEX "cocktails_name_key" ON "public"."cocktails" ("name");
-- Create index "idx_cocktails_deleted_at" to table: "cocktails"
CREATE INDEX "idx_cocktails_deleted_at" ON "public"."cocktails" ("deleted_at");
-- Create "patrons" table
CREATE TABLE "public"."patrons" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "first_name" character varying(48) NULL,
  "last_name" character varying(48) NULL,
  "total_alcohol" numeric NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_patrons_deleted_at" to table: "patrons"
CREATE INDEX "idx_patrons_deleted_at" ON "public"."patrons" ("deleted_at");
-- Create "patron_cocktails" table
CREATE TABLE "public"."patron_cocktails" (
  "patron_id" bigint NOT NULL,
  "cocktail_id" bigint NOT NULL,
  PRIMARY KEY ("patron_id", "cocktail_id"),
  CONSTRAINT "fk_patron_cocktails_cocktail" FOREIGN KEY ("cocktail_id") REFERENCES "public"."cocktails" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_patron_cocktails_patron" FOREIGN KEY ("patron_id") REFERENCES "public"."patrons" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "ingredients" table
CREATE TABLE "public"."ingredients" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(48) NULL,
  "abv" integer NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_ingredients_deleted_at" to table: "ingredients"
CREATE INDEX "idx_ingredients_deleted_at" ON "public"."ingredients" ("deleted_at");
-- Create "recipes" table
CREATE TABLE "public"."recipes" (
  "id" bigserial NOT NULL,
  "cocktail_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_cocktails_recipe" FOREIGN KEY ("cocktail_id") REFERENCES "public"."cocktails" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "recipe_ingredients" table
CREATE TABLE "public"."recipe_ingredients" (
  "recipe_id" bigint NOT NULL,
  "ingredient_id" bigint NOT NULL,
  PRIMARY KEY ("recipe_id", "ingredient_id"),
  CONSTRAINT "fk_recipe_ingredients_ingredient" FOREIGN KEY ("ingredient_id") REFERENCES "public"."ingredients" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_recipe_ingredients_recipe" FOREIGN KEY ("recipe_id") REFERENCES "public"."recipes" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
