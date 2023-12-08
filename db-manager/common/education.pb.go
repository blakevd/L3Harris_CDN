// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/education.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EducationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country                             string  `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Latitude                            float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude                           float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	OosrPre0PrimaryAgeMale              int32   `protobuf:"varint,4,opt,name=oosr_pre0primary_age_male,json=oosrPre0primaryAgeMale,proto3" json:"oosr_pre0primary_age_male,omitempty"`
	OosrPre0PrimaryAgeFemale            int32   `protobuf:"varint,5,opt,name=oosr_pre0primary_age_female,json=oosrPre0primaryAgeFemale,proto3" json:"oosr_pre0primary_age_female,omitempty"`
	OosrPrimaryAgeMale                  int32   `protobuf:"varint,6,opt,name=oosr_primary_age_male,json=oosrPrimaryAgeMale,proto3" json:"oosr_primary_age_male,omitempty"`
	OosrPrimaryAgeFemale                int32   `protobuf:"varint,7,opt,name=oosr_primary_age_female,json=oosrPrimaryAgeFemale,proto3" json:"oosr_primary_age_female,omitempty"`
	OosrLowerSecondaryAgeMale           int32   `protobuf:"varint,8,opt,name=oosr_lower_secondary_age_male,json=oosrLowerSecondaryAgeMale,proto3" json:"oosr_lower_secondary_age_male,omitempty"`
	OosrLowerSecondaryAgeFemale         int32   `protobuf:"varint,9,opt,name=oosr_lower_secondary_age_female,json=oosrLowerSecondaryAgeFemale,proto3" json:"oosr_lower_secondary_age_female,omitempty"`
	OosrUpperSecondaryAgeMale           int32   `protobuf:"varint,10,opt,name=oosr_upper_secondary_age_male,json=oosrUpperSecondaryAgeMale,proto3" json:"oosr_upper_secondary_age_male,omitempty"`
	OosrUpperSecondaryAgeFemale         int32   `protobuf:"varint,11,opt,name=oosr_upper_secondary_age_female,json=oosrUpperSecondaryAgeFemale,proto3" json:"oosr_upper_secondary_age_female,omitempty"`
	CompletionRatePrimaryMale           int32   `protobuf:"varint,12,opt,name=completion_rate_primary_male,json=completionRatePrimaryMale,proto3" json:"completion_rate_primary_male,omitempty"`
	CompletionRatePrimaryFemale         int32   `protobuf:"varint,13,opt,name=completion_rate_primary_female,json=completionRatePrimaryFemale,proto3" json:"completion_rate_primary_female,omitempty"`
	CompletionRateLowerSecondaryMale    int32   `protobuf:"varint,14,opt,name=completion_rate_lower_secondary_male,json=completionRateLowerSecondaryMale,proto3" json:"completion_rate_lower_secondary_male,omitempty"`
	CompletionRateLowerSecondaryFemale  int32   `protobuf:"varint,15,opt,name=completion_rate_lower_secondary_female,json=completionRateLowerSecondaryFemale,proto3" json:"completion_rate_lower_secondary_female,omitempty"`
	CompletionRateUpperSecondaryMale    int32   `protobuf:"varint,16,opt,name=completion_rate_upper_secondary_male,json=completionRateUpperSecondaryMale,proto3" json:"completion_rate_upper_secondary_male,omitempty"`
	CompletionRateUpperSecondaryFemale  int32   `protobuf:"varint,17,opt,name=completion_rate_upper_secondary_female,json=completionRateUpperSecondaryFemale,proto3" json:"completion_rate_upper_secondary_female,omitempty"`
	Grade_2_3ProficiencyReading         int32   `protobuf:"varint,18,opt,name=grade_2_3_proficiency_reading,json=grade23ProficiencyReading,proto3" json:"grade_2_3_proficiency_reading,omitempty"`
	Grade_2_3ProficiencyMath            int32   `protobuf:"varint,19,opt,name=grade_2_3_proficiency_math,json=grade23ProficiencyMath,proto3" json:"grade_2_3_proficiency_math,omitempty"`
	PrimaryEndProficiencyReading        int32   `protobuf:"varint,20,opt,name=primary_end_proficiency_reading,json=primaryEndProficiencyReading,proto3" json:"primary_end_proficiency_reading,omitempty"`
	PrimaryEndProficiencyMath           int32   `protobuf:"varint,21,opt,name=primary_end_proficiency_math,json=primaryEndProficiencyMath,proto3" json:"primary_end_proficiency_math,omitempty"`
	LowerSecondaryEndProficiencyReading int32   `protobuf:"varint,22,opt,name=lower_secondary_end_proficiency_reading,json=lowerSecondaryEndProficiencyReading,proto3" json:"lower_secondary_end_proficiency_reading,omitempty"`
	LowerSecondaryEndProficiencyMath    int32   `protobuf:"varint,23,opt,name=lower_secondary_end_proficiency_math,json=lowerSecondaryEndProficiencyMath,proto3" json:"lower_secondary_end_proficiency_math,omitempty"`
	Youth_15_24LiteracyRateMale         int32   `protobuf:"varint,24,opt,name=youth_15_24_literacy_rate_male,json=youth1524LiteracyRateMale,proto3" json:"youth_15_24_literacy_rate_male,omitempty"`
	Youth_15_24LiteracyRateFemale       int32   `protobuf:"varint,25,opt,name=youth_15_24_literacy_rate_female,json=youth1524LiteracyRateFemale,proto3" json:"youth_15_24_literacy_rate_female,omitempty"`
	BirthRate                           float64 `protobuf:"fixed64,26,opt,name=birth_rate,json=birthRate,proto3" json:"birth_rate,omitempty"`
	GrossPrimaryEducationEnrollment     float64 `protobuf:"fixed64,27,opt,name=gross_primary_education_enrollment,json=grossPrimaryEducationEnrollment,proto3" json:"gross_primary_education_enrollment,omitempty"`
	GrossTertiaryEducationEnrollment    float64 `protobuf:"fixed64,28,opt,name=gross_tertiary_education_enrollment,json=grossTertiaryEducationEnrollment,proto3" json:"gross_tertiary_education_enrollment,omitempty"`
	UnemploymentRate                    float64 `protobuf:"fixed64,29,opt,name=unemployment_rate,json=unemploymentRate,proto3" json:"unemployment_rate,omitempty"`
}

func (x *EducationData) Reset() {
	*x = EducationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_education_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EducationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EducationData) ProtoMessage() {}

func (x *EducationData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_education_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EducationData.ProtoReflect.Descriptor instead.
func (*EducationData) Descriptor() ([]byte, []int) {
	return file_proto_education_proto_rawDescGZIP(), []int{0}
}

func (x *EducationData) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *EducationData) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *EducationData) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *EducationData) GetOosrPre0PrimaryAgeMale() int32 {
	if x != nil {
		return x.OosrPre0PrimaryAgeMale
	}
	return 0
}

func (x *EducationData) GetOosrPre0PrimaryAgeFemale() int32 {
	if x != nil {
		return x.OosrPre0PrimaryAgeFemale
	}
	return 0
}

func (x *EducationData) GetOosrPrimaryAgeMale() int32 {
	if x != nil {
		return x.OosrPrimaryAgeMale
	}
	return 0
}

func (x *EducationData) GetOosrPrimaryAgeFemale() int32 {
	if x != nil {
		return x.OosrPrimaryAgeFemale
	}
	return 0
}

func (x *EducationData) GetOosrLowerSecondaryAgeMale() int32 {
	if x != nil {
		return x.OosrLowerSecondaryAgeMale
	}
	return 0
}

func (x *EducationData) GetOosrLowerSecondaryAgeFemale() int32 {
	if x != nil {
		return x.OosrLowerSecondaryAgeFemale
	}
	return 0
}

func (x *EducationData) GetOosrUpperSecondaryAgeMale() int32 {
	if x != nil {
		return x.OosrUpperSecondaryAgeMale
	}
	return 0
}

func (x *EducationData) GetOosrUpperSecondaryAgeFemale() int32 {
	if x != nil {
		return x.OosrUpperSecondaryAgeFemale
	}
	return 0
}

func (x *EducationData) GetCompletionRatePrimaryMale() int32 {
	if x != nil {
		return x.CompletionRatePrimaryMale
	}
	return 0
}

func (x *EducationData) GetCompletionRatePrimaryFemale() int32 {
	if x != nil {
		return x.CompletionRatePrimaryFemale
	}
	return 0
}

func (x *EducationData) GetCompletionRateLowerSecondaryMale() int32 {
	if x != nil {
		return x.CompletionRateLowerSecondaryMale
	}
	return 0
}

func (x *EducationData) GetCompletionRateLowerSecondaryFemale() int32 {
	if x != nil {
		return x.CompletionRateLowerSecondaryFemale
	}
	return 0
}

func (x *EducationData) GetCompletionRateUpperSecondaryMale() int32 {
	if x != nil {
		return x.CompletionRateUpperSecondaryMale
	}
	return 0
}

func (x *EducationData) GetCompletionRateUpperSecondaryFemale() int32 {
	if x != nil {
		return x.CompletionRateUpperSecondaryFemale
	}
	return 0
}

func (x *EducationData) GetGrade_2_3ProficiencyReading() int32 {
	if x != nil {
		return x.Grade_2_3ProficiencyReading
	}
	return 0
}

func (x *EducationData) GetGrade_2_3ProficiencyMath() int32 {
	if x != nil {
		return x.Grade_2_3ProficiencyMath
	}
	return 0
}

func (x *EducationData) GetPrimaryEndProficiencyReading() int32 {
	if x != nil {
		return x.PrimaryEndProficiencyReading
	}
	return 0
}

func (x *EducationData) GetPrimaryEndProficiencyMath() int32 {
	if x != nil {
		return x.PrimaryEndProficiencyMath
	}
	return 0
}

func (x *EducationData) GetLowerSecondaryEndProficiencyReading() int32 {
	if x != nil {
		return x.LowerSecondaryEndProficiencyReading
	}
	return 0
}

func (x *EducationData) GetLowerSecondaryEndProficiencyMath() int32 {
	if x != nil {
		return x.LowerSecondaryEndProficiencyMath
	}
	return 0
}

func (x *EducationData) GetYouth_15_24LiteracyRateMale() int32 {
	if x != nil {
		return x.Youth_15_24LiteracyRateMale
	}
	return 0
}

func (x *EducationData) GetYouth_15_24LiteracyRateFemale() int32 {
	if x != nil {
		return x.Youth_15_24LiteracyRateFemale
	}
	return 0
}

func (x *EducationData) GetBirthRate() float64 {
	if x != nil {
		return x.BirthRate
	}
	return 0
}

func (x *EducationData) GetGrossPrimaryEducationEnrollment() float64 {
	if x != nil {
		return x.GrossPrimaryEducationEnrollment
	}
	return 0
}

func (x *EducationData) GetGrossTertiaryEducationEnrollment() float64 {
	if x != nil {
		return x.GrossTertiaryEducationEnrollment
	}
	return 0
}

func (x *EducationData) GetUnemploymentRate() float64 {
	if x != nil {
		return x.UnemploymentRate
	}
	return 0
}

var File_proto_education_proto protoreflect.FileDescriptor

var file_proto_education_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x0e, 0x0a, 0x0d, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x39, 0x0a,
	0x19, 0x6f, 0x6f, 0x73, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x30, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x16, 0x6f, 0x6f, 0x73, 0x72, 0x50, 0x72, 0x65, 0x30, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x41, 0x67, 0x65, 0x4d, 0x61, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x6f, 0x73, 0x72,
	0x5f, 0x70, 0x72, 0x65, 0x30, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x67, 0x65,
	0x5f, 0x66, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6f,
	0x6f, 0x73, 0x72, 0x50, 0x72, 0x65, 0x30, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x67,
	0x65, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x6f, 0x6f, 0x73, 0x72, 0x5f,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6f, 0x6f, 0x73, 0x72, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x41, 0x67, 0x65, 0x4d, 0x61, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x6f, 0x6f,
	0x73, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x66,
	0x65, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6f, 0x6f, 0x73,
	0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x67, 0x65, 0x46, 0x65, 0x6d, 0x61, 0x6c,
	0x65, 0x12, 0x40, 0x0a, 0x1d, 0x6f, 0x6f, 0x73, 0x72, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61,
	0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x6f, 0x6f, 0x73, 0x72, 0x4c, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x41, 0x67, 0x65, 0x4d,
	0x61, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x1f, 0x6f, 0x6f, 0x73, 0x72, 0x5f, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x67, 0x65, 0x5f,
	0x66, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x6f, 0x6f,
	0x73, 0x72, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x41, 0x67, 0x65, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x1d, 0x6f, 0x6f, 0x73,
	0x72, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x19, 0x6f, 0x6f, 0x73, 0x72, 0x55, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x41, 0x67, 0x65, 0x4d, 0x61, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x1f, 0x6f,
	0x6f, 0x73, 0x72, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x6f, 0x6f, 0x73, 0x72, 0x55, 0x70, 0x70, 0x65, 0x72, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x41, 0x67, 0x65, 0x46, 0x65, 0x6d, 0x61, 0x6c,
	0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6c,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x61,
	0x6c, 0x65, 0x12, 0x43, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x65,
	0x6d, 0x61, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6c, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x4d, 0x61, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x65, 0x6d, 0x61, 0x6c,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x22, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x61, 0x72, 0x79, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x24, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x6d,
	0x61, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x55, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4d, 0x61, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x26, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x66,
	0x65, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x22, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x55, 0x70, 0x70, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x12,
	0x40, 0x0a, 0x1d, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x32, 0x5f, 0x33, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x67, 0x72, 0x61, 0x64, 0x65, 0x32, 0x33, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x3a, 0x0a, 0x1a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x32, 0x5f, 0x33, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x68, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x67, 0x72, 0x61, 0x64, 0x65, 0x32, 0x33, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x61, 0x74, 0x68, 0x12, 0x45, 0x0a,
	0x1f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x45,
	0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x1c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x6d, 0x61, 0x74, 0x68, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63,
	0x79, 0x4d, 0x61, 0x74, 0x68, 0x12, 0x54, 0x0a, 0x27, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x23, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x4e, 0x0a, 0x24, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d,
	0x61, 0x74, 0x68, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x61, 0x74, 0x68, 0x12, 0x41, 0x0a, 0x1e, 0x79,
	0x6f, 0x75, 0x74, 0x68, 0x5f, 0x31, 0x35, 0x5f, 0x32, 0x34, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x19, 0x79, 0x6f, 0x75, 0x74, 0x68, 0x31, 0x35, 0x32, 0x34, 0x4c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x79, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6c, 0x65, 0x12, 0x45,
	0x0a, 0x20, 0x79, 0x6f, 0x75, 0x74, 0x68, 0x5f, 0x31, 0x35, 0x5f, 0x32, 0x34, 0x5f, 0x6c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x65, 0x6d, 0x61,
	0x6c, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x79, 0x6f, 0x75, 0x74, 0x68, 0x31,
	0x35, 0x32, 0x34, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x63, 0x79, 0x52, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x6d, 0x61, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x22, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x1f, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x45, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x4d, 0x0a, 0x23, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x74, 0x65, 0x72, 0x74, 0x69,
	0x61, 0x72, 0x79, 0x5f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x20,
	0x67, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x65, 0x72, 0x74, 0x69, 0x61, 0x72, 0x79, 0x45, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x75, 0x6e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x75, 0x6e, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_education_proto_rawDescOnce sync.Once
	file_proto_education_proto_rawDescData = file_proto_education_proto_rawDesc
)

func file_proto_education_proto_rawDescGZIP() []byte {
	file_proto_education_proto_rawDescOnce.Do(func() {
		file_proto_education_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_education_proto_rawDescData)
	})
	return file_proto_education_proto_rawDescData
}

var file_proto_education_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_education_proto_goTypes = []interface{}{
	(*EducationData)(nil), // 0: EducationData
}
var file_proto_education_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_education_proto_init() }
func file_proto_education_proto_init() {
	if File_proto_education_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_education_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EducationData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_education_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_education_proto_goTypes,
		DependencyIndexes: file_proto_education_proto_depIdxs,
		MessageInfos:      file_proto_education_proto_msgTypes,
	}.Build()
	File_proto_education_proto = out.File
	file_proto_education_proto_rawDesc = nil
	file_proto_education_proto_goTypes = nil
	file_proto_education_proto_depIdxs = nil
}
