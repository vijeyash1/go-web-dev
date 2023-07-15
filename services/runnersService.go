package services

import (
	"net/http"
	"runners-postgresql/models"
	"strconv"
	"time"
)

type RunnersService struct {
	runnersRepository *repositories.RunnersRepository
	resultsRepository *repository.ResultsRepository
}

func NewRunnersService(runnersRepository *repositories.RunnersRepository, resultsRepository *repository.ResultsRepository) RunnersService {
	return RunnersService{
		runnersRepository: runnersRepository,
		resultsRepository: resultsRepository,
	}
}

func (rs RunnersService) CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseError) {
	responseErr := validateRunner(runner)
	if responseErr != nil {
		return nil, responseErr
	}
	return rs.runnersRepository.CreateRunner(runner)
}
func (rs RunnersService) UpdateRunner(runner *models.Runner) *models.ResponseError {
	responseErr := validateRunnerId(runner.ID)
	if responseErr != nil {
		return responseErr
	}
	responseErr = validateRunner(runner)
	if responseErr != nil {
		return responseErr
	}
	return rs.runnersRepository.UpdateRunner(runner)
}
func (rs RunnersService) DeleteRunner(runnerId string) *models.ResponseError {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil {
		return responseErr
	}
	return rs.runnersRepository.DeleteRunner(runnerId)
}
func (rs RunnersService) GetRunner(runnerId string) (*models.Runner, *models.ResponseError) {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}
	runner, responseErr := rs.runnersRepository.GetRunner(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}
	results, responseErr := rs.resultsRepository.GetAllRunnersResults(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}
	runner.Result = results
	return runner, nil
}
func (rs RunnersService) GetRunnersBatch(country, year string) (*[]models.Runner, *models.ResponseError) {
	if country != "" && year != "" {
		return nil, &models.ResponseError{
			Message: "Only one parameter can be passed",
			Status:  http.StatusBadRequest,
		}
	}
	if country != "" {
		return rs.runnersRepository.GetRunnersByCountry(country)
	}
	if year != "" {
		intYear, err := strconv.Atoi(year)
		if err != nil {
			return nil, &models.ResponseError{
				Message: "Invalid Year",
				Status:  http.StatusBadRequest,
			}
		}
		currentYear := time.Now().Year()
		if intYear < 0 || intYear > currentYear {
			return nil, &models.ResponseError{
				Message: "Invalid Year",
				Status:  http.StatusBadRequest,
			}
		}
		return rs.runnersRepository.GetRunnerByYear(intYear)
	}
	return rs.runnersRepository.GetAllRunners()
}
func validateRunner(runner *models.Runner) *models.ResponseError {
	if runner.FirstName == "" {
		return &models.ResponseError{
			Message: "Invalid First Name",
			Status:  http.StatusBadRequest,
		}
	}
	if runner.LastName == "" {
		return &models.ResponseError{
			Message: "Invalid Last Name",
			Status:  http.StatusBadRequest,
		}
	}
	if runner.Age <= 16 || runner.Age > 125 {
		return &models.ResponseError{
			Message: "Invalid Age",
			Status:  http.StatusBadRequest,
		}
	}
	if runner.Country == "" {
		return &models.ResponseError{
			Message: "Invalid Country",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func validateRunnerId(runnerId string) *models.ResponseError {
	if runnerId == "" {
		return &models.ResponseError{
			Message: "Invalid runner ID",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
