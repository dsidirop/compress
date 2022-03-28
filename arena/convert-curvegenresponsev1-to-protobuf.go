package arena

func ConvertCurvegenResponseV1ToProtobuf(x *CurveGenReplyV1) *PBCurveGenReplyV1 {
	return &PBCurveGenReplyV1{
		//state           protoimpl.MessageState    // no need
		//sizeCache       protoimpl.SizeCache       // no need
		//unknownFields   protoimpl.UnknownFields   // no need

		ClientId: x.ClientId,
		Spec:     protobufConvertCurveSpecificationV1(x.Spec),
		Lead1:    x.Lead1,
		Lead2:    x.Lead2,
		Lead3:    x.Lead3,
		Lead4:    x.Lead4,
		Lead5:    x.Lead5,
		Lead6:    x.Lead6,
		Lead7:    x.Lead7,
		Lead8:    x.Lead8,
		Lead9:    x.Lead9,
		Lead10:   x.Lead10,
		Lead11:   x.Lead11,
		Lead12:   x.Lead12,
		Abp:      x.Abp,
		Cvp:      x.Cvp,
		Pap:      x.Pap,
		Spo2:     x.Spo2,
		Wp:       x.Wp,
		Tags:     protobufConvertTags(x.Tags),
	}
}

func protobufConvertCurveSpecificationV1(spec CurveSpecificationV1) *PBCurveSpecificationV1 {

	results := &PBCurveSpecificationV1{}
	results.Tenant = spec.Tenant
	results.EndTime = spec.EndTime
	results.StartTime = spec.StartTime
	results.CurveTypes = int64(spec.CurveTypes)
	results.SimulatorName = spec.SimulatorName
	results.SampleInterval = int64(spec.SampleInterval)

	return results
}

func protobufConvertIntArrayToInt64Array(array []int) []int64 {
	if array == nil {
		return nil
	}

	results := make([]int64, 0)
	for i := 0; i < len(array); i++ {
		results = append(results, int64(array[i]))
	}

	return results
}

func protobufConvertTags(array []Tag) []*PBTag {
	if array == nil {
		return nil
	}

	results := make([]*PBTag, 0)
	for i := 0; i < len(array); i++ {

		thtag := &PBTag{}
		thtag.TagTime = array[i].TagTime
		thtag.TagType = int64(array[i].TagType)

		results = append(results, thtag)
	}

	return results
}
