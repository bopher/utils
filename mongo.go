package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoArray generate primitive.A
func MongoArray(args ...interface{}) primitive.A {
	return args
}

// MongoMap generate primitive.M
//
// Args count must even
func MongoMap(args ...interface{}) primitive.M {
	res := make(primitive.M)
	if len(args)%2 == 0 {
		for i := 0; i < len(args); i++ {
			if i%2 == 0 {
				if k, ok := args[i].(string); ok {
					res[k] = args[i+1]
				}
			}
		}
	}
	return res
}

// MongoMaps generate []primitive.M
//
// Args count must even
func MongoMaps(args ...interface{}) []primitive.M {
	res := make([]primitive.M, 0)
	if len(args)%2 == 0 {
		for i := 0; i < len(args); i++ {
			if i%2 == 0 {
				if k, ok := args[i].(string); ok {
					res = append(res, primitive.M{k: args[i+1]})
				}
			}
		}
	}
	return res
}

// MongoDoc generate primitive.D from args
//
// Args count must even
// Example: MongoDoc("_id", 1, "name", "John")
func MongoDoc(args ...interface{}) primitive.D {
	res := make([]primitive.E, 0)
	if len(args)%2 == 0 {
		for i := 0; i < len(args); i++ {
			if i%2 == 0 {
				if k, ok := args[i].(string); ok {
					res = append(res, primitive.E{Key: k, Value: args[i+1]})
				}
			}
		}
	}
	return res
}

// MongoOr generate $or
func MongoOr(v interface{}) primitive.D {
	return MongoDoc("$or", v)
}

// MongoAnd generate $and
func MongoAnd(v interface{}) primitive.D {
	return MongoDoc("$and", v)
}

// MongoIn generate $in
func MongoIn(k string, v interface{}) primitive.D {
	return MongoDoc(k, MongoMap("$in", v))
}

// MongoSet generate $set
func MongoSet(v interface{}) primitive.D {
	return MongoDoc("$set", v)
}

// MongoNestedSet generate nested $set
func MongoNestedSet(k string, v interface{}) primitive.D {
	return MongoDoc("$set", MongoMap(k, v))
}

// MongoMatch generate $match
func MongoMatch(v interface{}) primitive.D {
	return MongoDoc("$match", v)
}

// MongoLimit generate $limit
func MongoLimit(limit uint) primitive.D {
	return MongoDoc("$limit", limit)
}

// MongoSort generate $sort
func MongoSort(sorts interface{}) primitive.D {
	return MongoDoc("$sort", sorts)
}

// MongoSkip generate $skip
func MongoSkip(skip uint64) primitive.D {
	return MongoDoc("$skip", skip)
}

// MongoLookup generate $lookup
func MongoLookup(from string, local string, foreign string, as string) primitive.D {
	return MongoDoc("$lookup", MongoMap("from", from, "localField", local, "foreignField", foreign, "as", as))
}

// MongoUnwind generate $unwind (with preserveNull)
func MongoUnwind(path string) primitive.D {
	return MongoDoc("$unwind", MongoMap("path", path, "preserveNullAndEmptyArrays", true))
}

// MongoUnwrap get first item of array and insert to doc using $addFields.
func MongoUnwrap(field string, as string) primitive.D {
	return MongoDoc("$addFields", MongoMap(as, MongoMap("$first", field)))
}

// MongoGroup generate $group
func MongoGroup(fields interface{}) primitive.D {
	return MongoDoc("$group", fields)
}

// MongoSetRoot generate $replaceRoot
func MongoSetRoot(newRoot interface{}) primitive.D {
	return MongoDoc("$replaceRoot", MongoMap("newRoot", newRoot))
}

// MongoMergeRoot generate $replaceRoot with $mergeObject
func MongoMergeRoot(fields ...interface{}) primitive.D {
	return MongoSetRoot(MongoMap("$mergeObjects", fields))
}

// MongoUnProject generate $project to remove fields from result
func MongoUnProject(fields ...string) primitive.D {
	_fields := make(map[string]uint)
	for _, v := range fields {
		_fields[v] = 0
	}
	return MongoDoc("$project", _fields)
}

// ParseObjectID parse object id from string
//
// @returns ObjectID
// @fail nil
func ParseObjectID(id string) *primitive.ObjectID {
	if oId, err := primitive.ObjectIDFromHex(id); err == nil && !oId.IsZero() {
		return &oId
	}
	return nil
}

// IsValidOId check if object id is valid and not zero
func IsValidOId(id *primitive.ObjectID) bool {
	return id != nil && !id.IsZero()
}
