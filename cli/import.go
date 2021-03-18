package cli

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/DataDrake/cuppa/version"
	"github.com/alecbcs/lookout/config"
	"github.com/alecbcs/lookout/database"
	"github.com/alecbcs/lookout/results"
	"github.com/alecbcs/lookout/update"
)

// Import adds a new app to the database.
var Import = cmd.Sub{
	Name:  "import",
	Alias: "im",
	Short: "Import a package in .yml format to the database.",
	Args:  &ImportArgs{},
	Run:   ImportRun,
}

// ImportArgs handles the specific arguments for the import command.
type ImportArgs struct {
	Filename string
}

type importData struct {
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	Source       string   `yaml:"source"`
	Dependencies []string `yaml:"dependencies"`
}

// ImportRun adds a new app entry to the database by parsing a YAML file.
// Import like the add command will check if the app is up to date
// before adding to the database.
func ImportRun(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*ImportArgs)
	db := database.Open(config.Global.Database.Path)
	defer db.Close()
	fileData, err := ioutil.ReadFile(args.Filename)
	if err != nil {
		log.Fatal(err)
	}
	var data importData
	err = yaml.Unmarshal(fileData, &data)
	if err != nil {
		log.Fatal(err)
	}
	update.Init(config.Global.Github.Key)
	result, found := update.CheckUpdate(data.Source)
	if !found {
		log.Fatal("Unable to find " + data.Name + " " + data.Source)
	}
	entry := results.New(
		data.Name,
		result.Location,
		result.Version,
		data.Source,
		version.NewVersion(data.Version),
		update.UpToDate(result.Version, version.NewVersion(data.Version)))
	entry.Dependencies = data.Dependencies
	database.Add(db, entry)
	database.ImportDeps(db, entry.ID, entry.Dependencies)
}
