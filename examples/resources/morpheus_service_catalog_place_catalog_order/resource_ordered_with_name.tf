
resource "morpheus_service_catalog_place_catalog_order" "tf_order_go_dev_server_name" {

  order_item {

    catalog_item_type_name = "Go Development Server"
    config                 = <<EOF
      {
        "instanceName": "Go Development Server 1211",
        "group": "1",
        "cloud": "2",
        "resourcePool": "pool-8",
        "goVersion": "1.21.1"
      }
    EOF
  }

}