package pokeapi

type PokemonData struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"ability,omitempty"`
		IsHidden bool    `json:"is_hidden,omitempty"`
		Slot     float64 `json:"slot,omitempty"`
	} `json:"abilities,omitempty"`
	BaseExperience float64 `json:"base_experience,omitempty"`
	Cries          struct {
		Latest string `json:"latest,omitempty"`
		Legacy string `json:"legacy,omitempty"`
	} `json:"cries,omitempty"`
	Forms []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"forms,omitempty"`
	GameIndices []struct {
		GameIndex float64 `json:"game_index,omitempty"`
		Version   struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"version,omitempty"`
	} `json:"game_indices,omitempty"`
	Height    float64 `json:"height,omitempty"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"item,omitempty"`
		VersionDetails []struct {
			Rarity  float64 `json:"rarity,omitempty"`
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"held_items,omitempty"`
	ID                     float64 `json:"id,omitempty"`
	IsDefault              bool    `json:"is_default,omitempty"`
	LocationAreaEncounters string  `json:"location_area_encounters,omitempty"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"move,omitempty"`
		VersionGroupDetails []struct {
			LevelLearnedAt  float64 `json:"level_learned_at,omitempty"`
			MoveLearnMethod struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"move_learn_method,omitempty"`
			Order        any `json:"order,omitempty"`
			VersionGroup struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version_group,omitempty"`
		} `json:"version_group_details,omitempty"`
	} `json:"moves,omitempty"`
	Name          string  `json:"name,omitempty"`
	Order         float64 `json:"order,omitempty"`
	PastAbilities []struct {
		Abilities []struct {
			Ability  any     `json:"ability,omitempty"`
			IsHidden bool    `json:"is_hidden,omitempty"`
			Slot     float64 `json:"slot,omitempty"`
		} `json:"abilities,omitempty"`
		Generation struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"generation,omitempty"`
	} `json:"past_abilities,omitempty"`
	PastStats []struct {
		Generation struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"generation,omitempty"`
		Stats []struct {
			BaseStat float64 `json:"base_stat,omitempty"`
			Effort   float64 `json:"effort,omitempty"`
			Stat     struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"stat,omitempty"`
		} `json:"stats,omitempty"`
	} `json:"past_stats,omitempty"`
	PastTypes []any `json:"past_types,omitempty"`
	Species   struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"species,omitempty"`
	Sprites struct {
		BackDefault      string `json:"back_default,omitempty"`
		BackFemale       string `json:"back_female,omitempty"`
		BackShiny        string `json:"back_shiny,omitempty"`
		BackShinyFemale  string `json:"back_shiny_female,omitempty"`
		FrontDefault     string `json:"front_default,omitempty"`
		FrontFemale      string `json:"front_female,omitempty"`
		FrontShiny       string `json:"front_shiny,omitempty"`
		FrontShinyFemale string `json:"front_shiny_female,omitempty"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontFemale  any    `json:"front_female,omitempty"`
			} `json:"dream_world,omitempty"`
			Home struct {
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      string `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale string `json:"front_shiny_female,omitempty"`
			} `json:"home,omitempty"`
			Official_artwork struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontShiny   string `json:"front_shiny,omitempty"`
			} `json:"official-artwork,omitempty"`
			Showdown struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackFemale       string `json:"back_female,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      string `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale string `json:"front_shiny_female,omitempty"`
			} `json:"showdown,omitempty"`
		} `json:"other,omitempty"`
		Versions struct {
			Generation_i struct {
				Red_blue struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackGray         string `json:"back_gray,omitempty"`
					BackTransparent  string `json:"back_transparent,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontGray        string `json:"front_gray,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"red-blue,omitempty"`
				Yellow struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackGray         string `json:"back_gray,omitempty"`
					BackTransparent  string `json:"back_transparent,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontGray        string `json:"front_gray,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"yellow,omitempty"`
			} `json:"generation-i,omitempty"`
			Generation_ii struct {
				Crystal struct {
					BackDefault           string `json:"back_default,omitempty"`
					BackShiny             string `json:"back_shiny,omitempty"`
					BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
					BackTransparent       string `json:"back_transparent,omitempty"`
					FrontDefault          string `json:"front_default,omitempty"`
					FrontShiny            string `json:"front_shiny,omitempty"`
					FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
					FrontTransparent      string `json:"front_transparent,omitempty"`
				} `json:"crystal,omitempty"`
				Gold struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"gold,omitempty"`
				Silver struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontTransparent string `json:"front_transparent,omitempty"`
				} `json:"silver,omitempty"`
			} `json:"generation-ii,omitempty"`
			Generation_iii struct {
				Emerald struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"emerald,omitempty"`
				Firered_leafgreen struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"firered-leafgreen,omitempty"`
				Ruby_sapphire struct {
					BackDefault  string `json:"back_default,omitempty"`
					BackShiny    string `json:"back_shiny,omitempty"`
					FrontDefault string `json:"front_default,omitempty"`
					FrontShiny   string `json:"front_shiny,omitempty"`
				} `json:"ruby-sapphire,omitempty"`
			} `json:"generation-iii,omitempty"`
			Generation_iv struct {
				Diamond_pearl struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"diamond-pearl,omitempty"`
				Heartgold_soulsilver struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"heartgold-soulsilver,omitempty"`
				Platinum struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"platinum,omitempty"`
			} `json:"generation-iv,omitempty"`
			Generation_ix struct {
				Scarlet_violet struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  any    `json:"front_female,omitempty"`
				} `json:"scarlet-violet,omitempty"`
			} `json:"generation-ix,omitempty"`
			Generation_v struct {
				Black_white struct {
					Animated struct {
						BackDefault      string `json:"back_default,omitempty"`
						BackFemale       string `json:"back_female,omitempty"`
						BackShiny        string `json:"back_shiny,omitempty"`
						BackShinyFemale  string `json:"back_shiny_female,omitempty"`
						FrontDefault     string `json:"front_default,omitempty"`
						FrontFemale      string `json:"front_female,omitempty"`
						FrontShiny       string `json:"front_shiny,omitempty"`
						FrontShinyFemale string `json:"front_shiny_female,omitempty"`
					} `json:"animated,omitempty"`
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       string `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  string `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"black-white,omitempty"`
			} `json:"generation-v,omitempty"`
			Generation_vi struct {
				Omegaruby_alphasapphire struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"omegaruby-alphasapphire,omitempty"`
				X_y struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"x-y,omitempty"`
			} `json:"generation-vi,omitempty"`
			Generation_vii struct {
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  any    `json:"front_female,omitempty"`
				} `json:"icons,omitempty"`
				Ultra_sun_ultra_moon struct {
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      string `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale string `json:"front_shiny_female,omitempty"`
				} `json:"ultra-sun-ultra-moon,omitempty"`
			} `json:"generation-vii,omitempty"`
			Generation_viii struct {
				Brilliant_diamond_shining_pearl struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  any    `json:"front_female,omitempty"`
				} `json:"brilliant-diamond-shining-pearl,omitempty"`
				Icons struct {
					FrontDefault string `json:"front_default,omitempty"`
					FrontFemale  string `json:"front_female,omitempty"`
				} `json:"icons,omitempty"`
			} `json:"generation-viii,omitempty"`
		} `json:"versions,omitempty"`
	} `json:"sprites,omitempty"`
	Stats []struct {
		BaseStat float64 `json:"base_stat,omitempty"`
		Effort   float64 `json:"effort,omitempty"`
		Stat     struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"stat,omitempty"`
	} `json:"stats,omitempty"`
	Types []struct {
		Slot float64 `json:"slot,omitempty"`
		Type struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"type,omitempty"`
	} `json:"types,omitempty"`
	Weight float64 `json:"weight,omitempty"`
}
