resource "google_service_account" "cloud_run_service_account" {
  account_id   = "cloud-run-service-account"
  display_name = "Cloud Run Service Account"
}

resource "google_project_iam_member" "cloud_run_service_account_iam" {
  project = local.project_id
  for_each = toset([
    "roles/secretmanager.secretAccessor",
  ])
  role   = each.value
  member = "serviceAccount:${google_service_account.cloud_run_service_account.email}"
}

resource "google_cloud_run_service_iam_member" "web-public-access" {
  location = google_cloud_run_v2_service.web.location
  project  = google_cloud_run_v2_service.web.project
  service  = google_cloud_run_v2_service.web.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_cloud_run_service_iam_member" "api-public-access" {
  location = google_cloud_run_v2_service.api.location
  project  = google_cloud_run_v2_service.api.project
  service  = google_cloud_run_v2_service.api.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}
