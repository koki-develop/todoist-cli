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

	// --project-id (for create)
	flagSectionProjectIDForCreate = &flags.String{
		Flag: &flags.Flag{
			Name:        "project-id",
			Description: "project id this section should belong to",
			Required:    true,
		},
	}

	// --order
	flagSectionOrder = &flags.Int{
		Flag: &flags.Flag{
			Name:        "order",
			Description: "order among other sections in a project",
		},
	}

	// --name
	flagSectionNameForUpdate = &flags.String{
		Flag: &flags.Flag{
			Name:        "name",
			Description: "section name",
			Required:    true,
		},
	}
)

// flags for tasks list
var (
	// --project-id
	flagTasksListProjectID = &flags.String{
		Flag: &flags.Flag{
			Name:        "project-id",
			Description: "filter tasks by project id",
		},
	}

	// --section-id
	flagTasksListSectionID = &flags.String{
		Flag: &flags.Flag{
			Name:        "section-id",
			Description: "filter tasks by section id",
		},
	}

	// --label
	flagTasksListLabel = &flags.String{
		Flag: &flags.Flag{
			Name:        "label",
			Description: "filter tasks by label name",
		},
	}

	// --filter
	flagTasksListFilter = &flags.String{
		Flag: &flags.Flag{
			Name:        "filter",
			Description: "filter by any supported filter",
		},
	}

	// --lang
	flagTasksListLang = &flags.String{
		Flag: &flags.Flag{
			Name:        "lang",
			Description: "IETF language tag defining what language filter is written in, if differs from default English",
		},
	}

	// --ids
	flagTasksListIDs = &flags.Strings{
		Flag: &flags.Flag{
			Name:        "ids",
			Description: "a list of the task ids to retrieve",
		},
	}
)
