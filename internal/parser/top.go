package parser

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"top-ranking/internal/domain"
)

func Filter(source []*domain.Table, total int, item string) ([]*domain.Table, error) {
	var filtered []*domain.Table
	found := 0

	for _, rec := range source {
		if rec.Language == item {
			filtered = append(filtered, rec)
			found++
		}
		if found == total {
			break
		}
	}
	fmt.Printf("Registros encontrados %d", found)
	return filtered, nil
}

func FilterByLanguage(source []*domain.Table, language string) ([]*domain.Table, int) {
	var filtered []*domain.Table
	for _, rec := range source {
		if strings.ToLower(rec.Language) == strings.ToLower(language) {
			filtered = append(filtered, rec)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Rank < filtered[j].Rank
	})

	return filtered, len(filtered)
}

func ConverToJson(source []*domain.Table) ([]byte, error) {
	return json.Marshal(source)
}
