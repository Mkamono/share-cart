
# resource "github_repository" "share-cart" {
#   name       = "share-cart"
#   visibility = "public"

#   allow_merge_commit = false
#   allow_rebase_merge = false
#   allow_squash_merge = true

#   allow_update_branch    = true
#   delete_branch_on_merge = true

#   has_downloads = false
#   has_issues    = true
#   has_projects  = false
#   has_wiki      = false

#   squash_merge_commit_title = "PR_TITLE"
# }

# resource "github_branch_protection" "share-cart-protecnion" {
#   repository_id = github_repository.share-cart.node_id
#   pattern       = "main"

#   required_pull_request_reviews {
#     required_approving_review_count = 0
#   }

#   enforce_admins = true

#   required_status_checks {
#     strict = true
#   }
# }
