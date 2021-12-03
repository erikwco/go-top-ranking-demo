package cli

import (
	"github.com/gocarina/gocsv"
	"net/http"
	"top-ranking/internal/domain"
)

func GetCSVData() ([]*domain.Table, error) {
	res, err := http.Get("https://raw.githubusercontent.com/EvanLi/Github-Ranking/master/Data/github-ranking-2018-12-18.csv")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var records []*domain.Table
	err = gocsv.Unmarshal(res.Body, &records)
	return records, err

}
