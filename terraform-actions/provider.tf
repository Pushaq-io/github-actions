terraform {
  backend "remote" {
    hostname = "app.terraform.io"
    organization = "caltamirano"
    workspaces {
      name = "terraform-github-actions"
    }
  }
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.75.0"
    }
  }
}

provider "azurerm" {
  # Configuration options
  features {
    
  }
}

