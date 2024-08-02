package models

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type Project struct {
	Id                 int      `json:"id"`
	ProjectName        string   `json:"projectName"`
	ProjectDescription string   `json:"projectDescription"`
	Link               string   `json:"link"`
	Technologies       []string `json:"technologies"`
	ImagePath          string   `json:"imagePath"`
}

func QueryProjects(database *pgx.Conn, projectType string) []Project {
	// Query command with dynamic table
	query, err := database.Query(context.Background(), "SELECT id, project_name, project_description, link, technologies, image_path FROM "+projectType+" ORDER BY id")
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
			link               string
			technologies       []string
			imagePath          string
		)
		err := query.Scan(&id, &projectName, &projectDescription, &link, &technologies, &imagePath)
		if err != nil {
			log.Err(err).Msg("[Database] Error scanning personal projects: %v\n")
			return nil
		}

		returnedProjects = append(returnedProjects, Project{
			id, projectName, projectDescription, link, technologies, imagePath,
		})
	}
	// Close query result on function return
	defer query.Close()

	return returnedProjects
}
