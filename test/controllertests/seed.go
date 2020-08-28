package controllertests

import "github.com/zioufang/mltrackapi/pkg/api/model"

func clearProjectTable() {
	server.DB.Exec("DELETE FROM projects;")
}

func seedProjectTable() []model.Project {
	clearProjectTable()
	projects := []model.Project{
		{
			Name:        "daoud",
			Description: "this is project daoud",
		},
		{
			Name:        "estobar",
			Description: "this is project estobar",
		},
	}
	server.DB.Create(&projects)
	return projects
}
