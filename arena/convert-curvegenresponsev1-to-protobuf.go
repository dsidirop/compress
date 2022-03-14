package arena

func ConvertCurvegenResponseV1ToProtobuf(x *CurveGenReplyV1) *PBCurveGenReplyV1 {
	return &PBCurveGenReplyV1{
		//state           protoimpl.MessageState    // no need
		//sizeCache       protoimpl.SizeCache       // no need
		//unknownFields   protoimpl.UnknownFields   // no need

		ClientId: x.ClientId,
		Spec:     protobufConvertCurveSpecificationV1(x.Spec),
		Lead1:    protobufConvertIntArrayToInt64Array(x.Lead1),
		Lead2:    protobufConvertIntArrayToInt64Array(x.Lead2),
		Lead3:    protobufConvertIntArrayToInt64Array(x.Lead3),
		Lead4:    protobufConvertIntArrayToInt64Array(x.Lead4),
		Lead5:    protobufConvertIntArrayToInt64Array(x.Lead5),
		Lead6:    protobufConvertIntArrayToInt64Array(x.Lead6),
		Lead7:    protobufConvertIntArrayToInt64Array(x.Lead7),
		Lead8:    protobufConvertIntArrayToInt64Array(x.Lead8),
		Lead9:    protobufConvertIntArrayToInt64Array(x.Lead9),
		Lead10:   protobufConvertIntArrayToInt64Array(x.Lead10),
		Lead11:   protobufConvertIntArrayToInt64Array(x.Lead11),
		Lead12:   protobufConvertIntArrayToInt64Array(x.Lead12),
		Abp:      protobufConvertIntArrayToInt64Array(x.Abp),
		Cvp:      protobufConvertIntArrayToInt64Array(x.Cvp),
		Pap:      protobufConvertIntArrayToInt64Array(x.Pap),
		Spo2:     protobufConvertIntArrayToInt64Array(x.Spo2),
		Wp:       protobufConvertIntArrayToInt64Array(x.Wp),
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
