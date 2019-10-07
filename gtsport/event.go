package gtsport

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

func UnmarshalEventInfo(data []byte) (EventInfo, error) {
	var r EventInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EventInfo struct {
	Event []EventInfoEvent `json:"event"`
}

type EventInfoEvent struct {
	EventID    string         `json:"event_id"`
	CreateTime string         `json:"create_time"`
	UpdateTime string         `json:"update_time"`
	Value      []ValueElement `json:"value"`
}

type ValueClass struct {
	GameParameter GameParameter `json:"GameParameter"`
}

type GameParameter struct {
	Championship int64                `json:"championship"`
	Events       []GameParameterEvent `json:"events"`
	Tracks       []Track              `json:"tracks"`
}

type GameParameterEvent struct {
	EventID                         int64       `json:"event_id"`
	ChampionshipID                  int64       `json:"championship_id"`
	SeasonID                        int64       `json:"season_id"`
	RoundID                         int64       `json:"round_id"`
	ChampionshipColor               string      `json:"championship_color"`
	GameMode                        string      `json:"game_mode"`
	EventType                       string      `json:"event_type"`
	SportsMode                      string      `json:"sports_mode"`
	Information                     Information `json:"information"`
	EntrySet                        []EntrySet  `json:"entry_set"`
	Race                            Race        `json:"race"`
	Ranking                         Ranking     `json:"ranking"`
	Regulation                      Regulation  `json:"regulation"`
	RefuelingSpeed                  int64       `json:"refueling_speed"`
	TrackIndex                      int64       `json:"track_index"`
	MatchingMethod                  int64       `json:"matching_method"`
	IsUncertainEvent                int64       `json:"is_uncertain_event"`
	IsSeasonalEvent                 int64       `json:"is_seasonal_event"`
	BeginDate                       time.Time   `json:"begin_date"`
	EndDate                         time.Time   `json:"end_date"`
	OnlineTimeattackRegistrationKey interface{} `json:"online_timeattack_registration_key"`
	OnlineTimeattackType            interface{} `json:"online_timeattack_type"`
	NPRegions                       interface{} `json:"np_regions"`
	Track                           Track
}

type EntrySet struct {
	EntryGenerate EntryGenerate `json:"entry_generate"`
}

type EntryGenerate struct {
	PlayerEntryBaseArray interface{} `json:"player_entry_base_array"`
}

type Information struct {
	Title              Description `json:"title"`
	OneLineTitle       Description `json:"one_line_title"`
	Description        Description `json:"description"`
	RegistrationNotice Description `json:"registration_notice"`
	EventTarget        Description `json:"event_target"`
	EventRestrict      Description `json:"event_restrict"`
	LogoImagePath      string      `json:"logo_image_path"`
	FlyerImagePath     string      `json:"flyer_image_path"`
}

type Description struct {
	Jp string `json:"JP"`
	Us string `json:"US"`
	GB string `json:"GB"`
	Fr string `json:"FR"`
	De string `json:"DE"`
	It string `json:"IT"`
	Es string `json:"ES"`
	Pt string `json:"PT"`
	Nl string `json:"NL"`
	Ru string `json:"RU"`
	Kr string `json:"KR"`
	Tw string `json:"TW"`
	El string `json:"EL"`
	Tr string `json:"TR"`
	Pl string `json:"PL"`
	Cz string `json:"CZ"`
	Bp string `json:"BP"`
	MS string `json:"MS"`
	Ar string `json:"AR"`
	Th string `json:"TH"`
}

type Race struct {
	BehaviorDamageType     string      `json:"behavior_damage_type"`
	BehaviorSlipStreamType string      `json:"behavior_slip_stream_type"`
	BoostLevel             interface{} `json:"boost_level"`
	LowMuType              string      `json:"low_mu_type"`
	EntryMax               int         `json:"entry_max"`
	StartType              string      `json:"start_type"`
	ConsumeFuel            int         `json:"consume_fuel"`
	ConsumeTire            int         `json:"consume_tire"`
	RaceLimitLaps          int         `json:"race_limit_laps"`
	RaceLimitMinute        int         `json:"race_limit_minute"`
	PitConstraint          int         `json:"pit_constraint"`
	NeedCompoundUse        string      `json:"need_compound_use"`
	LimitFuelCapacity      int         `json:"limit_fuel_capacity"`
}

type Ranking struct {
	BoardID   int64       `json:"board_id"`
	BeginDate interface{} `json:"begin_date"`
	EndDate   interface{} `json:"end_date"`
}

type Regulation struct {
	LimitTireF       int         `json:"limit_tire_f"`
	NeedTireF        int         `json:"need_tire_f"`
	UseBop           int         `json:"use_bop"`
	Tuning           int         `json:"tuning"`
	NeedDriverClass  string      `json:"need_driver_class"`
	CarCategoryTypes []string    `json:"car_category_types"`
	Cars             interface{} `json:"cars"`
	CarsCount        int         `json:"cars_count"`
}

type Track struct {
	CourseCode  string      `json:"course_code"`
	WeatherList interface{} `json:"WeatherList"`
}

type ValueElement struct {
	String     *string
	ValueClass *ValueClass
}

func (x *ValueElement) UnmarshalJSON(data []byte) error {
	x.ValueClass = nil
	var c ValueClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ValueClass = &c
	}
	return nil
}

func (x *ValueElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ValueClass != nil, x.ValueClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
