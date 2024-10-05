# supabase
resource "google_secret_manager_secret" "supabase_user" {
  secret_id = "supabase_user"
  labels = {
    service = "supabase"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "supabase_password" {
  secret_id = "supabase_password"
  labels = {
    service = "supabase"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "supabase_db_dev" {
  secret_id = "supabase_db_dev"
  labels = {
    service = "supabase"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "supabase_host" {
  secret_id = "supabase_host"
  labels = {
    service = "supabase"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "supabase_port" {
  secret_id = "supabase_port"
  labels = {
    service = "supabase"
  }
  replication {
    auto {}
  }
}

# auth0
resource "google_secret_manager_secret" "auth0_audience" {
  secret_id = "auth0_audience"
  labels = {
    service = "auth0"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "auth0_domain" {
  secret_id = "auth0_domain"
  labels = {
    service = "auth0"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "auth0_client_id" {
  secret_id = "auth0_client_id"
  labels = {
    service = "auth0"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "auth0_client_secret" {
  secret_id = "auth0_client_secret"
  labels = {
    service = "auth0"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "auth0_secret" {
  secret_id = "auth0_secret"
  labels = {
    service = "auth0"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "auth0_base_url" {
  secret_id = "auth0_base_url"
  labels = {
    service = "auth0"
  }
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret" "auth0_issuer_base_url" {
  secret_id = "auth0_issuer_base_url"
  labels = {
    service = "auth0"
  }
  replication {
    auto {}
  }
}
