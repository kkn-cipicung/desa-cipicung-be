package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"cipicung.id/be/utils"
)

type Service interface {
	Create(ctx context.Context, payload AddProfilePayload) error
	Detail(ctx context.Context) (*ProfileOutput, error)
	FindRegionBoundary(ctx context.Context) (*ProfileRegionBoundaryResponse, error)
	FindVisionMission(ctx context.Context) (*ProfileVisionMissionOutput, error)
	FindGovernmentStructure(ctx context.Context) ([]GovernmentStructureResponse, error)
	FindResourcePotential(ctx context.Context) (ResourcePotentialResponse, error)
	Delete(ctx context.Context, payload ProfilePayload) error
	CreateOfficial(ctx context.Context, payload AddOfficialPayload) error
	ListOfficials(ctx context.Context, payload ListOfficialPayload) ([]OfficialResponse, error)
	FindOfficialByID(ctx context.Context, payload OfficialPayload) (*OfficialResponse, error)
	UpdateOfficial(ctx context.Context, payload EditOfficialPayload) error
	DeleteOfficial(ctx context.Context, payload OfficialPayload) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, payload AddProfilePayload) error {
	normalizeProfilePayload(&payload)
	if err := validateProfilePayload(payload); err != nil {
		return err
	}
	first, err := s.repository.FindFirst(ctx)
	if err == nil && first != nil && first.ID != 0 {
		payload.ID = first.ID
		return s.repository.Update(ctx, EditProfilePayload{AddProfilePayload: payload})
	}
	if payload.ID != 0 {
		return s.repository.Update(ctx, EditProfilePayload{AddProfilePayload: payload})
	}
	return s.repository.Create(ctx, payload)
}

func (s *service) Detail(ctx context.Context) (*ProfileOutput, error) {
	item, err := s.repository.Detail(ctx)
	if err != nil {
		return nil, err
	}
	output := mapProfileOutput(*item)
	if item.ID != 0 {
		headmen, err := s.repository.FindHeadmen(ctx, item.ID)
		if err == nil {
			output.Headmen = headmen
		}
	}
	return &output, nil
}

func (s *service) FindRegionBoundary(ctx context.Context) (*ProfileRegionBoundaryResponse, error) {
	return s.repository.FindRegionBoundary(ctx)
}

func (s *service) FindVisionMission(ctx context.Context) (*ProfileVisionMissionOutput, error) {
	item, err := s.repository.FindVisionMission(ctx)
	if err != nil {
		return nil, err
	}
	return &ProfileVisionMissionOutput{
		Vision:  item.Vision,
		Mission: []string(item.Mission),
	}, nil
}

func (s *service) FindGovernmentStructure(ctx context.Context) ([]GovernmentStructureResponse, error) {
	return s.repository.FindGovernmentStructure(ctx)
}

func (s *service) FindResourcePotential(ctx context.Context) (ResourcePotentialResponse, error) {
	return s.repository.FindResourcePotential(ctx)
}

func (s *service) Delete(ctx context.Context, payload ProfilePayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: profile id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.Delete(ctx, payload)
}

func normalizeProfilePayload(payload *AddProfilePayload) {
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Province = strings.TrimSpace(payload.Province)
	payload.Regency = strings.TrimSpace(payload.Regency)
	payload.District = strings.TrimSpace(payload.District)
	payload.PostalCode = strings.TrimSpace(payload.PostalCode)
	payload.Address = strings.TrimSpace(payload.Address)
	payload.Phone = strings.TrimSpace(payload.Phone)
	payload.Email = strings.TrimSpace(payload.Email)
	payload.Website = strings.TrimSpace(payload.Website)
	payload.Vision = strings.TrimSpace(payload.Vision)
	for index, mission := range payload.Mission {
		payload.Mission[index] = strings.TrimSpace(mission)
	}
	payload.History = strings.TrimSpace(payload.History)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.Region = strings.TrimSpace(payload.Region)
	payload.NorthBorder = strings.TrimSpace(payload.NorthBorder)
	payload.EastBorder = strings.TrimSpace(payload.EastBorder)
	payload.SouthBorder = strings.TrimSpace(payload.SouthBorder)
	payload.WestBorder = strings.TrimSpace(payload.WestBorder)
	payload.Area = strings.TrimSpace(payload.Area)
	if payload.RTHamletOne+payload.RTHamletTwo > 0 {
		payload.TotalRT = payload.RTHamletOne + payload.RTHamletTwo
	}
	if payload.RWHamletOne+payload.RWHamletTwo > 0 {
		payload.TotalRW = payload.RWHamletOne + payload.RWHamletTwo
	}
	payload.Population = calculateProfilePopulation(payload.HamletOne, payload.HamletTwo)
	if payload.Headman != nil {
		normalizeHeadmanPayload(payload.Headman)
	}
	for index := range payload.Headmen {
		normalizeHeadmanPayload(&payload.Headmen[index])
	}
	for index := range payload.Officials {
		normalizeGovernmentOfficial(&payload.Officials[index])
	}
	if payload.ResourcePotential != nil {
		payload.ResourcePotential.Title = strings.TrimSpace(payload.ResourcePotential.Title)
		payload.ResourcePotential.Detail = strings.TrimSpace(payload.ResourcePotential.Detail)
		payload.ResourcePotential.Description = strings.TrimSpace(payload.ResourcePotential.Description)
	}
}

func validateProfilePayload(payload AddProfilePayload) error {
	if payload.Name == "" {
		return fmt.Errorf("%w: name is required", utils.ErrInvalidPayload)
	}
	if payload.Province == "" {
		return fmt.Errorf("%w: province is required", utils.ErrInvalidPayload)
	}
	if payload.Regency == "" {
		return fmt.Errorf("%w: regency is required", utils.ErrInvalidPayload)
	}
	if payload.District == "" {
		return fmt.Errorf("%w: district is required", utils.ErrInvalidPayload)
	}
	if payload.Address == "" {
		return fmt.Errorf("%w: address is required", utils.ErrInvalidPayload)
	}
	if payload.HamletOne < 0 {
		return fmt.Errorf("%w: hamlet_one cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.HamletTwo < 0 {
		return fmt.Errorf("%w: hamlet_two cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.TotalFamily < 0 {
		return fmt.Errorf("%w: total_family cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.TotalRT < 0 {
		return fmt.Errorf("%w: total_rt cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.TotalRW < 0 {
		return fmt.Errorf("%w: total_rw cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.RTHamletOne < 0 {
		return fmt.Errorf("%w: rt_hamlet_one cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.RTHamletTwo < 0 {
		return fmt.Errorf("%w: rt_hamlet_two cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.RWHamletOne < 0 {
		return fmt.Errorf("%w: rw_hamlet_one cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.RWHamletTwo < 0 {
		return fmt.Errorf("%w: rw_hamlet_two cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.TotalMale < 0 {
		return fmt.Errorf("%w: total_male cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.TotalFemale < 0 {
		return fmt.Errorf("%w: total_female cannot be negative", utils.ErrInvalidPayload)
	}
	if payload.Headman != nil {
		if err := validateHeadmanPayload(*payload.Headman); err != nil {
			return err
		}
	}
	for _, headman := range payload.Headmen {
		if err := validateHeadmanPayload(headman); err != nil {
			return err
		}
	}
	if err := validateHeadmanPeriods(payload.Headmen); err != nil {
		return err
	}
	for _, official := range payload.Officials {
		if official.Name == "" || official.Position == "" {
			return fmt.Errorf("%w: official name and position are required", utils.ErrInvalidPayload)
		}
		if utils.GenerateSlug(official.Position) == headmanPosition || strings.EqualFold(official.Position, "kepala desa") || strings.EqualFold(official.Position, "kepala-desa") {
			return fmt.Errorf("%w: use headmen for officials with kepala-desa position", utils.ErrInvalidPayload)
		}
	}
	if payload.ResourcePotential != nil && (payload.ResourcePotential.Title == "" || payload.ResourcePotential.Detail == "") {
		return fmt.Errorf("%w: resource_potential title and detail are required", utils.ErrInvalidPayload)
	}
	return nil
}

func calculateProfilePopulation(hamletOne, hamletTwo int64) string {
	return strconv.FormatInt(hamletOne+hamletTwo, 10)
}

func normalizeGovernmentOfficial(payload *GovernmentOfficialInput) {
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Position = utils.GenerateSlug(strings.TrimSpace(payload.Position))
}

func validateOptionalOfficialDates(startDate, finishDate *string) error {
	if startDate == nil && finishDate == nil {
		return nil
	}
	if startDate == nil {
		return fmt.Errorf("%w: official start_date is required when finish_date is provided", utils.ErrInvalidPayload)
	}
	start, err := time.Parse("2006-01-02", *startDate)
	if err != nil {
		return fmt.Errorf("%w: official start_date must use YYYY-MM-DD format", utils.ErrInvalidPayload)
	}
	if finishDate == nil {
		return nil
	}
	finish, err := time.Parse("2006-01-02", *finishDate)
	if err != nil {
		return fmt.Errorf("%w: official finish_date must use YYYY-MM-DD format", utils.ErrInvalidPayload)
	}
	if finish.Before(start) {
		return fmt.Errorf("%w: official finish_date cannot be before start_date", utils.ErrInvalidPayload)
	}
	return nil
}

func normalizeHeadmanPayload(payload *ProfileOfficialInput) {
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Position = utils.GenerateSlug(strings.TrimSpace(payload.Position))
	if payload.Position == "" {
		payload.Position = headmanPosition
	}
	payload.Phone = strings.TrimSpace(payload.Phone)
	payload.Email = strings.TrimSpace(payload.Email)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.StartDate = strings.TrimSpace(payload.StartDate)
	if payload.FinishDate != nil {
		finishDate := strings.TrimSpace(*payload.FinishDate)
		if finishDate == "" {
			payload.FinishDate = nil
		} else {
			payload.FinishDate = &finishDate
		}
	}
}

func validateHeadmanPayload(payload ProfileOfficialInput) error {
	if payload.Name == "" {
		return fmt.Errorf("%w: headman name is required", utils.ErrInvalidPayload)
	}
	if payload.Position != headmanPosition {
		return fmt.Errorf("%w: headman position must be kepala-desa", utils.ErrInvalidPayload)
	}
	if payload.StartDate == "" {
		return fmt.Errorf("%w: headman start_date is required", utils.ErrInvalidPayload)
	}
	startDate, err := time.Parse("2006-01-02", payload.StartDate)
	if err != nil {
		return fmt.Errorf("%w: headman start_date must use YYYY-MM-DD format", utils.ErrInvalidPayload)
	}
	if payload.FinishDate != nil && *payload.FinishDate != "" {
		finishDate, err := time.Parse("2006-01-02", *payload.FinishDate)
		if err != nil {
			return fmt.Errorf("%w: headman finish_date must use YYYY-MM-DD format", utils.ErrInvalidPayload)
		}
		if finishDate.Before(startDate) {
			return fmt.Errorf("%w: headman finish_date cannot be before start_date", utils.ErrInvalidPayload)
		}
	}
	return nil
}

func validateHeadmanPeriods(headmen []ProfileOfficialInput) error {
	for currentIndex := 0; currentIndex < len(headmen); currentIndex++ {
		currentStart, _ := time.Parse("2006-01-02", headmen[currentIndex].StartDate)
		currentFinishYear := int(^uint(0) >> 1)
		if headmen[currentIndex].FinishDate != nil {
			currentFinish, _ := time.Parse("2006-01-02", *headmen[currentIndex].FinishDate)
			currentFinishYear = currentFinish.Year()
		}

		for comparedIndex := currentIndex + 1; comparedIndex < len(headmen); comparedIndex++ {
			comparedStart, _ := time.Parse("2006-01-02", headmen[comparedIndex].StartDate)
			comparedFinishYear := int(^uint(0) >> 1)
			if headmen[comparedIndex].FinishDate != nil {
				comparedFinish, _ := time.Parse("2006-01-02", *headmen[comparedIndex].FinishDate)
				comparedFinishYear = comparedFinish.Year()
			}

			periodsOverlap := currentStart.Year() <= comparedFinishYear && comparedStart.Year() <= currentFinishYear
			if periodsOverlap {
				return fmt.Errorf(
					"%w: headman terms for %q and %q cannot contain the same year",
					utils.ErrInvalidPayload,
					headmen[currentIndex].Name,
					headmen[comparedIndex].Name,
				)
			}
		}
	}

	return nil
}

func mapProfileOutput(item ProfileResponse) ProfileOutput {
	output := ProfileOutput{
		ID:                    item.ID,
		Name:                  item.Name,
		Province:              item.Province,
		Regency:               item.Regency,
		District:              item.District,
		PostalCode:            item.PostalCode,
		Address:               item.Address,
		Phone:                 item.Phone,
		Email:                 item.Email,
		Website:               item.Website,
		Latitude:              item.Latitude,
		Longitude:             item.Longitude,
		Vision:                item.Vision,
		Mission:               []string(item.Mission),
		History:               item.History,
		Description:           item.Description,
		Region:                item.Region,
		HamletOne:             item.HamletOne,
		HamletTwo:             item.HamletTwo,
		TotalFamily:           item.TotalFamily,
		TotalRT:               item.TotalRT,
		TotalRW:               item.TotalRW,
		RTHamletOne:           item.RTHamletOne,
		RTHamletTwo:           item.RTHamletTwo,
		RWHamletOne:           item.RWHamletOne,
		RWHamletTwo:           item.RWHamletTwo,
		NorthBorder:           item.NorthBorder,
		EastBorder:            item.EastBorder,
		SouthBorder:           item.SouthBorder,
		WestBorder:            item.WestBorder,
		Area:                  item.Area,
		Population:            item.Population,
		TotalMale:             item.TotalMale,
		TotalFemale:           item.TotalFemale,
		DemographicReligions:  parseDemographicJSON[DemographicSlice](item.DemographicReligionsJSON),
		DemographicReligionRT: parseDemographicJSON[DemographicBar](item.DemographicReligionRTJSON),
		DemographicEducation:  parseDemographicJSON[DemographicGenderBar](item.DemographicEducationJSON),
		DemographicOccupation: parseDemographicJSON[DemographicGenderBar](item.DemographicOccupationJSON),
		DemographicAges:       parseDemographicJSON[DemographicSingleBar](item.DemographicAgesJSON),
		IsActive:              item.IsActive,
		CreatedAt:             utils.FormatTimestamp(item.CreatedAt),
		UpdatedAt:             utils.FormatTimestamp(item.UpdatedAt),
	}
	if item.HeadmanID != 0 {
		output.Headman = &ProfileOfficialOutput{
			ID:          item.HeadmanID,
			Name:        item.HeadmanName,
			Position:    item.HeadmanPosition,
			Phone:       item.HeadmanPhone,
			Email:       item.HeadmanEmail,
			Description: item.HeadmanDescription,
			OrderNumber: item.HeadmanOrderNumber,
			IsActive:    item.HeadmanIsActive,
			StartDate:   formatOptionalDate(item.HeadmanStartDate),
			FinishDate:  formatOptionalDatePointer(item.HeadmanFinishDate),
		}
	}
	return output
}

func parseDemographicJSON[T any](value string) []T {
	if strings.TrimSpace(value) == "" {
		return []T{}
	}
	var result []T
	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return []T{}
	}
	if result == nil {
		return []T{}
	}
	return result
}

func formatOptionalDate(value *time.Time) string {
	if value == nil {
		return ""
	}
	return value.Format("2006-01-02")
}

func formatOptionalDatePointer(value *time.Time) *string {
	if value == nil {
		return nil
	}
	formatted := value.Format("2006-01-02")
	return &formatted
}

func (s *service) CreateOfficial(ctx context.Context, payload AddOfficialPayload) error {
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Position = strings.TrimSpace(payload.Position)
	if payload.Name == "" {
		return fmt.Errorf("%w: name is required", utils.ErrInvalidPayload)
	}
	if payload.Position == "" {
		return fmt.Errorf("%w: position is required", utils.ErrInvalidPayload)
	}
	if utils.GenerateSlug(payload.Position) == headmanPosition || strings.EqualFold(payload.Position, "kepala desa") || strings.EqualFold(payload.Position, "kepala-desa") {
		return fmt.Errorf("%w: position kepala-desa cannot be added as official", utils.ErrInvalidPayload)
	}
	return s.repository.CreateOfficial(ctx, payload)
}

func (s *service) ListOfficials(ctx context.Context, payload ListOfficialPayload) ([]OfficialResponse, error) {
	utils.NormalizePagination(&payload.Limit, &payload.Index)
	return s.repository.ListOfficials(ctx, payload)
}

func (s *service) FindOfficialByID(ctx context.Context, payload OfficialPayload) (*OfficialResponse, error) {
	if payload.ID == 0 {
		return nil, fmt.Errorf("%w: official id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.FindOfficialByID(ctx, payload)
}

func (s *service) UpdateOfficial(ctx context.Context, payload EditOfficialPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: official id must be greater than 0", utils.ErrInvalidPayload)
	}
	payload.Name = strings.TrimSpace(payload.Name)
	payload.Position = strings.TrimSpace(payload.Position)
	if payload.Name == "" {
		return fmt.Errorf("%w: name is required", utils.ErrInvalidPayload)
	}
	if payload.Position == "" {
		return fmt.Errorf("%w: position is required", utils.ErrInvalidPayload)
	}
	if utils.GenerateSlug(payload.Position) == headmanPosition || strings.EqualFold(payload.Position, "kepala desa") || strings.EqualFold(payload.Position, "kepala-desa") {
		return fmt.Errorf("%w: position kepala-desa cannot be set for official", utils.ErrInvalidPayload)
	}
	return s.repository.UpdateOfficial(ctx, payload)
}

func (s *service) DeleteOfficial(ctx context.Context, payload OfficialPayload) error {
	if payload.ID == 0 {
		return fmt.Errorf("%w: official id must be greater than 0", utils.ErrInvalidPayload)
	}
	return s.repository.DeleteOfficial(ctx, payload)
}
