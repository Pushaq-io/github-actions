include "root" {
  path = find_in_parent_folders()
}



terraform {
  source = "../../../../modules/rg"
}

inputs = {
  name = "rg-github-actions-qa-01-app1"
}