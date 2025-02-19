package consts

import (
	"fmt"
	"strings"
)

type Stats struct {
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

func (s Stats) String() string {
	return fmt.Sprintf(
		"Stats:\n  -hp: %d\n  -attack: %d\n  -defense: %d\n  -special-attack: %d\n  -special-defense: %d\n  -speed: %d",
		s.HP, s.Attack, s.Defense, s.SpecialAttack, s.SpecialDefense, s.Speed,
	)
}

type PokemonInspectInfo struct {
	Name   string
	Height int
	Weight int
	Stats  Stats
	Types  []string
}

func (p PokemonInspectInfo) String() string {
	return fmt.Sprintf(
		"Name: %s\nHeight: %d\nWeight: %d\n%s\n%s",
		p.Name, p.Height, p.Weight, p.Stats.String(), strings.Join(append([]string{"Types:"}, p.Types...), "\n  -"),
	)
}

type PlayerInfo struct {
	CaughtPokemons map[string]*PokemonInspectInfo
}

type Command func(args []string, playerInfo *PlayerInfo) error

type VersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource `json:"encounter_method"`
	VersionDetails  []VersionDetail  `json:"version_details"`
}

type Name struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type Encounter struct {
	Pokemon        NamedAPIResource `json:"pokemon"`
	VersionDetails []VersionDetail  `json:"version_details"`
}

type LocationInfo struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	Location             NamedAPIResource      `json:"location"`
	PokemonEncounters    []Encounter           `json:"pokemon_encounters"`
}

type Ability struct {
	IsHidden    bool             `json:"is_hidden"`
	Slot        int              `json:"slot"`
	AbilityInfo NamedAPIResource `json:"ability"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type HeldItem struct {
	Item          NamedAPIResource `json:"item"`
	VersionDetail VersionDetail    `json:"version_detail"`
}

type VersionDetailGroup struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
}

type MoveDetail struct {
	Move                NamedAPIResource     `json:"move"`
	VersionDetailGroups []VersionDetailGroup `json:"version_group_details"`
}

type SpriteDetail struct {
	BackDefault      string         `json:"back_default"`
	BackFemale       string         `json:"back_female"`
	BackShiny        string         `json:"back_shiny"`
	BackShinyFemale  string         `json:"back_shiny_female"`
	FrontDefault     string         `json:"front_default"`
	FrontFemale      string         `json:"front_female"`
	FrontShiny       string         `json:"front_shiny"`
	FrontShinyFemale string         `json:"front_shiny_female"`
	Other            map[string]any `json:"other"`
	Versions         map[string]any `json:"version"`
}

type CryInfo struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type StatInfo struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

type TypeInfo struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type GenerationType struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []TypeInfo       `json:"types"`
}

type APIResource struct {
	URL string `json:"url"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonInfo struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []Ability          `json:"abilities"`
	Forms                  []NamedAPIResource `json:"form"`
	GameIndices            []GameIndex        `json:"game_indices"`
	HeldItem               []HeldItem         `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []MoveDetail       `json:"move"`
	Species                NamedAPIResource   `json:"species"`
	Sprites                SpriteDetail       `json:"sprites"`
	Cries                  CryInfo            `json:"cries"`
	Stats                  []StatInfo         `json:"stats"`
	Types                  []TypeInfo         `json:"types"`
	PastTypes              []GenerationType   `json:"past_types"`
}

type PokedexEntry struct {
	EntryNumber int              `json:"entry_number"`
	Pokedex     NamedAPIResource `json:"pokedex"`
}

type TextEntry struct {
	FlavorText string           `json:"flavor_text"`
	Language   NamedAPIResource `json:"language"`
	Version    NamedAPIResource `json:"version"`
}

type FormDescription struct {
	Description string           `json:"description"`
	Language    NamedAPIResource `json:"language"`
}

type Genera struct {
	Genus    string           `json:"genus"`
	Language NamedAPIResource `json:"language"`
}

type Variety struct {
	IsDefault bool             `json:"is_default"`
	Pokemon   NamedAPIResource `json:"pokemon"`
}

type PokemonSpecies struct {
	ID                   int                `json:"id"`
	Name                 string             `json:"name"`
	Order                int                `json:"order"`
	GenderRate           int                `json:"gender_rate"`
	CaptureRate          int                `json:"capture_rate"`
	BaseHappiness        int                `json:"base_happiness"`
	IsBaby               bool               `json:"is_baby"`
	IsLegendary          bool               `json:"is_legendary"`
	IsMythical           bool               `json:"is_mythical"`
	HatchCounter         int                `json:"hatch_country"`
	HasGenderDifferences bool               `json:"has_gender_differences"`
	FormsSwitchable      bool               `json:"forms_switchable"`
	GrowthRate           NamedAPIResource   `json:"growth_rate"`
	PokedexNumbers       []PokedexEntry     `json:"pokedex_numbers"`
	EggGroups            []NamedAPIResource `json:"egg_groups"`
	Color                NamedAPIResource   `json:"color"`
	Shape                NamedAPIResource   `json:"shape"`
	EvolvesFromSpecies   NamedAPIResource   `json:"evolves_from_species"`
	EvolutionChain       APIResource        `json:"evolution_chain"`
	Habitat              NamedAPIResource   `json:"habitat"`
	Generation           NamedAPIResource   `json:"generation"`
	Names                []Name             `json:"names"`
	FlavorTextEntries    []TextEntry        `json:"flavor_text_entries"`
	FormDescriptions     []FormDescription  `json:"form_descriptions"`
	Genera               []Genera           `json:"genera"`
	Varieties            []Variety          `json:"varieties"`
}
