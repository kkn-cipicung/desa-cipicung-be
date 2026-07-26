package infografis

import (
	"context"
	"encoding/json"
	"strings"
)

type Service interface {
	Detail(ctx context.Context) (*InfografisOutput, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Detail(ctx context.Context) (*InfografisOutput, error) {
	stats, err := s.repository.Detail(ctx)
	if err != nil {
		return nil, err
	}

	return mapInfografisOutput(stats), nil
}

func mapInfografisOutput(stats *VillageStats) *InfografisOutput {
	if stats == nil {
		stats = &VillageStats{}
	}

	return &InfografisOutput{
		Summary: InfografisSummary{
			Population: stats.Population,
			Family:     stats.Family,
			Male:       stats.Male,
			Female:     stats.Female,
		},
		Hamlets: []InfografisSlice{
			{Label: "Dusun I", Value: stats.HamletOne, Color: "#8f6b4f"},
			{Label: "Dusun II", Value: stats.HamletTwo, Color: "#314f43"},
		},
		Religions:  parseJSON[InfografisSlice](stats.ReligionsJSON),
		ReligionRT: parseJSON[InfografisBar](stats.ReligionRTJSON),
		Education:  parseJSON[InfografisGenderBar](stats.EducationJSON),
		Occupation: parseJSON[InfografisGenderBar](stats.OccupationJSON),
		Ages:       parseJSON[InfografisSingleBar](stats.AgesJSON),
		Source:     "RPJM Desa Cipicung Perubahan 2025-2029, data BPS & Posyandu 2024",
	}
}

func parseJSON[T any](value string) []T {
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
