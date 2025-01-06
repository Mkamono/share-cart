module "lb-http" {
  source  = "terraform-google-modules/lb-http/google//modules/serverless_negs"
  version = "~> 10.0"

  name    = "web-lb"
  project = local.project_id

  ssl                             = true
  https_redirect                  = true
  managed_ssl_certificate_domains = [local.web_domain]
  labels                          = { "service" = "share-cart" }

  backends = {
    default = {
      description = "share cart web load balancer"
      groups = [
        {
          group = google_compute_region_network_endpoint_group.web_neg.id
        }
      ]
      enable_cdn = false

      iap_config = {
        enable = false
      }
      log_config = {
        enable = true
      }
    }
  }
}


resource "google_compute_region_network_endpoint_group" "web_neg" {
  provider              = google-beta
  name                  = "web-neg"
  network_endpoint_type = "SERVERLESS"
  region                = local.region
  cloud_run {
    service = google_cloud_run_v2_service.web.name
  }
}
