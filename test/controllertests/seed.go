package controllertests

import "github.com/zioufang/mltrackapi/pkg/api/model"

func resetProjectTable() {
	server.DB.DropTableIfExists(&model.Project{})
	server.DB.AutoMigrate(&model.Project{})
}

func seedProjectTable() []model.Project {
	resetProjectTable()
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
