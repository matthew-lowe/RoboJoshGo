package framework

func TagToUserId(tag string) string {
	return tag[3:][:len(tag)-4]
}
