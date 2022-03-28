package arena

import "github.com/klauspost/compress/arena/thvitalstemplate"

func ConvertVitalsTemplateToThrift(x *VitalsTemplate) *thvitalstemplate.THVitalsTemplate {
	return &thvitalstemplate.THVitalsTemplate{
		Spec: thriftConvertVitalsTemplateSpec(x.Spec),
	}
}

func thriftConvertVitalsTemplateSpec(spec VitalsTemplateSpec) *thvitalstemplate.THVitalsTemplateSpec {
	results := &thvitalstemplate.THVitalsTemplateSpec{}

	results.RespiratoryRate = spec.RespiratoryRate
	results.SpO2 = spec.SpO2
	results.BloodPressure = thriftConvertBloodPressure(spec.BloodPressure)
	results.EtCO2 = spec.EtCO2
	results.TemperatureBlood = spec.TemperatureBlood
	results.TemperaturePeripheral = spec.TemperaturePeripheral
	results.Ecg = thriftConvertEcg(spec.Ecg)
	results.Eyes = thriftConvertEyes(spec.Eyes)
	results.AirwayObstruction = thriftConvertAirwayObstruction(spec.AirwayObstruction)
	results.LungSounds = thriftConvertLungSounds(spec.LungSounds)
	results.HeartSounds = thriftConvertHeartSounds(spec.HeartSounds)
	results.BowelSounds = thriftConvertBowelSounds(spec.BowelSounds)
	results.Pulses = thriftConvertPulses(spec.Pulses)

	return results
}

func thriftConvertPulses(pulses *Pulses) *thvitalstemplate.THPulses {
	if pulses == nil {
		return nil
	}

	results := &thvitalstemplate.THPulses{}
	results.Central = pulses.Central
	results.CentralLimited = pulses.CentralLimited
	results.RightLeg = pulses.RightLeg
	results.RightLegLimited = pulses.RightLegLimited
	results.LeftLeg = pulses.LeftLeg
	results.LeftLegLimited = pulses.LeftLegLimited
	results.RightFoot = pulses.RightFoot
	results.RightFootLimited = pulses.RightFootLimited
	results.LeftFoot = pulses.LeftFoot
	results.LeftFootLimited = pulses.LeftFootLimited
	results.RightArm = pulses.RightArm
	results.RightArmLimited = pulses.RightArmLimited
	results.LeftArm = pulses.LeftArm
	results.LeftArmLimited = pulses.LeftArmLimited
	results.RightHand = pulses.RightHand
	results.RightHandLimited = pulses.RightHandLimited
	results.LeftHand = pulses.LeftHand
	results.LeftHandLimited = pulses.LeftHandLimited

	return results
}

func thriftConvertBowelSounds(bowelSounds *BowelSounds) *thvitalstemplate.THBowelSounds {
	if bowelSounds == nil {
		return nil
	}

	results := &thvitalstemplate.THBowelSounds{}
	results.Bowel = bowelSounds.Bowel
	results.BowelVolume = bowelSounds.BowelVolume

	return results
}

func thriftConvertLungSounds(lungSounds *LungSounds) *thvitalstemplate.THLungSounds {
	if lungSounds == nil {
		return nil
	}

	results := &thvitalstemplate.THLungSounds{}
	results.Left = thriftConvertSingleLungSounds(lungSounds.Left)
	results.Right = thriftConvertSingleLungSounds(lungSounds.Right)

	return results
}

func thriftConvertSingleLungSounds(singleLungSounds *SingleLungSounds) *thvitalstemplate.THSingleLungSounds {
	if singleLungSounds == nil {
		return nil
	}

	results := &thvitalstemplate.THSingleLungSounds{}
	results.AnteriorUpper = singleLungSounds.AnteriorUpper
	results.AnteriorUpperVolume = singleLungSounds.AnteriorUpperVolume
	results.AnteriorLower = singleLungSounds.AnteriorLower
	results.AnteriorLowerVolume = singleLungSounds.AnteriorLowerVolume
	results.AnteriorMiddle = singleLungSounds.AnteriorMiddle
	results.AnteriorMiddleVolume = singleLungSounds.AnteriorMiddleVolume
	results.PosteriorUpper = singleLungSounds.PosteriorUpper
	results.PosteriorUpperVolume = singleLungSounds.PosteriorUpperVolume
	results.PosteriorLower = singleLungSounds.PosteriorLower
	results.PosteriorLowerVolume = singleLungSounds.PosteriorLowerVolume

	return results
}

func thriftConvertHeartSounds(heartSounds *HeartSounds) *thvitalstemplate.THHeartSounds {
	if heartSounds == nil {
		return nil
	}

	results := &thvitalstemplate.THHeartSounds{}
	results.Aortic = heartSounds.Aortic
	results.Mitral = heartSounds.Mitral
	results.Pulmonary = heartSounds.Pulmonary
	results.Tricuspid = heartSounds.Tricuspid
	results.MitralVolume = heartSounds.MitralVolume
	results.AorticVolume = heartSounds.AorticVolume
	results.PulmonaryVolume = heartSounds.PulmonaryVolume
	results.TricuspidVolume = heartSounds.TricuspidVolume

	return results
}

func thriftConvertAirwayObstruction(airwayObstruction *AirwayObstruction) *thvitalstemplate.THAirwayObstruction {
	if airwayObstruction == nil {
		return nil
	}

	results := &thvitalstemplate.THAirwayObstruction{}
	results.TongueEdema = airwayObstruction.TongueEdema

	return results
}

func thriftConvertEyes(eyes *Eyes) *thvitalstemplate.THEyes {
	if eyes == nil {
		return nil
	}

	results := &thvitalstemplate.THEyes{}

	return results
}

func thriftConvertBloodPressure(bloodPressure *BloodPressure) *thvitalstemplate.THBloodPressure {
	if bloodPressure == nil {
		return nil
	}

	results := &thvitalstemplate.THBloodPressure{}
	results.Systolic = bloodPressure.Systolic
	results.Diastolic = bloodPressure.Diastolic

	return results
}

func thriftConvertEcg(ecg *Ecg) *thvitalstemplate.THEcg {
	if ecg == nil {
		return nil
	}

	results := &thvitalstemplate.THEcg{}
	results.EmdPea = ecg.EmdPea
	results.Severity = ecg.Severity
	results.HeartRate = ecg.HeartRate
	results.AllowPacing = ecg.AllowPacing
	results.BasicRhythm = ecg.BasicRhythm
	results.Extrasystole = thriftConvertExtraSystole(ecg.Extrasystole)

	return results
}

func thriftConvertExtraSystole(extrasystole *Extrasystole) *thvitalstemplate.THExtrasystole {
	if extrasystole == nil {
		return nil
	}

	results := &thvitalstemplate.THExtrasystole{}
	results.Type = extrasystole.Type
	results.Probability = extrasystole.Probability

	return results
}
