package arena

import (
	"github.com/klauspost/compress/arena/thcurvegenresponsev1"
)

func ConvertCurvegenResponseV1ToThrift(x *CurveGenReplyV1) *thcurvegenresponsev1.THCurveGenReplyV1 {
	y := thcurvegenresponsev1.NewTHCurveGenReplyV1()

	y.ClientId = x.ClientId

	y.Spec = thriftConvertCurveSpecificationV1(x.Spec)

	y.Lead1 = thriftConvertIntArrayToInt64Array(x.Lead1)
	y.Lead2 = thriftConvertIntArrayToInt64Array(x.Lead2)
	y.Lead3 = thriftConvertIntArrayToInt64Array(x.Lead3)
	y.Lead4 = thriftConvertIntArrayToInt64Array(x.Lead4)
	y.Lead5 = thriftConvertIntArrayToInt64Array(x.Lead5)
	y.Lead6 = thriftConvertIntArrayToInt64Array(x.Lead6)
	y.Lead7 = thriftConvertIntArrayToInt64Array(x.Lead7)
	y.Lead8 = thriftConvertIntArrayToInt64Array(x.Lead8)
	y.Lead9 = thriftConvertIntArrayToInt64Array(x.Lead9)
	y.Lead10 = thriftConvertIntArrayToInt64Array(x.Lead10)
	y.Lead11 = thriftConvertIntArrayToInt64Array(x.Lead11)
	y.Lead12 = thriftConvertIntArrayToInt64Array(x.Lead12)

	y.Wp = thriftConvertIntArrayToInt64Array(x.Wp)
	y.Abp = thriftConvertIntArrayToInt64Array(x.Abp)
	y.Cvp = thriftConvertIntArrayToInt64Array(x.Cvp)
	y.Pap = thriftConvertIntArrayToInt64Array(x.Pap)
	y.Spo2 = thriftConvertIntArrayToInt64Array(x.Spo2)

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
