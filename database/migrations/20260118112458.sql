-- Create "follows" table
CREATE TABLE "follows" (
  "follower_id" character(26) NOT NULL,
  "followee_id" character(26) NOT NULL,
  "created_at" timestamp NOT NULL,
  PRIMARY KEY ("follower_id", "followee_id")
);
-- Create index "follows_followee_id_idx" to table: "follows"
CREATE INDEX "follows_followee_id_idx" ON "follows" ("followee_id");
-- Create index "follows_follower_id_idx" to table: "follows"
CREATE INDEX "follows_follower_id_idx" ON "follows" ("follower_id");
