// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    seasonInfo, err := UnmarshalSeasonInfo(bytes)
//    bytes, err = seasonInfo.Marshal()

package gtsport

import (
	"encoding/json"
	"time"
)

func UnmarshalSeasonInfo(data []byte) (SeasonInfo, error) {
	var r SeasonInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SeasonInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SeasonInfo struct {
	SeasonInformation SeasonInformation `json:"season_information"`
}

type SeasonInformation struct {
	Seasons []Season `json:"seasons"`
}

type Season struct {
	SeasonID        int64   `json:"season_id"`
	ChampionshipID  int64   `json:"championship_id"`
	GlobalRankingID int64   `json:"global_ranking_id"`
	Color           Color   `json:"color"`
	Rounds          []Round `json:"rounds"`
}

type Color struct {
	A string `json:"a"`
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

type Round struct {
	RoundID                    string          `json:"round_id"`
	StartDate                  time.Time       `json:"start_date"`
	EndDate                    time.Time       `json:"end_date"`
	CarCategoryType            CarCategoryType `json:"car_category_type"`
	EventID                    string          `json:"event_id"`
	SportsMode                 SportsMode      `json:"sports_mode"`
	ChampionshipValidRaceCount string          `json:"championship_valid_race_count"`
	LogoImagePath              LogoImagePath   `json:"logo_image_path"`
	FlyerImagePath             FlyerImagePath  `json:"flyer_image_path"`
	Title                      EventTarget     `json:"title"`
	EventTarget                EventTarget     `json:"event_target"`
	CourseCode                 string          `json:"course_code"`
	MatchingMethod             int64           `json:"matching_method"`
	IsUncertainEvent           int64           `json:"is_uncertain_event"`
	RegulationCarsCount        *int64          `json:"regulation_cars_count,omitempty"`
}

type EventTarget struct {
	Jp *string `json:"JP,omitempty"`
	Us *string `json:"US,omitempty"`
	GB *string `json:"GB,omitempty"`
	Fr *string `json:"FR,omitempty"`
	De *string `json:"DE,omitempty"`
	It *string `json:"IT,omitempty"`
	Es *string `json:"ES,omitempty"`
	Pt *string `json:"PT,omitempty"`
	Nl *string `json:"NL,omitempty"`
	Ru *string `json:"RU,omitempty"`
	Kr *string `json:"KR,omitempty"`
	Tw *string `json:"TW,omitempty"`
	El *string `json:"EL,omitempty"`
	Tr *string `json:"TR,omitempty"`
	Pl *string `json:"PL,omitempty"`
	Cz *string `json:"CZ,omitempty"`
	Bp *string `json:"BP,omitempty"`
	MS *string `json:"MS,omitempty"`
	Ar *string `json:"AR,omitempty"`
	Th *string `json:"TH,omitempty"`
}

type CarCategoryType string

const (
	Gr1  CarCategoryType = "GR1"
	Gr2  CarCategoryType = "GR2"
	Gr3  CarCategoryType = "GR3"
	Gr4  CarCategoryType = "GR4"
	Grb  CarCategoryType = "GRB"
	Grx  CarCategoryType = "GRX"
	N100 CarCategoryType = "N100"
	N200 CarCategoryType = "N200"
	N300 CarCategoryType = "N300"
	N400 CarCategoryType = "N400"
	N500 CarCategoryType = "N500"
	N600 CarCategoryType = "N600"
)

type FlyerImagePath string

const (
	Img132_Ad7Ed60Cd617294486A19Ebec6F70C5BPNG FlyerImagePath = "img_132_ad7ed60cd617294486a19ebec6f70c5b.png"
	Img133_C15C75C6B35D3930C8Eab34E27Fee07EPNG FlyerImagePath = "img_133_c15c75c6b35d3930c8eab34e27fee07e.png"
	Img14PNG                                   FlyerImagePath = "img_14.png"
	Img15PNG                                   FlyerImagePath = "img_15.png"
	Img23PNG                                   FlyerImagePath = "img_23.png"
	Img25PNG                                   FlyerImagePath = "img_25.png"
	Img71PNG                                   FlyerImagePath = "img_71.png"
	Img73PNG                                   FlyerImagePath = "img_73.png"
	Img85PNG                                   FlyerImagePath = "img_85.png"
	Img88PNG                                   FlyerImagePath = "img_88.png"
	Img90PNG                                   FlyerImagePath = "img_90.png"
	Img92PNG                                   FlyerImagePath = "img_92.png"
	Img98PNG                                   FlyerImagePath = "img_98.png"
)

type LogoImagePath string

const (
	Img131_6B91F4D083B5866C964Db2C483559F03PNG LogoImagePath = "img_131_6b91f4d083b5866c964db2c483559f03.png"
	Img134_Af12039Fa84730789Adce2Ac9439A827PNG LogoImagePath = "img_134_af12039fa84730789adce2ac9439a827.png"
	Img22PNG                                   LogoImagePath = "img_22.png"
	Img24PNG                                   LogoImagePath = "img_24.png"
	Img6PNG                                    LogoImagePath = "img_6.png"
	Img70PNG                                   LogoImagePath = "img_70.png"
	Img72PNG                                   LogoImagePath = "img_72.png"
	Img7PNG                                    LogoImagePath = "img_7.png"
	Img84PNG                                   LogoImagePath = "img_84.png"
	Img87PNG                                   LogoImagePath = "img_87.png"
	Img89PNG                                   LogoImagePath = "img_89.png"
	Img91PNG                                   LogoImagePath = "img_91.png"
	Img97PNG                                   LogoImagePath = "img_97.png"
)

type SportsMode string

const (
	GeneralChampionship SportsMode = "GENERAL_CHAMPIONSHIP"
	Manufacturer        SportsMode = "MANUFACTURER"
	Nations             SportsMode = "NATIONS"
)
