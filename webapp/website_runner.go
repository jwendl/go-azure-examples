package main

import (
	"context"
	"fmt"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/resources"
	"github.com/subosito/gotenv"
)

var (
	webName               = "az-samples-go-web-" + helpers.GetRandomLetterSequence(5)
	resourceGroupName     = "canttouchthis"
	resourceGroupLocation = "southcentralus"
)

func main() {
	parseArgs()
	exampleCreateWebSite()
}

func parseArgs() error {
	gotenv.Load()
	err := helpers.ParseArgs()
	if err != nil {
		return fmt.Errorf("cannot parse args: %v", err)
	}

	return nil
}

func exampleCreateWebSite() {
	helpers.SetResourceGroupName(resourceGroupName)
	ctx := context.Background()
	defer resources.Cleanup(ctx)
	_, err := resources.CreateGroup(ctx, helpers.ResourceGroupName())
	if err != nil {
		helpers.PrintAndLog("Failed Web Site creation.")
		helpers.PrintAndLog(err.Error())
	}
	_, err = CreateAppServicePlan(ctx, webName)
	if err != nil {
		helpers.PrintAndLog("Failed App Service Plan creation.")
		helpers.PrintAndLog(err.Error())
	} else {
		helpers.PrintAndLog("Created App Service Plan")
		_, err = CreateWebSite(ctx, webName)
		if err != nil {
			helpers.PrintAndLog(err.Error())
		} else {
			helpers.PrintAndLog("Created Website")
		}
	}
}
