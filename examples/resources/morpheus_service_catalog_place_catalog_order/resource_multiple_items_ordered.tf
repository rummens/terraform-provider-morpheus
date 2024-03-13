resource "morpheus_service_catalog_place_catalog_order" "tf_example_service_catalog_order_on_name" {

  order_item {
    catalog_item_type_name = "Ubuntu 22.04"
    config    = <<EOF
      {
        "customInstanceName": "My Ubuntu Instance",
        "optionTypeKey 2":" value 2",
        "optionTypeKey 3":" value 3"
      }
    EOF
  }

  order_item {
    catalog_item_type_id = "9"
    config    = <<EOF
      {
        "customInstanceName": "My Ubuntu Instance 2",
        "optionTypeKey 2":" value 2",
        "optionTypeKey 3":" value 3"
      }
    EOF
  }

}