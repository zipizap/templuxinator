package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"

	"github.com/ghodss/yaml"
	cli "github.com/urfave/cli/v2"

	"github.com/Masterminds/sprig"
)

func initializations() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})
	spew.Config.Indent = "  "
}

/*
  ## USEFULL NOTES :)

  spew.Dump(...any-object...)

  log.Fatal(err)
  log.Info("info message")
*/

// This is just a placeholder to store flag values, so they can be set by app.Flags and from then on read from here
var ArgsAndFlags = struct {
	args  []string
	flags struct {
		templateFile string
		valuesFile   string
		outputFile   string
	}
}{}

// This function is where all the interesting stuff happens
func app_action(c *cli.Context) error {

	/*
		// *arguments* - working with them
		if c.NArg() > 0 {
			ArgsAndFlags.args = c.Args().Slice()
			fmt.Printf("First arg is: %q \n", ArgsAndFlags.args[0])
		}

		// *flags* - working with them
		fmt.Printf("Flag --config is: %+v \n", ArgsAndFlags.flags.configfile)
		fmt.Printf("Flag --config is: %+v \n", c.String("configfile"))

		// Exit-code
		if ArgsAndFlags.args[0] == "exit" {
			// exit from an err:
			err := errors.New("This is my homemade error! Spuff!")
			return cli.Exit(err.Error(), 86)

			// or exit "manually":
			//return cli.Exit("My exit message here - bye bye with exit code 86 ", 86)
		}
		return nil
	*/

	templateFileBytes, err := ioutil.ReadFile(ArgsAndFlags.flags.templateFile)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	templateFileString := string(templateFileBytes)
	t, err := template.
		New(ArgsAndFlags.flags.templateFile).
		Option("missingkey=error").
		Funcs(sprig.FuncMap()).
		Parse(templateFileString)
	if err != nil {
		return cli.Exit(err.Error(), 2)
	}
	// read values and use it to execute template
	valuesAsMap, err := readValuesFromYamlFile(ArgsAndFlags.flags.valuesFile)
	//spew.Dump(valuesAsMap)
	if err != nil {
		return cli.Exit(err.Error(), 3)
	}

	foutput, err := os.Create(ArgsAndFlags.flags.outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer foutput.Close()

	multiwriter := io.MultiWriter(os.Stdout, foutput)
	err = t.Execute(multiwriter, valuesAsMap)
	if err != nil {
		return cli.Exit(err.Error(), 4)
	}

	return nil
}

// :) round of applause to https://github.com/amitsaha/golang-templates-demo/blob/master/render-arbitrary-template/main.go
func readValuesFromYamlFile(path string) (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{}, 0)
	if err := yaml.Unmarshal(content, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	initializations()

	app := &cli.App{
		Name:    "templuxinator",
		Usage:   "Templating similar to helm-chart, but for any files :)",
		Version: "1.0.0", // --version | -v
		Action:  app_action,
		Flags: []cli.Flag{
			// Global options flags
			&cli.StringFlag{
				Name:        "template",
				Aliases:     []string{"t"},
				Usage:       "Load template from `mytemplatefile`",
				Destination: &ArgsAndFlags.flags.templateFile,
				//Value:       "./mytemplatefile",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "values",
				Aliases:     []string{"f"},
				Usage:       "Load values from `myvaluesfile`",
				Destination: &ArgsAndFlags.flags.valuesFile,
				//Value:       "./myvaluesfile",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Output file to create with generated text `myresultfile`",
				Destination: &ArgsAndFlags.flags.outputFile,
				Value:       "./myresultfile",
				Required:    true,
			},
		},
	}
	app.UseShortOptionHandling = true // let bool-shortnames to be joined: -a -b -c FILE => -abc FILE
	app.EnableBashCompletion = true   // and then users must copy a file "bash_autocomplete" and setup shell with: PROG=myprogram source path/to/cli/autocomplete/bash_autocomplete

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
