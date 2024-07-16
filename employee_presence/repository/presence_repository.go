package repository

import (
	"employee_presence/model"
	"fmt"
	"time"
)

type PresenceRepository struct {
	data   []model.Employee
	lastId int
}

func NewPresenceRepository() *PresenceRepository {
	return &PresenceRepository{}
}

func (p *PresenceRepository) Save(data model.Employee) {

	// Set id from lastId
	p.lastId++
	data.ID = p.lastId
	data.Presence = true
	data.CreatedAt = time.Now().Format(time.RFC822)

	p.data = append(p.data, data)
}

func (p *PresenceRepository) GetAll() []model.Employee {
	return p.data
}

func (p *PresenceRepository) Update(data model.Employee) error {

	for i, v := range p.data {
		if data.ID == v.ID {
			p.data[i] = data
			return nil
		}
	}

	return fmt.Errorf("Employee with id: %d not found", data.ID)
}

func (p *PresenceRepository) Delete(id int) error {

	for i, v := range p.data {
		if id == v.ID {
			p.data = append(p.data[:i], p.data[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Employee with id: %d not found", id)
}

func (p *PresenceRepository) FindById(employee *model.Employee) error {

	for _, v := range p.data {
		if employee.ID == v.ID {
			*employee = v
			return nil
		}
	}

	return fmt.Errorf("Employee with id: %d not found", employee.ID)
}
