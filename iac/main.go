package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvVar struct {
	Key   string
	Value string
	Type  string
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		envVars := []EnvVar{
			{Key: "JWT_SECRET", Type: "SECRET"},
			{
				Key:  "GRAPHQL_PLAYGROUND_PASSWORD",
				Type: "SECRET",
			},
			{
				Key:  "MONGODB_DATABASE_NAME",
				Type: "SECRET",
			},
		}
		for i := range envVars {
			envVar := &envVars[i]
			name := fmt.Sprintf("APP_%s", envVar.Key)
			envVar.Value = os.Getenv(name)
			if envVar.Value == "" {
				panic("Please set the environment variable " + name)
			}
		}

		appSpecServiceEnvArgs := make(digitalocean.AppSpecServiceEnvArray, len(envVars))
		for i, envVar := range envVars {
			key := pulumi.String(envVar.Key)
			value := pulumi.String(envVar.Value)
			envType := pulumi.String(envVar.Type)

			envArg := digitalocean.AppSpecServiceEnvArgs{
				Key:   key,
				Type:  envType,
				Value: value,
			}
			appSpecServiceEnvArgs[i] = envArg
		}

		appSpecServiceEnvArgs = append(appSpecServiceEnvArgs, digitalocean.AppSpecServiceEnvArgs{
			Key:   pulumi.String("ENVIRONMENT"),
			Type:  pulumi.String("GENERAL"),
			Value: pulumi.String("production"),
		})

		portString := os.Getenv("APP_PORT")
		if portString == "" {
			panic("Please set the environment variable APP_PORT")
		}
		port, err := strconv.Atoi(portString)
		if err != nil {
			panic("APP_PORT must be a number")
		}

		db, err := digitalocean.NewDatabaseCluster(
			ctx,
			"mongodb",
			&digitalocean.DatabaseClusterArgs{
				Region:    pulumi.String("fra1"),
				Version:   pulumi.String("6.0"),
				Size:      pulumi.String("db-s-1vcpu-1gb"),
				NodeCount: pulumi.Int(1),
				Engine:    pulumi.String("mongodb"),
			},
		)
		if err != nil {
			return err
		}

		appSpecServiceEnvArgs = append(appSpecServiceEnvArgs, digitalocean.AppSpecServiceEnvArgs{
			Key:  pulumi.String("MONGODB_CONNECTION_STRING"),
			Type: pulumi.String("SECRET"),
			Value: pulumi.Sprintf(
				"mongodb+srv://%v:%v@%v/%v?replicaSet=%v&tls=true&authSource=admin",
				db.User,
				db.Password,
				db.Host,
				envVars[2].Value, // MONGODB_DATABASE_NAME
				db.Name,
			),
		})

		app, err := digitalocean.NewApp(ctx, "habit-tracker", &digitalocean.AppArgs{
			Spec: &digitalocean.AppSpecArgs{
				Name:   pulumi.String("habit-tracker"),
				Region: pulumi.String("fra1"),
				Databases: digitalocean.AppSpecDatabaseArray{
					&digitalocean.AppSpecDatabaseArgs{
						Name:        pulumi.String("mongodb"),
						ClusterName: db.Name,
						Engine:      pulumi.String("MONGODB"),
						Production:  pulumi.Bool(true),
					},
				},
				Services: digitalocean.AppSpecServiceArray{
					&digitalocean.AppSpecServiceArgs{
						Name:             pulumi.String("gql-api"),
						InstanceCount:    pulumi.Int(1),
						InstanceSizeSlug: pulumi.String("basic-xs"),
						Image: digitalocean.AppSpecServiceImageArgs{
							RegistryType: pulumi.String("DOCR"),
							Repository: pulumi.String(
								"habit_tracker",
							),
							Tag: pulumi.String("latest"),
							DeployOnPushes: digitalocean.AppSpecServiceImageDeployOnPushArray{
								digitalocean.AppSpecServiceImageDeployOnPushArgs{
									Enabled: pulumi.Bool(true),
								},
							},
						},
						Envs:     appSpecServiceEnvArgs,
						HttpPort: pulumi.Int(port),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		trustedSource, err := digitalocean.NewDatabaseFirewall(
			ctx,
			"trusted-souce",
			&digitalocean.DatabaseFirewallArgs{
				ClusterId: db.ID(),
				Rules: digitalocean.DatabaseFirewallRuleArray{
					&digitalocean.DatabaseFirewallRuleArgs{
						Type:  pulumi.String("app"),
						Value: app.ID(),
					},
				},
			},
		)

		ctx.Export("db_name", db.Name)
		ctx.Export("db_uri", db.Uri)

		ctx.Export("app_name", app.Spec.Name())
		ctx.Export("app_url", app.LiveUrl)

		ctx.Export("trusted_source", trustedSource)

		return nil
	})
}
