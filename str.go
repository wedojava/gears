package gears

import (
	"fmt"
	"github.com/axgle/mahonia"
	chardet2 "github.com/chennqqi/chardet"
	mapset "github.com/deckarep/golang-set"
	"github.com/gogs/chardet"
	"log"
)

func PrintSlice(x []string) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// StrSliceDeDupl is used for string slice deduplication
func StrSliceDeDupl(items []string) []string {
	result := make([]string, 0, len(items))
	temp := map[string]struct{}{}
	for _, item := range items {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func ConvertToUtf8(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func StrDetector2(s string) string {
	// detectors := chardet2.Possible([]byte(s))
	detector := chardet2.Mostlike([]byte(s))
	return detector
}

func StrDetector(s string) string {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf(
	//         "Detected charset is %s, language is %s",
	//         result.Charset,
	//         result.Language,
	// )

	return result.Charset
}

// StrSliceDiff return strings in sl1 but not in sl2
func StrSliceDiff(sl1, sl2 []string) (ret []string) {
	for _, s1 := range sl1 {
		i, j := 0, 0
		for _, s2 := range sl2 {
			if s1 == s2 {
				j++
				continue
			} else {
				i++
			}
		}
		if j == 0 {
			ret = append(ret, s1)
		}
	}

	return
}

// StrSliceDiff2 return strings in sl1 but not in sl2
func StrSliceDiff2(sl1, sl2 []string) (ret []string) {
	s1 := mapset.NewSet()
	for _, s := range sl1 {
		s1.Add(s)
	}
	s2 := mapset.NewSet()
	for _, s := range sl2 {
		s2.Add(s)
	}
	r := s1.Difference(s2).ToSlice()
	for _, i := range r {
		ret = append(ret, fmt.Sprintf("%v", i))
	}

	return
}
