package storage

import (
	"context"
	gen "recruitment/genprotos"
)

type StorageInterface interface {
	Company() CompanyInterface
	Job() JobInterface
	Vacancy() VacancyInterface
}


type CompanyInterface interface {
	Create(ctx context.Context, req *gen.CreateCompany) (*gen.Void, error)
	GetByID(ctx context.Context, req *gen.Byid) (*gen.GetByID, error)
	Update(ctx context.Context, req *gen.CompanyUpdate) (*gen.Void, error)
	Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error)
	GetAll(ctx context.Context, req *gen.GetAllCompanyRequest) (*gen.GetAllCompanyResponse, error)
}


type JobInterface interface {
	Create(ctx context.Context, req *gen.JobApplicationCreate) (*gen.Void, error)
	GetByID(ctx context.Context, req *gen.Byid) (*gen.JobApplitcationGetRes, error)
	Update(ctx context.Context, req *gen.JobApplicationUpdate) (*gen.Void, error)
	Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error)
	GetAll(ctx context.Context, req *gen.JobApplicationGetAll) (*gen.JobApplicationGetAllRes, error)
}


type VacancyInterface interface {
	Create(ctx context.Context, req *gen.VacancyCreate) (*gen.Void, error)
	GetByID(ctx context.Context, req *gen.Byid) (*gen.VacancyGetResUpdateReq, error)
	Update(ctx context.Context, req *gen.VacancyGetResUpdateReq) (*gen.Void, error)
	Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error)
	GetAll(ctx context.Context, req *gen.VacancyGetAllReq) (*gen.VacancyGetAllRes, error)
}