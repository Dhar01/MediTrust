package main

import (
	_ "github.com/lib/pq"
)

func main() {

	// medApp := api.MedicineApp{
	// 	DB:     dbQueries,
	// 	Router: router,
	// }

	// router := gin.Default()

	// // for collection
	// router.POST("/medicines", medApp.CreateMedicine)
	// router.GET("/medicines", medApp.GetMedicine)
	// // for single item
	// router.DELETE("/medicines/:medicineID", medApp.DeleteMedicine)
	// router.PUT("/medicines/:medicineID", medApp.UpdateMedicine)
	// router.GET("/medicines/:medicineID", medApp.GetMedicine)

	// router.Run(":8080")

	// cfg := config.LoadConfig()
	// medApp := app.NewMedicineApp(cfg)

	// mux := http.NewServeMux()

	// mux.HandleFunc("/medicines", )

}
