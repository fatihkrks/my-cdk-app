package main

import (
	"my-cdk-app/vpc"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := awscdk.NewApp(nil)

	vpc.NewVpcStack(app, "basevpc", &vpc.VpcStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil

}
