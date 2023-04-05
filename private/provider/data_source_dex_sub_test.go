package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"dex": providerserver.NewProtocol6WithError(NewDexProvider()),
	}
)

func TestDexSubDataSource_Subjects(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "dex_sub" "gitlab" {
						user_id = "vpupkin"
						conn_id = "gitlab"
					}

					data "dex_sub" "github" {
						user_id = "vasily.pupkin"
						conn_id = "github"
					}
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.dex_sub.gitlab", "sub", "Cgd2cHVwa2luEgZnaXRsYWI"),
					resource.TestCheckResourceAttr("data.dex_sub.github", "sub", "Cg12YXNpbHkucHVwa2luEgZnaXRodWI"),
				),
			},
		},
	})

}
