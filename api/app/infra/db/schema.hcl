table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = integer
    identity {
      generated = ALWAYS
    }
  }
  column "create_time" {
    null = true
    type = date
  }
  column "name" {
    null = true
    type = character_varying(255)
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
  comment = "standard public schema"
}
