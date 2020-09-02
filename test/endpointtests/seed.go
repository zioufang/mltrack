package endpointtests

import "github.com/zioufang/mltrackapi/pkg/api/model"

func resetTables() {
	server.DB.DropTableIfExists(&model.Project{})
	server.DB.DropTableIfExists(&model.Model{})
	server.DB.AutoMigrate(&model.Project{}, &model.Model{})
}

func seedProjectTable() []model.Project {
	resetTables()
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
	for i := range projects {
		server.DB.Create(&projects[i])

	}
	return projects
}

func seedModelTable() []model.Model {
	projects := seedProjectTable()
	models := []model.Model{
		{

			Name:        "kramer",
			ProjectID:   projects[0].ID,
			Status:      "experiment",
			Description: "this is model kramer",
		},
		{
			Name:        "sherpa",
			ProjectID:   projects[0].ID,
			Status:      "production",
			Description: "this is model sherpa",
		},
		{
			Name:        "owen",
			ProjectID:   projects[1].ID,
			Status:      "experiment",
			Description: "this is model owen",
		},
	}
	for i := range models {
		server.DB.Create(&models[i])
	}
	return models
}
