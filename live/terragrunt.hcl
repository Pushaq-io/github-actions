generate "provider" {
  path = "provider.tf"
  if_exists = "overwrite_terragrunt"
  contents = <<EOF
provider "azurerm" {
  # Configuration options
  features {
    
  }
}
EOF
}