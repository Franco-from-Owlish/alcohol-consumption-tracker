variable "db_host" {
  type    = string
  default = getenv("DB_HOST")
}

variable "db_driver" {
  type    = string
  default = getenv("DB_DRIVER")
}

variable "db_user" {
  type    = string
  default = getenv("DB_USER")
}

variable "db_password" {
  type    = string
  default = getenv("DB_PASSWORD")
}

variable "db_name" {
  type    = string
  default = getenv("DB_NAME")
}

variable "db_port" {
  type    = string
  default = getenv("DB_PORT")
}

locals {
  db_url  = "${var.db_driver}://${var.db_user}:${var.db_password}@${var.db_host}:${var.db_port}/${var.db_name}?sslmode=disable"
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/migrations/main.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev"
  url = "docker://postgres/15/dev"
  migration {
    dir = "file://backend/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "deploy" {
  url = local.db_url
  migration {
    dir = "file://backend/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}