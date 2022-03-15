package arena

import (
	"github.com/klauspost/compress/arena/thcurvegenresponsev1"
)

func ConvertCurvegenResponseV1ToThrift(x *CurveGenReplyV1) *thcurvegenresponsev1.THCurveGenReplyV1 {
	y := thcurvegenresponsev1.NewTHCurveGenReplyV1()

	y.ClientId = x.ClientId

	y.Spec = thriftConvertCurveSpecificationV1(x.Spec)

	y.Lead1 = x.Lead1
	y.Lead2 = x.Lead2
	y.Lead3 = x.Lead3
	y.Lead4 = x.Lead4
	y.Lead5 = x.Lead5
	y.Lead6 = x.Lead6
	y.Lead7 = x.Lead7
	y.Lead8 = x.Lead8
	y.Lead9 = x.Lead9
	y.Lead10 = x.Lead10
	y.Lead11 = x.Lead11
	y.Lead12 = x.Lead12

	y.Wp = x.Wp
	y.Abp = x.Abp
	y.Cvp = x.Cvp
	y.Pap = x.Pap
	y.Spo2 = x.Spo2

	y.Tags = thriftConvertTags(x.Tags)

	return y
}

func thriftConvertIntArrayToInt64Array(array []int) []int64 {
	if array == nil {
		return nil
	}

	results := make([]int64, 0)
	for i := 0; i < len(array); i++ {
		results = append(results, int64(array[i]))
	}

	return results
}

func thriftConvertTags(array []Tag) []*thcurvegenresponsev1.THTag {
	if array == nil {
		return nil
	}

	results := make([]*thcurvegenresponsev1.THTag, 0)
	for i := 0; i < len(array); i++ {

		thtag := thcurvegenresponsev1.NewTHTag()
		thtag.TagTime = array[i].TagTime
		thtag.TagType = int64(array[i].TagType)

		results = append(results, thtag)
	}

	return results
}

func thriftConvertCurveSpecificationV1(spec CurveSpecificationV1) *thcurvegenresponsev1.THCurveSpecificationV1 {

	results := &thcurvegenresponsev1.THCurveSpecificationV1{}
	results.Tenant = spec.Tenant
	results.EndTime = spec.EndTime
	results.StartTime = spec.StartTime
	results.CurveTypes = int64(spec.CurveTypes)
	results.SimulatorName = spec.SimulatorName
	results.SampleInterval = int64(spec.SampleInterval)

	return results
}
