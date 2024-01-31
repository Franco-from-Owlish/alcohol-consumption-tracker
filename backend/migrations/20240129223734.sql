-- Modify "recipe_ingredients" table
ALTER TABLE "public"."recipe_ingredients" DROP CONSTRAINT "fk_recipe_ingredients_ingredient", DROP CONSTRAINT "fk_recipe_ingredients_recipe", ADD COLUMN "amount" numeric NULL, ADD COLUMN "unit" character varying(12) NULL;
