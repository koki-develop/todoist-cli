package cmd

import (
	"fmt"
	"strings"

	"github.com/koki-develop/todoist-cli/pkg/flags"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

// basic flags
var (
	// --api-token
	flagAPIToken = &flags.String{Flag: &flags.Flag{
		Name:        "api-token",
		Description: "todoist api token",
	}}

	// -f, --format
	flagFormat = &flags.String{Flag: &flags.Flag{
		Name:        "format",
		ShortName:   "f",
		Description: fmt.Sprintf("output format (%s)", strings.Join(renderer.Formats, "|")),
	}}
)

// flags for projects create
var (
	// --parent-id
	flagProjectsCreateParentID = &flags.String{Flag: &flags.Flag{
		Name:        "parent-id",
		Description: "parent project id",
	}}

	// --color
	flagProjectsCreateColor = &flags.String{Flag: &flags.Flag{
		Name:        "color",
		Description: "the color of the project icon",
	}}

	// --favorite
	flagProjectsCreateFavorite = &flags.Bool{Flag: &flags.Flag{
		Name:        "favorite",
		Description: "whether the project is a favorite",
	}}
)

// flags for columns
var (
	flagColumns = &flags.Flag{Name: "columns", Description: "table columns"}

	// project
	flagColumnsProject = &flags.Strings{Flag: flagColumns,
		Default: []string{"id", "name"},
	}

	// section
	flagColumnsSection = &flags.Strings{Flag: flagColumns,
		Default: []string{"id", "name", "project_id"},
	}

	// task
	flagColumnsTask = &flags.Strings{Flag: flagColumns,
		Default: []string{"id", "content", "labels", "project_id", "section_id"},
	}
)

// flags for projects update
var (
	// --name
	flagProjectsUpdateName = &flags.String{Flag: &flags.Flag{
		Name:        "name",
		Description: "name of the project",
	}}

	// --color
	flagProjectsUpdateColor = &flags.String{Flag: &flags.Flag{
		Name:        "color",
		Description: "the color of the project icon",
	}}

	// --favorite
	flagProjectsUpdateFavorite = &flags.Bool{Flag: &flags.Flag{
		Name:        "favorite",
		Description: "whether the project is a favorite",
	}}
)

// flags for sections list
var (
	// --project-id
	flagSectionsListProjectID = &flags.String{Flag: &flags.Flag{
		Name:        "project-id",
		Description: "filter sections by project id",
	}}
)

// flags for sections create
var (
	// --project-id (for create)
	flagSectionsCreateProjectID = &flags.String{Flag: &flags.Flag{
		Name:        "project-id",
		Description: "project id this section should belong to",
		Required:    true,
	}}

	// --order
	flagSectionsCreateOrder = &flags.Int{Flag: &flags.Flag{
		Name:        "order",
		Description: "order among other sections in a project",
	}}
)

// flags for sections update
var (
	// --name
	flagSectionsUpdateName = &flags.String{Flag: &flags.Flag{
		Name:        "name",
		Description: "section name",
		Required:    true,
	}}
)

// flags for tasks list
var (
	// --project-id
	flagTasksListProjectID = &flags.String{Flag: &flags.Flag{
		Name:        "project-id",
		Description: "filter tasks by project id",
	}}

	// --section-id
	flagTasksListSectionID = &flags.String{Flag: &flags.Flag{
		Name:        "section-id",
		Description: "filter tasks by section id",
	}}

	// --label
	flagTasksListLabel = &flags.String{Flag: &flags.Flag{
		Name:        "label",
		Description: "filter tasks by label name",
	}}

	// --filter
	flagTasksListFilter = &flags.String{Flag: &flags.Flag{
		Name:        "filter",
		Description: "filter by any supported filter",
	}}

	// --lang
	flagTasksListLang = &flags.String{Flag: &flags.Flag{
		Name:        "lang",
		Description: "IETF language tag defining what language filter is written in, if differs from default English",
	}}

	// --ids
	flagTasksListIDs = &flags.Ints{Flag: &flags.Flag{
		Name:        "ids",
		Description: "a list of the task ids to retrieve",
	}}
)

// flags for tasks create
var (
	// --description
	flagTasksCreateDescription = &flags.String{Flag: &flags.Flag{
		Name:        "description",
		Description: "a description for the task",
	}}

	// --project-id
	flagTasksCreateProjectID = &flags.String{Flag: &flags.Flag{
		Name:        "project-id",
		Description: "task project id",
	}}

	// --section-id
	flagTasksCreateSectionID = &flags.String{Flag: &flags.Flag{
		Name:        "section-id",
		Description: "id of section to put task into",
	}}

	// --parent-id
	flagTasksCreateParentID = &flags.String{Flag: &flags.Flag{
		Name:        "parent-id",
		Description: "parent task id",
	}}

	// --order
	flagTasksCreateOrder = &flags.Int{Flag: &flags.Flag{
		Name:        "order",
		Description: "non-zero integer value to sort tasks under the same parent",
	}}

	// --labels
	flagTasksCreateLabels = &flags.Strings{Flag: &flags.Flag{
		Name:        "labels",
		Description: "the task's labels",
	}}

	// --priority
	flagTasksCreatePriority = &flags.Int{Flag: &flags.Flag{
		Name:        "priority",
		Description: "task priority 1 to 4",
	}}

	// --due-string
	flagTasksCreateDueString = &flags.String{Flag: &flags.Flag{
		Name:        "due-string",
		Description: "human defined task due tate",
	}}

	// --due-date
	flagTasksCreateDueDate = &flags.String{Flag: &flags.Flag{
		Name:        "due-date",
		Description: "specific date in YYYY-MM-DD format relative to user's timezone",
	}}

	// --due-datetime
	flagTasksCreateDueDatetime = &flags.String{Flag: &flags.Flag{
		Name:        "due-datetime",
		Description: "specific date and time in RFC3339 format in UTC",
	}}

	// --due-lang
	flagTasksCreateDueLang = &flags.String{Flag: &flags.Flag{
		Name:        "due-lang",
		Description: "2-letter code specifying language in case `due_string` is not written in English",
	}}

	// --assignee-id
	flagTasksCreateAssigneeID = &flags.String{Flag: &flags.Flag{
		Name:        "assignee-id",
		Description: "the responsible user id",
	}}
)

// flags for tasks update
var (
	// --content
	flagTasksUpdateContent = &flags.String{Flag: &flags.Flag{
		Name:        "content",
		Description: "task content",
	}}

	// --description
	flagTasksUpdateDescription = &flags.String{Flag: &flags.Flag{
		Name:        "description",
		Description: "a description for the task",
	}}

	// --labels
	flagTasksUpdateLabels = &flags.Strings{Flag: &flags.Flag{
		Name:        "labels",
		Description: "the task's labels",
	}}

	// --priority
	flagTasksUpdatePriority = &flags.Int{Flag: &flags.Flag{
		Name:        "priority",
		Description: "task priority 1 to 4",
	}}

	// --due-string
	flagTasksUpdateDueString = &flags.String{Flag: &flags.Flag{
		Name:        "due-string",
		Description: "human defined task due tate",
	}}

	// --due-date
	flagTasksUpdateDueDate = &flags.String{Flag: &flags.Flag{
		Name:        "due-date",
		Description: "specific date in YYYY-MM-DD format relative to user's timezone",
	}}

	// --due-datetime
	flagTasksUpdateDueDatetime = &flags.String{Flag: &flags.Flag{
		Name:        "due-datetime",
		Description: "specific date and time in RFC3339 format in UTC",
	}}

	// --due-lang
	flagTasksUpdateDueLang = &flags.String{Flag: &flags.Flag{
		Name:        "due-lang",
		Description: "2-letter code specifying language in case `due_string` is not written in English",
	}}

	// --assignee-id
	flagTasksUpdateAssigneeID = &flags.String{Flag: &flags.Flag{
		Name:        "assignee-id",
		Description: "the responsible user id",
	}}
)
