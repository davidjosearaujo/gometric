terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

resource "docker_image" "gometric" {
  name = "gometric"
  build {
    context = "."
    tag     = ["gometric:dev"]
    label = {
      author : "davidjosearaujo"
    }
  }
  keep_locally = false
}

resource "docker_container" "test1" {
  image = docker_image.gometric.image_id
  name  = "gometric1"

  ports {
    internal = 7000
    external = 7001
  }
}

resource "docker_container" "test2" {
  image = docker_image.gometric.image_id
  name  = "gometric2"

  ports {
    internal = 7000
    external = 7002
  }
}

resource "docker_container" "test3" {
  image = docker_image.gometric.image_id
  name  = "gometric3"

  ports {
    internal = 7000
    external = 7003
  }
}
