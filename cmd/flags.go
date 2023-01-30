package cmd

import (
	"fmt"
	"strings"

	"github.com/koki-develop/todoist-cli/pkg/flags"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	// --api-token
	flagAPIToken = &flags.String{
		Flag: &flags.Flag{
			Name:        "api-token",
			Description: "todoist api token",
		},
	}

	// -f, --format
	flagFormat = &flags.String{
		Flag: &flags.Flag{
			Name:        "format",
			ShortName:   "f",
			Description: fmt.Sprintf("output format (%s)", strings.Join(renderer.Formats, "|")),
		},
	}
)

// flags for projects
var (
	// --parent-id
	flagProjectParentID = &flags.String{
		Flag: &flags.Flag{
			Name:        "parent-id",
			Description: "parent project id",
		},
	}

	// --name
	flagProjectName = &flags.String{
		Flag: &flags.Flag{
			Name:        "name",
			Description: "name of the project",
		},
	}

	// --color
	flagProjectColor = &flags.String{
		Flag: &flags.Flag{
			Name:        "color",
			Description: "the color of the project icon",
		},
	}

	// --favorite
	flagProjectFavorite = &flags.Bool{
		Flag: &flags.Flag{
			Name:        "favorite",
			Description: "whether the project is a favorite",
		},
	}
)

// flags for sections
var (
	// --project-id
	flagSectionProjectID = &flags.String{
		Flag: &flags.Flag{
			Name:        "project-id",
			Description: "filter sections by project id",
		},
	}
)
