package util

import "testing"

func TestReplaceHangulPipe(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{"[현대]싼타페DM 2WD 2.0 PREMIUMㅣ2013년식ㅣ21,056 kmㅣ수원ㅣ무사고ㅣ2,220만원", "[현대]싼타페DM 2WD 2.0 PREMIUM|2013년식|21,056 km|수원|무사고|2,220만원"},
		{"[현대]제네시스DH G380 파이니스트 에디션 AWDㅣ2015년식ㅣ38,772 kmㅣ수원ㅣ무사고ㅣ4,690만원", "[현대]제네시스DH G380 파이니스트 에디션 AWD|2015년식|38,772 km|수원|무사고|4,690만원"},
	}

	for _, tc := range testCases {
		s := ReplaceHangulPipe(tc.str)
		if tc.expected != s {
			t.Error("Expected", tc.expected, "got", s)
		}
	}
}

func TestGetBrand(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{"[현대]싼타페DM 2WD 2.0 PREMIUM|2013년식|21,056 km|수원|무사고|2,220만원", "현대"},
		{"[기아]|레이프레스티지|휘발유|12년형|49,243km|오토|은하색|수원| 840만원|", "기아"},
		{"[쉐보래] 스파크LT | 2014년식| 107km|인천|무사고|730만원|인증딜러박현진", "쉐보래"},
		{"[삼성] 뉴 sm3(신형)(비흡연.여성1인신조)|2010년식|30000km|인천|무사고|500만원|인증딜러박현진", "삼성"},
	}

	for _, tc := range testCases {
		b := GetBrand(tc.str)
		if tc.expected != b {
			t.Error("Expected", tc.expected, "got", b)
		}
	}
}

func TestGetName(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{"[현대]제네시스DH G380 파이니스트 에디션 AWD|2016년식|38,335 km|수원|무사고|4,850만원", "제네시스DH G380 파이니스트 에디션 AWD"},
		{"[기아]|K31.6GDI 럭셔리|휘발유|15년형|32,517km|오토|흰색|수원| 1230만원|", "K31.6GDI 럭셔리"},
		{"[쉐보래] 스파크LT | 2014년식| 107km|인천|무사고|730만원|인증딜러박현진", "스파크LT"},
		{"[삼성] 뉴 sm3(신형)(비흡연.여성1인신조)|2010년식|30000km|인천|무사고|500만원|인증딜러박현진", "뉴 sm3(신형)(비흡연.여성1인신조)"},
	}

	for _, tc := range testCases {
		n := GetName(tc.str)
		if tc.expected != n {
			t.Error("Expected", tc.expected, "got", n)
		}
	}
}

func TestGetYear(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{"[현대]제네시스DH G380 파이니스트 에디션 AWD|2016년식|38,335 km|수원|무사고|4,850만원", "2016"},
		{"[기아]|K31.6GDI 럭셔리|휘발유|15년형|32,517km|오토|흰색|수원| 1230만원|", "15"},
		{"[쉐보래] 스파크LT | 2014년식| 107km|인천|무사고|730만원|인증딜러박현진", "2014"},
		{"[삼성] 뉴 sm3(신형)(비흡연.여성1인신조)|2010년식|30000km|인천|무사고|500만원|인증딜러박현진", "2010"},
	}

	for _, tc := range testCases {
		n := GetYear(tc.str)
		if tc.expected != n {
			t.Error("Expected", tc.expected, "got", n)
		}
	}
}

func TestGetDistance(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{"[현대]제네시스DH G380 파이니스트 에디션 AWD|2016년식|38,335 km|수원|무사고|4,850만원", "38,335"},
		{"[기아]|K31.6GDI 럭셔리|휘발유|15년형|32,517km|오토|흰색|수원| 1230만원|", "32,517"},
		{"[쉐보래] 스파크LT | 2014년식| 107km|인천|무사고|730만원|인증딜러박현진", "107"},
		{"[삼성] 뉴 sm3(신형)(비흡연.여성1인신조)|2010년식|30000km|인천|무사고|500만원|인증딜러박현진", "30000"},
	}

	for _, tc := range testCases {
		n := GetDistance(tc.str)
		if tc.expected != n {
			t.Error("Expected", tc.expected, "got", n)
		}
	}
}

func TestGetPrice(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{"[현대]제네시스DH G380 파이니스트 에디션 AWD|2016년식|38,335 km|수원|무사고|4,850만원", "4,850"},
		{"[기아]|K31.6GDI 럭셔리|휘발유|15년형|32,517km|오토|흰색|수원| 1230만원|", "1230"},
		{"[쉐보래] 스파크LT | 2014년식| 107km|인천|무사고|730만원|인증딜러박현진", "730"},
		{"[삼성] 뉴 sm3(신형)(비흡연.여성1인신조)|2010년식|30000km|인천|무사고|500만원|인증딜러박현진", "500"},
	}

	for _, tc := range testCases {
		n := GetPrice(tc.str)
		if tc.expected != n {
			t.Error("Expected", tc.expected, "got", n)
		}
	}
}
