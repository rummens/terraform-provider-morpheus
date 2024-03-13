package morpheus

import (
	"context"
	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceServiceCatalogPlaceCatalogOrder() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a Morpheus service catalog place catalog order, resource",
		CreateContext: resourceServiceCatalogPlaceCatalogOrderCreate,

		// this resource can only be used to make an order, there is no state management/crud
		ReadContext:   noOp,
		UpdateContext: noOp,
		DeleteContext: noOp,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "The id of the created order",
				Computed:    true,
			},
			"order_item": {
				Type:        schema.TypeList,
				Description: "Order item configuration",
				Optional:    false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"catalog_item_type_id": {
							Type:        schema.TypeInt,
							Description: "The id of the catalog item type to be ordered",
							Required:    false,
						},
						"catalog_item_type_name": {
							Type:        schema.TypeString,
							Description: "The name of the catalog item type to be ordered",
							Required:    false,
						},
						"config": {
							Type:        schema.TypeString,
							Description: "JSON object of key/value pairs which represent the catalog item type associated inputs/option types",
							Required:    true,
						},
						"context": {
							Type:        schema.TypeString,
							Description: "If workflow catalog item type, specify 'instance', 'server' or 'appliance'",
							Required:    false,
						},
						"target": {
							Type:        schema.TypeInt,
							Description: "If workflow catalog item type, resource (Instance or Server) id for context when running workflow",
							Required:    false,
						},
					},
				},
			},
		},
	}
}

func resourceServiceCatalogPlaceCatalogOrderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*morpheus.Client)

	var diags diag.Diagnostics

	catalogOrder := make(map[string]interface{})
	catalogOrder["items"] = parseCatalogOrderItems(d.Get("order_item").([]interface{}))

	req := &morpheus.Request{
		Body: map[string]interface{}{
			"order": catalogOrder,
		},
	}

	resp, err := client.CreateCatalogItem(req)
	if err != nil {
		log.Printf("API FAILURE: %s - %s", resp, err)
		return diag.FromErr(err)
	}
	log.Printf("API RESPONSE: %s", resp)

	result := resp.Result.(*morpheus.CreateCatalogItemResult)
	catalogItemType := result.CatalogItem

	d.SetId(int64ToString(catalogItemType.ID))

	return diags
}

func parseCatalogOrderItems(orderItemList []interface{}) []map[string]interface{} {
	var orderItems []map[string]interface{}

	for i := 0; i < len(orderItemList); i++ {
		oI := map[string]interface{}{}
		oIConfig := orderItemList[i].(map[string]interface{})
		for k, v := range oIConfig {
			switch k {
			case "catalog_item_type_id":
				oI["type"] = map[string]int{
					"id": v.(int),
				}
			case "catalog_item_type_name":
				oI["type"] = map[string]string{
					"name": v.(string),
				}
			case "config":
				oI["config"] = v.(string)
			case "context":
				oI["context"] = v.(string)
			case "target":
				oI["target"] = v.(int)
			}
		}
		orderItems = append(orderItems, oI)
	}

	return orderItems

}

// not sure if we can nil the config so providing a noOp which returns empty diag.Diagnostics
func noOp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}
