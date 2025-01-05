package main

import (
	ctrl "medicine-app/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// cfg := ctrl.Config{
	// 	Platform: platform,
	// 	DB:       dbQueries,
	// }

	// mux := http.NewServeMux()
	// mux.HandleFunc("/medicines", cfg.CreateMedicineHandler)
	// mux.HandleFunc("/medicines/:medID", cfg.DeleteMedicine)
	// port := "8080"
	// srv := http.Server{
	// 	Handler: mux,
	// 	Addr: ":" + port,
	// }
	// log.Printf("Serving on port: %v\n", port)
	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Printf("%v\n", err.Error())
	// }

	medCtrl := ctrl.MedicineController{}

	router := gin.Default()

	router.POST("/medicines", medCtrl.CreateMedicineHandler)

	router.Run("8080")
	// router.POST("/medicines", cfg.CreateMedicineHandler)
}
