package excel

import "strings"

type cTag struct {
	tag        string  // tag名称，如head
	param      string  // tag参数
	next       *cTag   // 下一个tag
	typeof     tagType // tag类型
	hasParam   bool    // 是否有参数
	isBlockEnd bool    // 是否是block结束
}

type tagType int

func parseFieldTags(tag string) (firstCtag, current *cTag) {
	var t string
	tags := strings.Split(tag, tagSeparator)

	for i := 0; i < len(tags); i++ {
		t = tags[i]
		parts := strings.SplitN(t, tagKeySeparator, 2)

		if i == 0 {
			current = &cTag{tag: parts[0], hasParam: len(parts) > 1}
			firstCtag = current
		} else {
			current.next = &cTag{tag: parts[0], hasParam: len(parts) > 1}
			current = current.next
		}

		if current.hasParam {
			current.param = parts[1]
		}

		switch parts[0] {
		case SubTagHeader:
			current.typeof = typeHead
		case SubTagType:
			current.typeof = typeType
		case SubTagRequired:
			current.typeof = typeRequired
		case SubTagMax:
			current.typeof = typeMax
		default:
			current.typeof = typeDefault
		}
		current.isBlockEnd = true
	}
	return
}

func getSubTagMap(tag string) map[string]*cTag {
	tagMap := make(map[string]*cTag)
	firstCtag, current := parseFieldTags(tag)

	for current = firstCtag; current != nil; current = current.next {
		tagMap[current.tag] = current
	}
	return tagMap
}
