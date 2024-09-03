table "items" {
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
  foreign_key "items_market_id_fk" {
    columns     = [column.market_id]
    ref_columns = [table.markets.column.id]
    on_delete   = NO_ACTION
    on_update   = NO_ACTION
  }

  column "name" {
    null = false
    type = text
  }
  unique "items_name" {
    columns = [column.name]
  }

  column "price" {
    null = false
    type = numeric
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
