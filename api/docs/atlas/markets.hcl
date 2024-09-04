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

table "market_images" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }

  column "market_id" {
    null = false
    type = uuid
  }
  foreign_key "market_images_market_id_fk" {
    columns     = [column.market_id]
    ref_columns = [table.markets.column.id]
  }

  column "url" {
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
