// https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password
resource random_password baskets {
  length = 16
}

// https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource
// https://www.terraform.io/language/resources/provisioners/local-exec
resource null_resource init_baskets_db {
  provisioner "local-exec" {
    command     = "psql --file sql/init_service_db.psql -v db=$DB -v user=$USER -v pass=$PASS ${local.db_conn}/postgres"
    environment = {
      DB   = "baskets"
      USER = "baskets_user"
      PASS = random_password.baskets.result
    }
  }
  depends_on = [
    null_resource.init_db,
    random_password.baskets
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/secret_v1
resource kubernetes_secret_v1 baskets {
  metadata {
    name      = "baskets-secrets"
    namespace = local.project
  }

  data = {
    PG_CONN = "host=${local.db_host} port=${local.db_port} dbname=baskets user=baskets_user password=${random_password.baskets.result} search_path=baskets,public"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
    null_resource.init_baskets_db
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment_v1
resource kubernetes_deployment_v1 baskets {
  metadata {
    name      = "baskets"
    namespace = local.project
    labels    = {
      app = "baskets"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        "app.kubernetes.io/name" = "baskets"
      }
    }
    template {
      metadata {
        name   = "baskets"
        labels = {
          "app.kubernetes.io/name" = "baskets"
        }
      }
      spec {
        hostname = "baskets"
        container {
          name              = "baskets"
          image             = "${local.aws_ecr_url}/baskets:latest"
          image_pull_policy = "Always"
          env_from {
            config_map_ref {
              name = "common-config-map"
            }
          }
          env_from {
            secret_ref {
              name = "baskets-secrets"
            }
          }
          port {
            protocol       = "TCP"
            container_port = 80
          }
          port {
            protocol       = "TCP"
            container_port = 9000
          }
          liveness_probe {
            http_get {
              path = "/liveness"
              port = 80
            }
            initial_delay_seconds = 3
            period_seconds        = 5
          }
        }
      }
    }
  }

  depends_on = [
    kubernetes_namespace_v1.namespace,
    kubernetes_config_map_v1.common,
    kubernetes_secret_v1.cosec,
    kubernetes_service_v1.nats,
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_v1
resource kubernetes_service_v1 baskets {
  metadata {
    name      = "baskets"
    namespace = local.project
    labels    = {
      app = "baskets"
    }
  }
  spec {
    selector = {
      "app.kubernetes.io/name" = "baskets"
    }
    session_affinity = "ClientIP"
    port {
      name        = "http"
      protocol    = "TCP"
      port        = 80
      target_port = 80
    }
    port {
      name        = "grpc"
      protocol    = "TCP"
      port        = 9000
      target_port = 9000
    }
    type = "NodePort"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
  ]
}

// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/ingress_v1
resource kubernetes_ingress_v1 baskets {
  metadata {
    name        = "baskets-ingress"
    namespace   = local.project
    annotations = {
      "alb.ingress.kubernetes.io/group.name"    = local.project
      "alb.ingress.kubernetes.io/scheme"        = "internet-facing"
      "alb.ingress.kubernetes.io/inbound-cidrs" = local.allowed_cidr_block
      "alb.ingress.kubernetes.io/target-type"   = "instance"
    }
  }

  spec {
    rule {
      http {
        path {
          path      = "/api/baskets"
          path_type = "Prefix"
          backend {
            service {
              name = "baskets"
              port {
                number = 80
              }
            }
          }
        }
        path {
          path      = "/baskets-spec/"
          path_type = "Prefix"
          backend {
            service {
              name = "baskets"
              port {
                number = 80
              }
            }
          }
        }
      }
    }
    ingress_class_name = "alb"
  }
  depends_on = [
    kubernetes_namespace_v1.namespace,
  ]
}
