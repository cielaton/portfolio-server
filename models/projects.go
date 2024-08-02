package models

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type Project struct {
	Id                 int
	ProjectName        string
	ProjectDescription string
	Technologies       []string
	ImagePath          string
}

func QueryProjects(database *pgx.Conn, projectType string) []Project {
	// Query command with dynamic table
	query, err := database.Query(context.Background(), "SELECT id, project_name, project_description, technologies, image_path FROM "+projectType)
	if err != nil {
		log.Err(err).Msg("[Database] Error querying personal projects: %v\n")
	}

	// Iterate through the query results and assign it into returned value
	var returnedProjects []Project
	for query.Next() {
		// Holding query values for each field
		var (
			id                 int
			projectName        string
			projectDescription string
			technologies       []string
			imagePath          string
		)
		err := query.Scan(&id, &projectName, &projectDescription, &technologies, &imagePath)
		if err != nil {
			log.Err(err).Msg("[Database] Error scanning personal projects: %v\n")
			return nil
		}

		returnedProjects = append(returnedProjects, Project{
			id, projectName, projectDescription, technologies, imagePath,
		})
	}
	// Close query result on function return
	defer query.Close()

	return returnedProjects
}
