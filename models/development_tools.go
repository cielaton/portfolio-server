package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Style struct {
	Span      int
	TextColor string
}

type Icons struct {
	Path string
	Name []string
}

// the development tools data model
type developmentTools struct {
	Id          int
	Field       string
	Description string
	Style       Style
	Icons       Icons
}

func QueryDevelopmentTools(database *pgx.Conn) []developmentTools {
	// Query command
	query, err := database.Query(context.Background(), "SELECT tools.id, tools.field, tools.descriptions, tools.text_color, tools.span, icons.path, icons.icon_names FROM favourite_tools as tools INNER JOIN favourite_tools_icons as icons on tools.field = icons.field")
	if err != nil {
		fmt.Printf("[Database] Unable to query development tools: %v\n", err)
	}

	// Iterate through the query results and assign it into returned value
	var returnedLiteral []developmentTools
	for query.Next() {
		// Holding query values for each field
		var field, description, textColor, path string
		var iconNames []string
		var id, span int

		err := query.Scan(&id, &field, &description, &textColor, &span, &path, &iconNames)
		if err != nil {
			fmt.Printf("[Database] Scan error: %v\n", err)
			return nil
		}

		returnedLiteral = append(returnedLiteral, developmentTools{
			id,
			field,
			description,
			Style{span, textColor},
			Icons{path, iconNames},
		})
	}

	// Close query result on function return
	defer query.Close()

	return returnedLiteral
}
