resource "azurerm_resource_group" "this" {
  name = var.name
  location = "East US"
  tags = {
    "Cost" = "dsd"
    "Environment" = "dsd"
    "Project" = "sd"
  }
}