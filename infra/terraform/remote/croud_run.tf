
resource "google_cloud_run_v2_service" "api" {
  name     = "api"
  location = local.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "asia-northeast1-docker.pkg.dev/coastal-fiber-430001-s0/api-repo/api:latest"
      resources {
        cpu_idle = true
        limits = {
          cpu    = "2"
          memory = "512Mi"
        }
      }
    }
  }
}
