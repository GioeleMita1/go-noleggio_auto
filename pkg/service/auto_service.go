package service

import (
	"context"
	"errors"
	"noleggio_auto/pkg/dto"
	"noleggio_auto/pkg/factory"
	"noleggio_auto/pkg/repository"
	"noleggio_auto/pkg/model"
)


type autoService struct {
	repo repository.AutoRepository
}

type AutoService interface {
	CreaNuovaAuto(ctx context.Context, req dto.CreateAutoRequest) (*model.Auto, error)
	OttieniAuto(ctx context.Context, id string) (*model.Auto, error)
	ListaAuto(ctx context.Context) ([]model.Auto, error)
	AggiornaAuto(ctx context.Context, id string, req dto.UpdateAutoRequest) error
	RimuoviAuto(ctx context.Context, id string) error
}


func NewAutoService(repo repository.AutoRepository) AutoService {
	return &autoService{
		repo: repo,
	}
}

//CREATE
func (s *autoService) CreaNuovaAuto(ctx context.Context, req dto.CreateAutoRequest) (*model.Auto, error) {
	auto := factory.ToModel(req)

	err := s.repo.Create(ctx, auto)
	if err != nil {
		return nil, err
	}
	return auto, nil
}

// READ
func (s *autoService) OttieniAuto(ctx context.Context, id string) (*model.Auto, error) {
	if id == "" {
		return nil, errors.New("id richiesto")
	}
	return s.repo.FindByID(ctx, id)
}

//READ ALL  ancora non funzionante 
func (s *autoService) ListaAuto(ctx context.Context) ([]model.Auto, error) {
	return s.repo.FindAll(ctx)
}

//UPDATE
func (s *autoService) AggiornaAuto(ctx context.Context, id string, req dto.UpdateAutoRequest) error {
	
	autoM := &model.Auto{
		Marca:        req.Marca,
		Modello:      req.Modello,
		Disponibile:  req.Disponibile,
	}
	
	return s.repo.Update(ctx, id, autoM)
}

// DELETE
func (s *autoService) RimuoviAuto(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}