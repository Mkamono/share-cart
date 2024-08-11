table "users" {
  schema = schema.public
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
  unique "user_name" {
    columns = [column.name]
  }

  // system managed columns
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
}

table "auth_subjects" {
  schema = schema.public
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
  foreign_key "auth_subjects_user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_delete   = NO_ACTION
    on_update   = NO_ACTION
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
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
}
