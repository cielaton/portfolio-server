package models

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type Style struct {
	Span      int    `json:"span"`
	TextColor string `json:"textColor"`
}

type Icons struct {
	Path string   `json:"path"`
	Name []string `json:"name"`
}

// the development tools data model
type developmentTools struct {
	Id          int    `json:"id"`
	Field       string `json:"field"`
	Description string `json:"description"`
	Style       Style  `json:"style"`
	Icons       Icons  `json:"icons"`
}

func QueryDevelopmentTools(database *pgx.Conn) []developmentTools {
	// Query command
	query, err := database.Query(context.Background(), "SELECT tools.id, tools.field, tools.descriptions, tools.text_color, tools.span, icons.path, icons.icon_names FROM favourite_tools as tools INNER JOIN favourite_tools_icons as icons on tools.field = icons.field")
	if err != nil {
		log.Err(err).Msg("[Database] Error querying development tools: %v\n")
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
			log.Err(err).Msg("[Database] Scan error: %v\n")
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
