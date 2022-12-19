package innovahandler

import "demo-fiber-mysql-gorm/api/domain"

type innovaHandler struct {
	InnovaService domain.InnovaService
}

func NewInnovaHandler(innovaService domain.InnovaService) domain.InnovaHandler {
	return &innovaHandler{InnovaService: innovaService}
}
