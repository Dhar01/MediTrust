package api

import "medicine-app/internal/services"

// type API struct {
// 	userAPI UserAPI
// 	authAPI AuthAPI
// 	medAPI  MedicineAPI
// }

type API struct {
	services   *services.Services
	MedService services.MedService
}

func NewAPI(srv *services.Services) API {
	if srv == nil {
		panic("services can't be nil")
	}

	return API{
		services:   srv,
		MedService: srv.MedService,
	}
}

// type AuthAPI struct {
// 	service *services.Services
// }

// type MedicineAPI struct {
// 	service *services.Services
// }

// func NewAPI(srv *services.Services) API {
// 	if srv == nil {
// 		panic("service can't be nil")
// 	}
// 	return API{
// 		userAPI: newUserAPI(srv),
// 		authAPI: newAuthAPI(srv),
// 		medAPI:  newMedicineAPI(srv),
// 	}
// }

// func newAuthAPI(srv *services.Services) AuthAPI {
// 	return AuthAPI{
// 		service: srv,
// 	}
// }

// func newMedicineAPI(srv *services.Services) MedicineAPI {
// 	return MedicineAPI{
// 		service: srv,
// 	}
// }
