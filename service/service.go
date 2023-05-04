package service

import (
	"sync"

	"github.com/brunoOchoa.com/api-REST-FULL/domain"
	"github.com/brunoOchoa.com/api-REST-FULL/repository"
	"github.com/brunoOchoa.com/api-REST-FULL/requests"
)

var once sync.Once

type clienteService struct {
	clienteRepository repository.ClienteRepository
}

type ClienteService interface {
	GetAllClientes() ([]domain.Cliente, error)
	GetCliente(string) (domain.Cliente, error)
	CreateCliente(requests.ClienteCreateRequest) (domain.Cliente, error)
	UpdateCliente(string, requests.ClienteUpdateRequest) error
	DeleteCliente(string) error
}

var instance *clienteService

func NewClienteService(r repository.ClienteRepository) ClienteService {
	once.Do(func() {
		instance = &clienteService{
			clienteRepository: r,
		}
	})

	return instance
}

func (*clienteService) GetAllClientes() ([]domain.Cliente, error) {
	return instance.clienteRepository.GetAllClientes()
}

func (*clienteService) GetCliente(id string) (domain.Cliente, error) {
	return instance.clienteRepository.GetCliente(id)
}

func (*clienteService) CreateCliente(request requests.ClienteCreateRequest) (domain.Cliente, error) {

	return instance.clienteRepository.CreateCliente(request)
}

func (*clienteService) UpdateCliente(id string, request requests.ClienteUpdateRequest) error {
	err := instance.clienteRepository.UpdateCliente(id, request)

	if err != nil {
		return err
	}

	return nil
}
func (*clienteService) DeleteCliente(string) error {
	return nil
}
