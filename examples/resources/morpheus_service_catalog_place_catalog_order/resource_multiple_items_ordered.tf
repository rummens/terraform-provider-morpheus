
resource "morpheus_service_catalog_place_catalog_order" "tf_order_go_dev_server_multiple" {

  order_item {

    catalog_item_type_name = "Go Development Server"
    config                 = <<EOF
      {
        "instanceName": "My Go Development Server 1211",
        "group": "1",
        "cloud": "2",
        "resourcePool": "pool-8",
        "goVersion": "1.21.1"
      }
    EOF
  }

  order_item {

    catalog_item_type_id = "1"
    config               = <<EOF
      {
        "instanceName": "Go Development Server 1201",
        "group": "1",
        "cloud": "2",
        "resourcePool": "pool-8",
        "goVersion": "1.20.1"
      }
    EOF
  }

}