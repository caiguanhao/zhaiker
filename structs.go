package zhaiker

import (
	"encoding/json"
	"strconv"
	"time"
)

type Gym struct {
	Address    string
	Admin      string
	AdminPhone string
	City       string
	GmtCreate  Timestamp
	GmtModify  Timestamp
	GpsX       float64
	GpsY       float64
	GymName    string
	Id         string
	Images     stringSlice
	LoginName  string
	Logo       string
	Province   string
	State      string
}

type Exam struct {
	AerobicGoal      float64
	Age              int
	Agility          int
	AnaGoal          float64
	Balance          int
	BalanceAngle     float64
	BloodMaxPressure int
	BloodMinPressure int
	BloodOxygen      int
	Bmi              float64
	Bmr              float64
	BodyAge          int
	BodyDetect       *ScoreXYByBodyParts
	BodyImage        string
	BodyShape        int
	BodyShapeRisk    int
	Bone             float64
	CaloriesInput    float64
	CustomerID       string
	DeviceFrom       string
	EnduGoal         float64
	Fat              float64
	FatControl       float64
	FatLeftArm       float64
	FatLeftLeg       float64
	FatRightArm      float64
	FatRightLeg      float64
	FatTrunk         float64
	GmtCreate        Timestamp
	GmtModify        Timestamp
	GymID            string
	GymName          string
	HeartFun         int
	Height           float64
	Hips             float64
	HumpbackRisk     int
	Id               string
	InFat            float64
	Language         string
	LegRisk          string
	Muscle           float64
	MuscleControl    float64
	MuscleLeftArm    float64
	MuscleLeftLeg    float64
	MuscleRightArm   float64
	MuscleRightLeg   float64
	MuscleTrunk      float64
	Name             string
	PelvisRisk       int
	PerfectWeight    float64
	Phone            string
	PostureRisk      int
	Protein          float64
	R100LeftArm      float64
	R100LeftLeg      float64
	R100RightArm     float64
	R100RightLeg     float64
	R100Trunk        float64
	R20LeftArm       float64
	R20LeftLeg       float64
	R20RightArm      float64
	R20RightLeg      float64
	R20Trunk         float64
	RestingHeartRate int
	Score            int
	Sex              string
	ShoulderRisk     int
	SideBodyDetect   *ScoreXYByBodyParts
	SideImage        string
	SportGoal        float64
	SportLevel       int
	SportSafeRisk    int
	SubFat           float64
	UnitType         string
	UserID           string
	VitalCapacity    int
	Waist            float64
	Water            float64
	Wc               float64
	Weight           float64
	WeightControl    float64
}

type ScoreXY struct {
	Score float64
	X     float64
	Y     float64
}

type ScoreXYByBodyParts struct {
	BodyDirection    int
	FullBody         bool
	LeftAnkle        ScoreXY
	LeftEar          ScoreXY
	LeftElbow        ScoreXY
	LeftEye          ScoreXY
	LeftHip          ScoreXY
	LeftKnee         ScoreXY
	LeftMouthCorner  ScoreXY
	LeftShoulder     ScoreXY
	LeftWrist        ScoreXY
	Neck             ScoreXY
	Nose             ScoreXY
	RightAnkle       ScoreXY
	RightEar         ScoreXY
	RightElbow       ScoreXY
	RightEye         ScoreXY
	RightHip         ScoreXY
	RightKnee        ScoreXY
	RightMouthCorner ScoreXY
	RightShoulder    ScoreXY
	RightWrist       ScoreXY
	TopHead          ScoreXY
}

func (s *ScoreXYByBodyParts) UnmarshalJSON(data []byte) error {
	raw, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	origin := scoreXYByBodyPartsOrigin{}
	err = json.Unmarshal([]byte(raw), &origin)
	if err != nil {
		return err
	}
	*s = origin.AsScoreXYByBodyParts()
	return nil
}

type scoreXYByBodyPartsOrigin struct {
	BodyDirection    int     `json:"bodyDirection"`
	FullBody         bool    `json:"fullBody"`
	LeftAnkle        ScoreXY `json:"left_ankle"`
	LeftEar          ScoreXY `json:"left_ear"`
	LeftElbow        ScoreXY `json:"left_elbow"`
	LeftEye          ScoreXY `json:"left_eye"`
	LeftHip          ScoreXY `json:"left_hip"`
	LeftKnee         ScoreXY `json:"left_knee"`
	LeftMouthCorner  ScoreXY `json:"left_mouth_corner"`
	LeftShoulder     ScoreXY `json:"left_shoulder"`
	LeftWrist        ScoreXY `json:"left_wrist"`
	Neck             ScoreXY `json:"neck"`
	Nose             ScoreXY `json:"nose"`
	RightAnkle       ScoreXY `json:"right_ankle"`
	RightEar         ScoreXY `json:"right_ear"`
	RightElbow       ScoreXY `json:"right_elbow"`
	RightEye         ScoreXY `json:"right_eye"`
	RightHip         ScoreXY `json:"right_hip"`
	RightKnee        ScoreXY `json:"right_knee"`
	RightMouthCorner ScoreXY `json:"right_mouth_corner"`
	RightShoulder    ScoreXY `json:"right_shoulder"`
	RightWrist       ScoreXY `json:"right_wrist"`
	TopHead          ScoreXY `json:"top_head"`
}

func (origin scoreXYByBodyPartsOrigin) AsScoreXYByBodyParts() ScoreXYByBodyParts {
	return ScoreXYByBodyParts{
		BodyDirection:    origin.BodyDirection,
		FullBody:         origin.FullBody,
		LeftAnkle:        origin.LeftAnkle,
		LeftEar:          origin.LeftEar,
		LeftElbow:        origin.LeftElbow,
		LeftEye:          origin.LeftEye,
		LeftHip:          origin.LeftHip,
		LeftKnee:         origin.LeftKnee,
		LeftMouthCorner:  origin.LeftMouthCorner,
		LeftShoulder:     origin.LeftShoulder,
		LeftWrist:        origin.LeftWrist,
		Neck:             origin.Neck,
		Nose:             origin.Nose,
		RightAnkle:       origin.RightAnkle,
		RightEar:         origin.RightEar,
		RightElbow:       origin.RightElbow,
		RightEye:         origin.RightEye,
		RightHip:         origin.RightHip,
		RightKnee:        origin.RightKnee,
		RightMouthCorner: origin.RightMouthCorner,
		RightShoulder:    origin.RightShoulder,
		RightWrist:       origin.RightWrist,
		TopHead:          origin.TopHead,
	}
}

type ExamWithStandards struct {
	Exam                 Exam
	BasicStandard        *BasicStandard
	WeightGrowthStandard *StandardValues
	HeightGrowthStandard *StandardValues
}

func (s *ExamWithStandards) UnmarshalJSON(data []byte) error {
	origin := examWithStandardsOrigin{}
	err := json.Unmarshal(data, &origin)
	if err != nil {
		return err
	}
	*s = origin.AsExamWithStandards()
	return nil
}

type examWithStandardsOrigin struct {
	Exam                 Exam                 `json:"DATA"`
	BasicStandard        *basicStandardOrigin `json:"standard"`
	WeightGrowthStandard *StandardValues      `json:"weightGrowthStandard"`
	HeightGrowthStandard *StandardValues      `json:"heightGrowthStandard"`
}

func (origin examWithStandardsOrigin) AsExamWithStandards() ExamWithStandards {
	var s *BasicStandard = nil
	if origin.BasicStandard != nil {
		x := origin.BasicStandard.AsBasicStandard()
		s = &x
	}
	return ExamWithStandards{
		Exam:                 origin.Exam,
		BasicStandard:        s,
		WeightGrowthStandard: origin.WeightGrowthStandard,
		HeightGrowthStandard: origin.HeightGrowthStandard,
	}
}

type BasicStandard struct {
	BmaxpStandard    StandardValue
	BmiStandard      StandardValue
	BminpStandard    StandardValue
	BmrStandard      StandardValue
	BoneStandard     StandardValue
	BottomFat        float64
	BottomMuscle     float64
	FatStandard      StandardValue
	HeartStandard    StandardValue
	HeightStandard   StandardValue
	MuscleStandard   StandardValue
	ProteinStandard  StandardValue
	ReactionStandard StandardValue
	SubfatStandard   StandardValue
	TopFat           float64
	TopMuscle        float64
	TrunkFat         float64
	TrunkMuscle      float64
	VisfatStandard   StandardValue
	WaterStandard    StandardValue
	WeightStandard   StandardValue
	WhrStandard      StandardValue
}

type basicStandardOrigin struct {
	BmaxpStandard    StandardRange
	BmiStandard      StandardRange
	BminpStandard    StandardRange
	BmrStandard      StandardRange
	BoneStandard     StandardRange
	BottomFat        float64
	BottomMuscle     float64
	FatStandard      StandardRange
	HeartStandard    StandardRange
	HeightStandard   StandardRange
	MuscleStandard   StandardRange
	ProteinStandard  StandardRange
	ReactionStandard StandardRange
	SubfatStandard   StandardRange
	TopFat           float64
	TopMuscle        float64
	TrunkFat         float64
	TrunkMuscle      float64
	VisfatStandard   StandardRange
	WaterStandard    StandardRange
	WeightStandard   StandardRange
	WhrStandard      StandardRange
}

func (o basicStandardOrigin) AsBasicStandard() BasicStandard {
	return BasicStandard{
		BmaxpStandard:    o.BmaxpStandard.AsStandardValue(),
		BmiStandard:      o.BmiStandard.AsStandardValue(),
		BminpStandard:    o.BminpStandard.AsStandardValue(),
		BmrStandard:      o.BmrStandard.AsStandardValue(),
		BoneStandard:     o.BoneStandard.AsStandardValue(),
		BottomFat:        o.BottomFat,
		BottomMuscle:     o.BottomMuscle,
		FatStandard:      o.FatStandard.AsStandardValue(),
		HeartStandard:    o.HeartStandard.AsStandardValue(),
		HeightStandard:   o.HeightStandard.AsStandardValue(),
		MuscleStandard:   o.MuscleStandard.AsStandardValue(),
		ProteinStandard:  o.ProteinStandard.AsStandardValue(),
		ReactionStandard: o.ReactionStandard.AsStandardValue(),
		SubfatStandard:   o.SubfatStandard.AsStandardValue(),
		TopFat:           o.TopFat,
		TopMuscle:        o.TopMuscle,
		TrunkFat:         o.TrunkFat,
		TrunkMuscle:      o.TrunkMuscle,
		VisfatStandard:   o.VisfatStandard.AsStandardValue(),
		WaterStandard:    o.WaterStandard.AsStandardValue(),
		WeightStandard:   o.WeightStandard.AsStandardValue(),
		WhrStandard:      o.WhrStandard.AsStandardValue(),
	}
}

type StandardRange struct {
	Age      int
	Fewer    float64
	Fewest   float64
	More     float64
	Most     float64
	Standard float64
}

func (sr StandardRange) AsStandardValue() StandardValue {
	return StandardValue{float64(sr.Age), sr.Fewest, sr.Fewer, sr.Standard, sr.More, sr.Most}
}

type StandardValues []StandardValue

type StandardValue [6]float64

type Device struct {
	DeviceID   string
	DeviceName string
	GmtCreate  Timestamp
	GmtModify  Timestamp
	GymID      string
	GymName    string
	IsDeleted  string
	Service    string
	State      string
	Type       string
}

type stringSlice []string

func (ts *stringSlice) UnmarshalJSON(data []byte) error {
	raw, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	var strs []string
	err = json.Unmarshal([]byte(raw), &strs)
	if err != nil {
		return err
	}
	*ts = strs
	return nil
}

type Timestamp time.Time

func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*ts = Timestamp(time.Unix(i/1000, 0).UTC())
	return nil
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ts))
}

func (ts Timestamp) String() string {
	return time.Time(ts).String()
}
