package main

import (
	"health-tracker/reports"
	"health-tracker/storage"
)

func main() {
	const filename = "user_data.json"

	firstUser, isExisting := storage.LoadOrRegister(filename)

	if isExisting {
			reports.UpdateWeight(&firstUser)
	} else {
		firstUser = reports.RegisterNewUser()
	}

	reports.ShowHealthReport(firstUser)
	reports.ShowWeightChange(firstUser)
	reports.ShowingAcitve(firstUser)
	reports.ShowTotalProgress(firstUser)

	storage.Save(firstUser, filename)

}
