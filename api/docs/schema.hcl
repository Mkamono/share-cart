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
    type = text
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

table "auth_subjects" {
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

  column "user_id" {
    null = false
    type = integer
  }
  foreign_key "user_fk" {
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
    on_delete = CASCADE
    on_update = NO_ACTION
  }

  column "subject" {
    null = false
    type = text
  }
  unique "subject" {
    columns = [column.subject]
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
