package dxb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// The names(Id/Xml/Json/Html) in codes may always be camel cased,
// which differ from the ones(ID/XML/JSON/HTML) in documents.
type ObjectId = primitive.ObjectID

type D = primitive.D
type M = primitive.M
type E = primitive.E

type Regex = primitive.Regex
type Raw = bson.Raw
type FindOptions = options.FindOptions

var NewObjectId = primitive.NewObjectID
var NewObjectIdFromHex = primitive.ObjectIDFromHex
