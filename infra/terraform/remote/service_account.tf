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
