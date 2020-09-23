package endpointtests

import "github.com/zioufang/mltrackapi/pkg/api/model"

func resetTables() {
	server.DB.Migrator().DropTable(&model.Project{})
	server.DB.Migrator().DropTable(&model.Model{})
	server.DB.Migrator().DropTable(&model.ModelRun{})
	server.DB.Migrator().DropTable(&model.RunNumAttr{})
	server.DB.Migrator().DropTable(&model.RunTag{})
	server.DB.AutoMigrate(
		&model.Project{},
		&model.Model{},
		&model.ModelRun{},
		&model.RunNumAttr{},
		&model.RunTag{},
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

func seedRunTagTable() []model.RunTag {
	modelRuns := seedModelRunTable()
	tags := []model.RunTag{
		{
			ModelRunID: modelRuns[0].ID,
			Key:        "git_hash",
			Value:      "dc823jfc",
		},
		{
			ModelRunID: modelRuns[0].ID,
			Key:        "input_data",
			Value:      "s3://input-bucket/data",
		},
		{
			ModelRunID: modelRuns[1].ID,
			Key:        "git_hash",
			Value:      "i49fm34y",
		},
		{
			ModelRunID: modelRuns[1].ID,
			Key:        "model_path",
			Value:      "s3://model-bucket/here",
		},
		{
			ModelRunID: modelRuns[1].ID,
			Key:        "modeler",
			Value:      "awesome guy",
		},
	}

	for i := range tags {
		server.DB.Create(&tags[i])
	}

	return tags
}

func getFloatPointer(val float32) *float32 {
	return &val
}
