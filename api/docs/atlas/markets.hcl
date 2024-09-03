table "markets" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }

  column "name" {
    null = false
    type = text
  }
  unique "markets_name" {
    columns = [column.name]
  }

  column "description" {
    null = false
    type = text
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
