package controller

import (
	"errors"
	"log"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
)

func (c *Conn) CreateBranch(req *models.CreateBranch) (*models.Branch, error) {

	log.Println(config.Info, "Create Branch resp >>>>> ", req)
	resp, err := c.storage.Branch().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDBranch(req *models.BranchPrimaryKey) (*models.Branch, error) {

	log.Println(config.Info, "GetByID Branch resp >>>>> ", req)
	resp, err := c.storage.Branch().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListBranch(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error) {

	log.Println(config.Info, "GetList Branch resp >>>>> ", req)
	resp, err := c.storage.Branch().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateBranch(req *models.UpdateBranch) (*models.Branch, error) {

	log.Println(config.Info, "Update Branch resp >>>>> ", req)
	rowsAffected, err := c.storage.Branch().Update(req)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}

	resp, err := c.storage.Branch().GetByID(&models.BranchPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteBranch(req *models.BranchPrimaryKey) error {

	log.Println(config.Info, "Delete Branch resp >>>>> ", req)
	err := c.storage.Branch().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
