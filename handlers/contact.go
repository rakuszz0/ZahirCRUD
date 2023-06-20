package handlers

import (
	contactsdto "CRUD/dto/contact"
	dto "CRUD/dto/result"
	"CRUD/models"
	"CRUD/repositories"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handler struct {
	ContactRepository repositories.ContactRepository
}

func HandlerContact(ContactRepository repositories.ContactRepository) *handler {
	return &handler{ContactRepository}
}

func (h *handler) FindContacts(c echo.Context) error {
	contacts, err := h.ContactRepository.FindContacts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: contacts})
}

func (h *handler) GetContact(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	contact, err := h.ContactRepository.GetContact(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(contact)})
}

func (h *handler) CreateContact(c echo.Context) error {
	request := new(contactsdto.CreateContactRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	contact := models.Contact{
		ID:        0,
		Name:      request.Name,
		Gender:    request.Gender,
		Phone:     request.Phone,
		Email:     request.Email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	data, err := h.ContactRepository.CreateContact(contact)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) GetContactList(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	filter := c.QueryParam("filter")
	page := c.QueryParam("page")
	pageSize := c.QueryParam("page_size")

	pageToInt, _ := strconv.Atoi(page)
	pageSizeToInt, _ := strconv.Atoi(pageSize)

	contacts, err := h.ContactRepository.FindContacts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	filteredContacts := filterContactsByName(contacts, filter)

	if sortBy != "" {
		sortContacts(filteredContacts, sortBy)
	}

	paginatedContacts := paginateContacts(filteredContacts, pageToInt, pageSizeToInt)

	// Return the paginated contact list data
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: paginatedContacts})
}

// Filter contacts by name
func filterContactsByName(contacts []models.Contact, filter string) []models.Contact {
	var filteredContacts []models.Contact
	for _, contact := range contacts {
		if strings.Contains(contact.Name, filter) {
			filteredContacts = append(filteredContacts, contact)
		}
	}
	return filteredContacts
}

// Sort contacts based on the given field
func sortContacts(contacts []models.Contact, sortBy string) {
	switch sortBy {
	case "name":
		sort.SliceStable(contacts, func(i, j int) bool {
			return contacts[i].Name < contacts[j].Name
		})
	case "email":
		sort.SliceStable(contacts, func(i, j int) bool {
			return contacts[i].Email < contacts[j].Email
		})
		// Add more cases for other sortable fields if needed
	}
}

// Paginate contacts based on the page and page size
func paginateContacts(contacts []models.Contact, page, pageSize int) []models.Contact {
	startIndex := (page - 1) * pageSize
	if startIndex >= len(contacts) {
		return []models.Contact{}
	}

	endIndex := startIndex + pageSize
	if endIndex > len(contacts) {
		endIndex = len(contacts)
	}

	return contacts[startIndex:endIndex]
}

func (h *handler) UpdateContact(c echo.Context) error {
	request := new(contactsdto.UpdateContactRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	contact, err := h.ContactRepository.GetContact(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		contact.Name = request.Name
	}

	if request.Phone != "" {
		contact.Phone = request.Phone
	}

	if request.Gender != "" {
		contact.Gender = request.Gender
	}

	if request.Email != "" {
		contact.Email = request.Email
	}

	data, err := h.ContactRepository.UpdateContact(contact)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) DeleteContact(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.ContactRepository.GetContact(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ContactRepository.DeleteContact(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func convertResponse(u models.Contact) contactsdto.ContactResponse {
	return contactsdto.ContactResponse{
		ID:     u.ID,
		Name:   u.Name,
		Gender: u.Gender,
		Phone:  u.Phone,
		Email:  u.Email,
	}
}

// func GetContactList(c echo.Context) error {
// 	sortBy := c.QueryParam("sort_by")
// 	filter := c.QueryParam("filter")
// 	page := c.QueryParam("page")
// 	pageSize := c.QueryParam("page_size")

// 	pageToInt, _ := strconv.Atoi(page)
// 	pageSizeToInt, _ := strconv.Atoi(pageSize)

// 	filteredContacts := dataJson.ContactList
// 	if filter != "" {
// 		filteredContacts = filterContactsByName(dataJson.ContactList, filter)
// 	}

// 	if sortBy != "" {
// 		sortContacts(filteredContacts, sortBy)
// 	}

// 	paginateContacts := paginateContacts(filteredContacts, pageToInt, pageSizeToInt)

// 	// return all contact list data
// 	return c.JSON(http.StatusOK, paginateContacts)
// }
