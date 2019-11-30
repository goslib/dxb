package dxb

func QueryLimit(option *FindOptions, limit int64) *FindOptions {
	if option == nil {
		option = &FindOptions{}
	}
	option.SetLimit(limit)
	return option
}

func QuerySortBy(option *FindOptions, key string, desc bool) *FindOptions {
	if option == nil {
		option = &FindOptions{}
	}
	if desc {
		option.Sort = M{key: -1}
	} else {
		option.Sort = M{key: 1}
	}
	return option
}
