resource "morpheus_service_catalog_place_catalog_order" "tf_example_service_catalog_order_on_id" {

  order_item {
    catalog_item_type_id = "9"
    config               = <<EOF
      {
        "customInstanceName": "My Ubuntu Instance 2",
        "optionTypeKey 2":" value 2",
        "optionTypeKey 3":" value 3"
      }
    EOF
  }

}