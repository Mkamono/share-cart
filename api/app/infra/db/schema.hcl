schema "main" {}

table "users" {
  schema = schema.main
  column "id" {
    null = false
    type = integer
    identity {
      generated = ALWAYS
    }
  }
  primary_key {
    columns = [column.id]
  }

  column "name" {
    null = false
    type = character_varying(255)
  }
  unique "name" {
    columns = [column.name]
  }

  // system managed columns
  column "created_at" {
    null = false
    type = timestamp
  }

  column "updated_at" {
    null = false
    type = timestamp
  }

  column "deleted_at" {
    null = true
    type = timestamp
  }
}
