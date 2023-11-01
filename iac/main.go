package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvVar struct {
	Key   string
	Value string
	Type  string
}

func main() {
	envVars := []EnvVar{
		{Key: "ENVIRONMENT", Type: "GENERAL"},
		{Key: "PORT", Type: "GENERAL"},
		{Key: "JWT_SECRET", Type: "SECRET"},
		{
			Key:  "GRAPHQL_PLAYGROUND_PASSWORD",
			Type: "SECRET",
		},
		{
			Key:  "MONGODB_CONNECTION_STRING",
			Type: "SECRET",
		},
		{
			Key:  "MONGODB_DATABASE_NAME",
			Type: "SECRET",
		},
	}
	for _, envVar := range envVars {
		name := fmt.Sprintf("APP_%s", envVar.Key)
		envVar.Value = os.Getenv(name)
		if envVar.Value == "" {
			panic("Please set the environment variable " + name)
		}
	}

	pulumi.Run(func(ctx *pulumi.Context) error {
		appSpecEnvArgs := make(digitalocean.AppSpecEnvArray, len(envVars))

		for i, envVar := range envVars {
			key := pulumi.String(envVar.Key)
			value := pulumi.String(envVar.Value)
			envType := pulumi.String(envVar.Type)

			envArg := digitalocean.AppSpecEnvArgs{
				Key:   key,
				Type:  envType,
				Value: value,
			}
			appSpecEnvArgs[i] = envArg
		}

		app, err := digitalocean.NewApp(ctx, "habit_tracker", &digitalocean.AppArgs{
			Spec: &digitalocean.AppSpecArgs{
				Name:   pulumi.String("gql_api"),
				Region: pulumi.String("fra1"),
				Services: digitalocean.AppSpecServiceArray{
					&digitalocean.AppSpecServiceArgs{
						InstanceCount:    pulumi.Int(1),
						InstanceSizeSlug: pulumi.String("basic-xs"),
					},
				},
				Envs: appSpecEnvArgs,
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("app_name", app.Spec.Name())

		return nil
	})
}
