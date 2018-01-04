package make

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

// VERSION is the version of go-make.
const VERSION = "0.1.0"

// CLIApp returns a cli.App with default build and clean commands.
func CLIApp(suite *Suite) *cli.App {
	app := cli.NewApp()
	app.Name = "make"
	app.Usage = "a go build suite"
	app.Version = VERSION

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "parallel, p",
			Usage: "parallelize targets",
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name: "build",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "os",
					Usage: "The operating system to build for.",
					Value: "native",
				},
				cli.StringFlag{
					Name:  "arch",
					Usage: "The architecture to build for.",
					Value: "native",
				},
				cli.BoolFlag{
					Name:  "release",
					Usage: "If set all available platforms will be built.",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Bool("release") {
					// Build all.
					var target Target
					targets := suite.LookupBuildTargets()
					if c.GlobalBool("parallel") {
						target = Parallelize(targets...)
					} else {
						target = Concatenate(false, targets...)
					}

					err := suite.Execute(target)
					return err
				}

				// Build only one.
				platform, err := ParsePlatform(c.String("os"), c.String("arch"))
				if err != nil {
					return cli.NewExitError(err, -1)
				}

				err = suite.ExecuteNamedTarget(BuildTargetNamePrefix + platform.String())
				if IsNotFound(err) {
					return cli.NewExitError(fmt.Sprintf("platform %s is not supported", platform), -2)
				} else if err != nil {
					return cli.NewExitError(err, -2)
				}

				return nil
			},
		},
		cli.Command{
			Name: "clean",
			Action: func(c *cli.Context) error {
				var target Target
				targets := suite.LookupCleanTargets()
				if c.GlobalBool("parallel") {
					target = Parallelize(targets...)
				} else {
					target = Concatenate(false, targets...)
				}

				err := suite.Execute(target)
				if err != nil {
					return cli.NewExitError(err, -1)
				}

				return nil
			},
		},
	}

	return app
}
