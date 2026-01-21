schema "gcook" {}

table "recipes" {
  schema = schema.gcook
  column "id" {
    type = uuid
  }
  column "user_id" {
    type = varchar(26)
  }
  column "title" {
    type = varchar(255)
  }
  column "description" {
    type = varchar(255)
  }
  column "created_at" {
    type = timestamp
  }
  column "updated_at" {
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
}

table "follows" {
  schema = schema.gcook
  column "follower_id" {
    type = char(26)
  }
  column "followee_id" {
    type = char(26)
  }
  column "created_at" {
    type = timestamp
  }
  primary_key {
    columns = [column.follower_id, column.followee_id]
  }
  index "follows_follower_id_idx" {
    columns = [column.follower_id]
  }
  index "follows_followee_id_idx" {
    columns = [column.followee_id]
  }
}
