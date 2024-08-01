package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func QueryDevelopmentTools(database *pgx.Conn) {
	developmentTools, err := database.Query(context.Background(), "SELECT tools.field, tools.descriptions, tools.text_color, tools.span, icons.path, icons.icon_names FROM favourite_tools as tools INNER JOIN favourite_tools_icons as icons on tools.field = icons.field", nil)
	if err != nil {
		fmt.Printf("[Database] Unable to query development tools: %v\n", err)
	}
	fmt.Println(developmentTools)
}
