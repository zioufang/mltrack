package endpointtests

import "github.com/zioufang/mltrackapi/pkg/api/model"

func resetTables() {
	server.DB.DropTableIfExists(&model.Project{})
	server.DB.DropTableIfExists(&model.Model{})
	server.DB.DropTableIfExists(&model.ModelRun{})
	server.DB.DropTableIfExists(&model.RunNumAttr{})
	server.DB.AutoMigrate(
		&model.Project{},
		&model.Model{},
		&model.ModelRun{},
		&model.RunNumAttr{},
	)
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

func seedModelRunTable() []model.ModelRun {
	models := seedModelTable()
	runs := []model.ModelRun{
		{

			Name:    "perceval",
			ModelID: models[0].ID,
		},
		{
			Name:    "burroughs",
			ModelID: models[0].ID,
		},
		{
			Name:    "paramythis",
			ModelID: models[1].ID,
		},
	}
	for i := range runs {
		server.DB.Create(&runs[i])
	}
	return runs
}

func seedRunNumAttrTable() []model.RunNumAttr {
	modelRuns := seedModelRunTable()
	attrs := []model.RunNumAttr{
		{
			ModelRunID: modelRuns[0].ID,
			Name:       "metric_1",
			Category:   "metric",
			Value:      getFloatPointer(0.1),
		},
		{
			ModelRunID: modelRuns[0].ID,
			Name:       "metric_2",
			Category:   "metric",
			Value:      getFloatPointer(0.0),
		},
		{
			ModelRunID: modelRuns[0].ID,
			Name:       "param_2",
			Category:   "param",
			Value:      getFloatPointer(0.0),
		},
		{
			ModelRunID: modelRuns[1].ID,
			Name:       "metric_1",
			Category:   "metric",
			Value:      getFloatPointer(0.1),
		},
		{
			ModelRunID: modelRuns[1].ID,
			Name:       "param_1",
			Category:   "param",
			Value:      getFloatPointer(1.0),
		},
		{
			ModelRunID: modelRuns[1].ID,
			Name:       "param_2",
			Category:   "param",
			Value:      getFloatPointer(0.0001),
		},
	}

	for i := range attrs {
		server.DB.Create(&attrs[i])
	}

	return attrs
}

func getFloatPointer(val float32) *float32 {
	return &val
}
