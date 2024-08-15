resource "google_artifact_registry_repository" "api-repo" {
  location      = local.region
  repository_id = "api-repo"
  description   = "share cart api repository"
  format        = "DOCKER"
}
