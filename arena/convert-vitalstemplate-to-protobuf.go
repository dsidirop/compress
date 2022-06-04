package arena

func ConvertVitalsTemplateToProtobuf(x *VitalsTemplate) *PBVitalsTemplate {
	return &PBVitalsTemplate{
		Spec: protobufConvertVitalsTemplateSpec(x.Spec),
	}
}

func protobufConvertVitalsTemplateSpec(spec VitalsTemplateSpec) *PBVitalsTemplateSpec {
	results := &PBVitalsTemplateSpec{}

	results.RespiratoryRate = spec.RespiratoryRate
	results.SpO2 = spec.SpO2
	results.BloodPressure = protobufConvertBloodPressure(spec.BloodPressure)
	results.EtCO2 = spec.EtCO2
	results.TemperatureBlood = spec.TemperatureBlood
	results.TemperaturePeripheral = spec.TemperaturePeripheral
	results.Ecg = protobufConvertEcg(spec.Ecg)
	results.Eyes = protobufConvertEyes(spec.Eyes)
	results.AirwayObstruction = protobufConvertAirwayObstruction(spec.AirwayObstruction)
	results.LungSounds = protobufConvertLungSounds(spec.LungSounds)
	results.HeartSounds = protobufConvertHeartSounds(spec.HeartSounds)
	results.BowelSounds = protobufConvertBowelSounds(spec.BowelSounds)
	results.Pulses = protobufConvertPulses(spec.Pulses)

	return results
}

func protobufConvertPulses(pulses *Pulses) *PBPulses {
	if pulses == nil {
		return nil
	}

	results := &PBPulses{}
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

func protobufConvertBowelSounds(bowelSounds *BowelSounds) *PBBowelSounds {
	if bowelSounds == nil {
		return nil
	}

	results := &PBBowelSounds{}
	results.Bowel = bowelSounds.Bowel
	results.BowelVolume = bowelSounds.BowelVolume

	return results
}

func protobufConvertLungSounds(lungSounds *LungSounds) *PBLungSounds {
	if lungSounds == nil {
		return nil
	}

	results := &PBLungSounds{}
	results.Left = protobufConvertSingleLungSounds(lungSounds.Left)
	results.Right = protobufConvertSingleLungSounds(lungSounds.Right)

	return results
}

func protobufConvertSingleLungSounds(singleLungSounds *SingleLungSounds) *PBSingleLungSounds {
	if singleLungSounds == nil {
		return nil
	}

	results := &PBSingleLungSounds{}
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

func protobufConvertHeartSounds(heartSounds *HeartSounds) *PBHeartSounds {
	if heartSounds == nil {
		return nil
	}

	results := &PBHeartSounds{}
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

func protobufConvertAirwayObstruction(airwayObstruction *AirwayObstruction) *PBAirwayObstruction {
	if airwayObstruction == nil {
		return nil
	}

	results := &PBAirwayObstruction{}
	results.TongueEdema = airwayObstruction.TongueEdema

	return results
}

func protobufConvertEyes(eyes *Eyes) *PBEyes {
	if eyes == nil {
		return nil
	}

	results := &PBEyes{}

	return results
}

func protobufConvertBloodPressure(bloodPressure *BloodPressure) *PBBloodPressure {
	if bloodPressure == nil {
		return nil
	}

	results := &PBBloodPressure{}
	results.Systolic = bloodPressure.Systolic
	results.Diastolic = bloodPressure.Diastolic

	return results
}

func protobufConvertEcg(ecg *Ecg) *PBEcg {
	if ecg == nil {
		return nil
	}

	results := &PBEcg{}
	results.EmdPea = ecg.EmdPea
	results.Severity = ecg.Severity
	results.HeartRate = ecg.HeartRate
	results.AllowPacing = ecg.AllowPacing
	results.BasicRhythm = ecg.BasicRhythm
	results.Extrasystole = protobufConvertExtraSystole(ecg.Extrasystole)

	return results
}

func protobufConvertExtraSystole(extrasystole *Extrasystole) *PBExtrasystole {
	if extrasystole == nil {
		return nil
	}

	results := &PBExtrasystole{}
	results.Type = extrasystole.Type
	results.Probability = extrasystole.Probability

	return results
}
