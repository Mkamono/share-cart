
resource "google_cloud_run_v2_service" "api" {
  name     = "api"
  location = local.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    service_account = google_service_account.cloud_run_service_account.email
    scaling {
      max_instance_count = 3
      min_instance_count = 0
    }
    containers {
      image = "asia-northeast1-docker.pkg.dev/coastal-fiber-430001-s0/api-repo/api:latest"
      resources {
        cpu_idle = true
        limits = {
          cpu    = "2"
          memory = "512Mi"
        }
      }
      ports {
        container_port = 8080
      }

      // シークレットの設定
      env {
        name = "POSTGRES_USER"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.supabase_user.secret_id
            version = "latest"
          }
        }
      }
      env {
        name = "POSTGRES_PASSWORD"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.supabase_password.secret_id
            version = "latest"
          }
        }
      }
      env {
        name = "POSTGRES_DB"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.supabase_db_dev.secret_id
            version = "latest"
          }
        }
      }
      env {
        name = "POSTGRES_HOST"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.supabase_host.secret_id
            version = "latest"
          }
        }
      }
      env {
        name = "POSTGRES_PORT"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.supabase_port.secret_id
            version = "latest"
          }
        }
      }
      env {
        name = "AUTH0_DOMAIN"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.auth0_domain.secret_id
            version = "latest"
          }
        }
      }
      env {
        name = "AUTH0_AUDIENCE"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.auth0_audience.secret_id
            version = "latest"
          }
        }
      }
      env {
        name  = "PROJECT_ID"
        value = local.project_id
      }
    }
  }
}

resource "google_cloud_run_v2_service" "web" {
  name     = "web"
  location = local.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    service_account = google_service_account.cloud_run_service_account.email
    scaling {
      max_instance_count = 3
      min_instance_count = 0
    }
    containers {
      image = "asia-northeast1-docker.pkg.dev/coastal-fiber-430001-s0/web-repo/web:latest"
      resources {
        cpu_idle = true
        limits = {
          cpu    = "2"
          memory = "512Mi"
        }
      }
      ports {
        container_port = 3000
      }

      // シークレットの設定
      env {
        name = "AUTH0_CLIENT_ID"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.auth0_client_id.secret_id
            version = "latest"
          }
        }
      }
      env {
        name = "AUTH0_CLIENT_SECRET"
        value_source {
          secret_key_ref {
            secret  = google_secret_manager_secret.auth0_client_secret.secret_id
            version = "latest"
          }
        }
      }
      env {
        name  = "API_HOST"
        value = google_cloud_run_v2_service.api.uri
      }
    }
  }
}
