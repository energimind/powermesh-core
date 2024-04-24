package mongoquery

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Filter is a MongoDB filter.
type Filter bson.M

// EQ adds an equality filter.
func (f Filter) EQ(key string, value any) Filter {
	f[key] = value

	return f
}

// NE adds a not equal filter.
func (f Filter) NE(key string, value any) Filter {
	f[key] = bson.M{"$ne": value}

	return f
}

// GT adds a greater than filter.
func (f Filter) GT(key string, value any) Filter {
	f[key] = bson.M{"$gt": value}

	return f
}

// GTE adds a greater than or equal filter.
func (f Filter) GTE(key string, value any) Filter {
	f[key] = bson.M{"$gte": value}

	return f
}

// LT adds a less than filter.
func (f Filter) LT(key string, value any) Filter {
	f[key] = bson.M{"$lt": value}

	return f
}

// LTE adds a less than or equal filter.
func (f Filter) LTE(key string, value any) Filter {
	f[key] = bson.M{"$lte": value}

	return f
}

// IN adds an in filter.
func (f Filter) IN(key string, values any) Filter {
	f[key] = bson.M{"$in": values}

	return f
}

// toBSON converts the filter to a BSON document.
func (f Filter) toBSON() bson.M {
	return bson.M(f)
}

func buildFilter(idOrFilter any) bson.M {
	if f, isFilter := idOrFilter.(Filter); isFilter {
		return f.toBSON()
	}

	if b, isBson := idOrFilter.(bson.M); isBson {
		return b
	}

	return bson.M{"id": idOrFilter}
}
