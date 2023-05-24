package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type location struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandMapF(p player) error {
	for i := 0; i < 20 && p.posId < p.posMaxId; i++ {
		res, err := http.Get("https://pokeapi.co/api/v2/location/" + strconv.Itoa(p.posId))
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			return err
		}
		if err != nil {
			fmt.Print("commandMapF IO error")
			return err
		}
		var obj location
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Print("commandMapF unmarshal error")
			return err
		}
		fmt.Print(obj.Name, "\n")
		p.posId++
	}
	return nil
}

func commandMapB(p player) error {
	for i := 0; i < 20 && p.posId > 1; i++ {
		p.posId--
		res, err := http.Get("https://pokeapi.co/api/v2/location/" + strconv.Itoa(p.posId))
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			return err
		}
		if err != nil {
			fmt.Print("commandMapF IO error")
			return err
		}
		var obj location
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Print("commandMapF unmarshal error")
			return err
		}
		fmt.Print(obj.Name, "\n")
	}
	return nil
}
