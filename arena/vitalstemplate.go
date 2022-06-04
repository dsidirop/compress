package arena

type VitalsTemplate struct {
	Spec VitalsTemplateSpec `                        json:"spec,omitempty"                       msg:"spec,omitempty"                     bson:"spec,omitempty"                      `
	Foo  *int               `                        json:"foo,omitempty"                        msg:"foo,omitempty"                      bson:"foo,omitempty"                       ` //stupid hack to dodge stupid msgp error
}

type VitalsTemplateSpec struct {
	RespiratoryRate       *int64             `       json:"respiratory_rate,omitempty"           msg:"respiratory_rate,omitempty"         bson:"respiratory_rate,omitempty"          `
	SpO2                  *int64             `       json:"spo2,omitempty"                       msg:"spo2,omitempty"                     bson:"spo2,omitempty"                      `
	BloodPressure         *BloodPressure     `       json:"blood_pressure,omitempty"             msg:"blood_pressure,omitempty"           bson:"blood_pressure,omitempty"            `
	EtCO2                 *int64             `       json:"etco2,omitempty"                      msg:"etco2,omitempty"                    bson:"etco2,omitempty"                     `
	TemperatureBlood      *float64           `       json:"temperature_blood,omitempty"          msg:"temperature_blood,omitempty"        bson:"temperature_blood,omitempty"         `
	TemperaturePeripheral *float64           `       json:"temperature_peripheral,omitempty"     msg:"temperature_peripheral,omitempty"   bson:"temperature_peripheral,omitempty"    `
	Ecg                   *Ecg               `       json:"ecg,omitempty"                        msg:"ecg,omitempty"                      bson:"ecg,omitempty"                       `
	Eyes                  *Eyes              `       json:"eyes,omitempty"                       msg:"eyes,omitempty"                     bson:"eyes,omitempty"                      `
	AirwayObstruction     *AirwayObstruction `       json:"airway_obstruction,omitempty"         msg:"airway_obstruction,omitempty"       bson:"airway_obstruction,omitempty"        `
	LungSounds            *LungSounds        `       json:"lung_sounds,omitempty"                msg:"lung_sounds,omitempty"              bson:"lung_sounds,omitempty"               `
	HeartSounds           *HeartSounds       `       json:"heart_sounds,omitempty"               msg:"heart_sounds,omitempty"             bson:"heart_sounds,omitempty"              `
	BowelSounds           *BowelSounds       `       json:"bowel_sounds,omitempty"               msg:"bowel_sounds,omitempty"             bson:"bowel_sounds,omitempty"              `
	Pulses                *Pulses            `       json:"pulses,omitempty"                     msg:"pulses,omitempty"                   bson:"pulses,omitempty"                    `
}

type BloodPressure struct {
	Systolic  *int64 `                               json:"systolic,omitempty"                   msg:"systolic,omitempty"                 bson:"systolic,omitempty"                  `
	Diastolic *int64 `                               json:"diastolic,omitempty"                  msg:"diastolic,omitempty"                bson:"diastolic,omitempty"                 `
}

type Ecg struct {
	HeartRate    *int64        `                     json:"heart_rate,omitempty"                 msg:"heart_rate,omitempty"               bson:"heart_rate,omitempty"                `
	BasicRhythm  *string       `                     json:"basic_rhythm,omitempty"               msg:"basic_rhythm,omitempty"             bson:"basic_rhythm,omitempty"              `
	Extrasystole *Extrasystole `                     json:"extrasystole,omitempty"               msg:"extrasystole,omitempty"             bson:"extrasystole,omitempty"              `
	Severity     *int64        `                     json:"severity,omitempty"                   msg:"severity,omitempty"                 bson:"severity,omitempty"                  `
	AllowPacing  *bool         `                     json:"allow_pacing,omitempty"               msg:"allow_pacing,omitempty"             bson:"allow_pacing,omitempty"              `
	EmdPea       *bool         `                     json:"emd_pea,omitempty"                    msg:"emd_pea,omitempty"                  bson:"emd_pea,omitempty"                   `
}

type Extrasystole struct {
	Type        *string `                            json:"type,omitempty"                       msg:"type,omitempty"                     bson:"type,omitempty"                      `
	Probability *int64  `                            json:"probability,omitempty"                msg:"probability,omitempty"              bson:"probability,omitempty"               ` //0 - 100
}

type Eyes struct {
	Left      *Eye    `                              json:"left,omitempty"                       msg:"left,omitempty"                     bson:"left,omitempty"                      `
	Right     *Eye    `                              json:"right,omitempty"                      msg:"right,omitempty"                    bson:"right,omitempty"                     `
	BlinkRate *string `                              json:"blink_rate,omitempty"                 msg:"blink_rate,omitempty"               bson:"blink_rate,omitempty"                `
}

type Eye struct {
	EyelidPosition      *string `                    json:"eyelid_position,omitempty"            msg:"eyelid_position,omitempty"          bson:"eyelid_position,omitempty"           `
	PupilSize           *int64  `                    json:"pupil_size,omitempty"                 msg:"pupil_size,omitempty"               bson:"pupil_size,omitempty"                `
	PupilResponsiveness *string `                    json:"pupil_responsiveness,omitempty"       msg:"pupil_responsiveness,omitempty"     bson:"pupil_responsiveness,omitempty"      `
}

type AirwayObstruction struct {
	TongueEdema *string `                            json:"tongue_edema,omitempty"               msg:"tongue_edema,omitempty"             bson:"tongue_edema,omitempty"              `
}

type LungSounds struct {
	Left  *SingleLungSounds `                        json:"left,omitempty"                       msg:"left,omitempty"                     bson:"left,omitempty"                      `
	Right *SingleLungSounds `                        json:"right,omitempty"                      msg:"right,omitempty"                    bson:"right,omitempty"                     `
}

type HeartSounds struct {
	Aortic          *string `                        json:"aortic,omitempty"                     msg:"aortic,omitempty"                   bson:"aortic,omitempty"                    `
	AorticVolume    *int64  `                        json:"aortic_volume,omitempty"              msg:"aortic_volume,omitempty"            bson:"aortic_volume,omitempty"             `
	Pulmonary       *string `                        json:"pulmonary,omitempty"                  msg:"pulmonary,omitempty"                bson:"pulmonary,omitempty"                 `
	PulmonaryVolume *int64  `                        json:"pulmonary_volume,omitempty"           msg:"pulmonary_volume,omitempty"         bson:"pulmonary_volume,omitempty"          `
	Tricuspid       *string `                        json:"tricuspid,omitempty"                  msg:"tricuspid,omitempty"                bson:"tricuspid,omitempty"                 `
	TricuspidVolume *int64  `                        json:"tricuspid_volume,omitempty"           msg:"tricuspid_volume,omitempty"         bson:"tricuspid_volume,omitempty"          `
	Mitral          *string `                        json:"mitral,omitempty"                     msg:"mitral,omitempty"                   bson:"mitral,omitempty"                    `
	MitralVolume    *int64  `                        json:"mitral_volume,omitempty"              msg:"mitral_volume,omitempty"            bson:"mitral_volume,omitempty"             `
}

type SingleLungSounds struct {
	AnteriorUpper        *string `                   json:"anterior_upper,omitempty"             msg:"anterior_upper,omitempty"           bson:"anterior_upper,omitempty"            `
	AnteriorUpperVolume  *int64  `                   json:"anterior_upper_volume,omitempty"      msg:"anterior_upper_volume,omitempty"    bson:"anterior_upper_volume,omitempty"     ` // [0-100]
	AnteriorLower        *string `                   json:"anterior_lower,omitempty"             msg:"anterior_lower,omitempty"           bson:"anterior_lower,omitempty"            `
	AnteriorLowerVolume  *int64  `                   json:"anterior_lower_volume,omitempty"      msg:"anterior_lower_volume,omitempty"    bson:"anterior_lower_volume,omitempty"     ` // [0-100]
	AnteriorMiddle       *string `                   json:"anterior_middle,omitempty"            msg:"anterior_middle,omitempty"          bson:"anterior_middle,omitempty"           ` // Note: left lung doesn't have anterior middle sound
	AnteriorMiddleVolume *int64  `                   json:"anterior_middle_volume,omitempty"     msg:"anterior_middle_volume,omitempty"   bson:"anterior_middle_volume,omitempty"    ` // [0-100]
	PosteriorUpper       *string `                   json:"posterior_upper,omitempty"            msg:"posterior_upper,omitempty"          bson:"posterior_upper,omitempty"           `
	PosteriorUpperVolume *int64  `                   json:"posterior_upper_volume,omitempty"     msg:"posterior_upper_volume,omitempty"   bson:"posterior_upper_volume,omitempty"    ` // [0-100]
	PosteriorLower       *string `                   json:"posterior_lower,omitempty"            msg:"posterior_lower,omitempty"          bson:"posterior_lower,omitempty"           `
	PosteriorLowerVolume *int64  `                   json:"posterior_lower_volume,omitempty"     msg:"posterior_lower_volume,omitempty"   bson:"posterior_lower_volume,omitempty"    ` // [0-100]
}

type BowelSounds struct {
	Bowel       *string `                            json:"bowel,omitempty"                      msg:"bowel,omitempty"                    bson:"bowel,omitempty"                     `
	BowelVolume *int64  `                            json:"bowel_volume,omitempty"               msg:"bowel_volume,omitempty"             bson:"bowel_volume,omitempty"              `
}

type Pulses struct {
	Central          *string `                       json:"central,omitempty"                    msg:"central,omitempty"                  bson:"central,omitempty"                   `
	CentralLimited   *string `                       json:"central_limited,omitempty"            msg:"central_limited,omitempty"          bson:"central_limited,omitempty"           `
	RightLeg         *string `                       json:"right_leg,omitempty"                  msg:"right_leg,omitempty"                bson:"right_leg,omitempty"                 `
	RightLegLimited  *string `                       json:"right_leg_limited,omitempty"          msg:"right_leg_limited,omitempty"        bson:"right_leg_limited,omitempty"         `
	LeftLeg          *string `                       json:"left_leg,omitempty"                   msg:"left_leg,omitempty"                 bson:"left_leg,omitempty"                  `
	LeftLegLimited   *string `                       json:"left_leg_limited,omitempty"           msg:"left_leg_limited,omitempty"         bson:"left_leg_limited,omitempty"          `
	RightFoot        *string `                       json:"right_foot,omitempty"                 msg:"right_foot,omitempty"               bson:"right_foot,omitempty"                `
	RightFootLimited *string `                       json:"right_foot_limited,omitempty"         msg:"right_foot_limited,omitempty"       bson:"right_foot_limited,omitempty"        `
	LeftFoot         *string `                       json:"left_foot,omitempty"                  msg:"left_foot,omitempty"                bson:"left_foot,omitempty"                 `
	LeftFootLimited  *string `                       json:"left_foot_limited,omitempty"          msg:"left_foot_limited,omitempty"        bson:"left_foot_limited,omitempty"         `
	RightArm         *string `                       json:"right_arm,omitempty"                  msg:"right_arm,omitempty"                bson:"right_arm,omitempty"                 `
	RightArmLimited  *string `                       json:"right_arm_limited,omitempty"          msg:"right_arm_limited,omitempty"        bson:"right_arm_limited,omitempty"         `
	LeftArm          *string `                       json:"left_arm,omitempty"                   msg:"left_arm,omitempty"                 bson:"left_arm,omitempty"                  `
	LeftArmLimited   *string `                       json:"left_arm_limited,omitempty"           msg:"left_arm_limited,omitempty"         bson:"left_arm_limited,omitempty"          `
	RightHand        *string `                       json:"right_hand,omitempty"                 msg:"right_hand,omitempty"               bson:"right_hand,omitempty"                `
	RightHandLimited *string `                       json:"right_hand_limited,omitempty"         msg:"right_hand_limited,omitempty"       bson:"right_hand_limited,omitempty"        `
	LeftHand         *string `                       json:"left_hand,omitempty"                  msg:"left_hand,omitempty"                bson:"left_hand,omitempty"                 `
	LeftHandLimited  *string `                       json:"left_hand_limited,omitempty"          msg:"left_hand_limited,omitempty"        bson:"left_hand_limited,omitempty"         `
}
