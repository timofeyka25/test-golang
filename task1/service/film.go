package service

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"testGolang"
	"time"
)

type data struct {
	Films []testGolang.Film `json:"films"`
}

func (s *Service) ReadFilms() ([]testGolang.Film, error) {
	byteValue, err := ioutil.ReadFile("films.json")
	if err != nil {
		return nil, err
	}

	var result data
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		err = s.writeFilms(nil)
		return nil, err
	}

	return result.Films, nil
}

func (s *Service) GetFilms(sortParam string) ([]testGolang.FilmOutput, error) {
	var output []testGolang.FilmOutput
	films, err := s.ReadFilms()
	if err != nil {
		return nil, err
	}
	for _, v := range films {
		output = append(output,
			testGolang.FilmOutput{Id: v.Id, Title: v.Title, ReleaseDate: v.ReleaseDate, CreatedAt: v.CreatedAt})
	}
	if sortParam == "releaseDate" {
		sort.SliceStable(output, func(i, j int) bool {
			return output[i].ReleaseDate.Before(output[j].ReleaseDate)
		})
	} else if sortParam == "-releaseDate" {
		sort.SliceStable(output, func(i, j int) bool {
			return output[i].ReleaseDate.After(output[j].ReleaseDate)
		})
	}

	return output, nil
}

func (s *Service) writeFilms(films []testGolang.Film) error {
	writeData := data{Films: films}

	file, err := json.MarshalIndent(writeData, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("films.json", file, 0644)
	return err
}

func (s *Service) newId(films []testGolang.Film) (int, error) {
	if films == nil {
		return 1, nil
	}
	id := 0
	for _, v := range films {
		if v.Id > id {
			id = v.Id
		}
	}
	return id + 1, nil
}

func (s *Service) AddFilm(film testGolang.Film) error {
	films, err := s.ReadFilms()
	if err != nil {
		return err
	}
	film.Id, err = s.newId(films)
	if err != nil {
		return err
	}
	film.CreatedAt = time.Now()
	films = append(films, film)

	return s.writeFilms(films)
}

func (s *Service) GetFilmById(id int) (*testGolang.Film, error) {
	films, err := s.ReadFilms()
	if err != nil {
		return nil, err
	}

	for _, v := range films {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, nil
}
