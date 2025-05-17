package product

import "github.com/labstack/echo/v4"

func ProductRoutes(router echo.Group) {
	repo := newMedicineRepo(nil)
	svc := newMedicineService(repo)
	_ = newMedicineHandler(svc)
}
