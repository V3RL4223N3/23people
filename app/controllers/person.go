package controllers

import (
	"errors"
	"github.com/V3RL4223N3/23people/app/models"
	"github.com/revel/revel"
)

type PersonController struct {
	*revel.Controller
}

func (c PersonController) Index() revel.Result {
	var (
		persons []models.Person
		err     error
	)
	persons, err = models.GetPersons()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(persons)
}

func (c PersonController) Show(id string) revel.Result {
	var (
		person models.Person
		err    error
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid person id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	//personID, err = convertToObjectIdHex(id)
	//if err != nil {
	//	errResp := buildErrResponse(errors.New("Invalid person id2 format"), "400")
	//	c.Response.Status = 400
	//	fmt.Println(err)
	//	return c.RenderJSON(errResp)
	//}

	person, err = models.GetPerson(id)
	if err != nil {
		errResp := buildErrResponse(err, "404")
		c.Response.Status = 404
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(person)
}

func (c PersonController) Create() revel.Result {
	var (
		person models.Person
		err    error
	)

	err = c.Params.BindJSON(&person)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	person, err = models.AddPerson(person)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(person)
}

func (c PersonController) Update(id string) revel.Result {
	var (
		person models.Person
		err    error
	)
	err = c.Params.BindJSON(&person)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	if id != person.NationalId {
		err = errors.New("JSON payload national_id is different from URL")
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)

	}

	err = person.UpdatePerson(id)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	return c.RenderJSON(person)
}

func (c PersonController) Delete(id string) revel.Result {
	var (
		err    error
		person models.Person
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid person id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	person, err = models.GetPerson(id)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = person.DeletePerson()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}
