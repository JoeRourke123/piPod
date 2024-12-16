package endpoints

import (
	"conductor/db"
	"conductor/util"
	"conductor/util/api"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"strings"
)

var (
	DatabaseEndpoints = []Endpoints{
		{
			Method:  "GET",
			Path:    api.Collections(),
			Handler: handleGetCollections,
		},
		{
			Method:  "GET",
			Path:    api.Collection(":name"),
			Handler: handleGetCollection,
		},
		{
			Method:  "DELETE",
			Path:    api.Collection(":name"),
			Handler: handleDeleteCollection,
		},
		{
			Method:  "PUT",
			Path:    api.Collection(":name"),
			Handler: handleUpdateCollection,
		},
	}
)

func handleGetCollections(ctx *fiber.Ctx) error {
	collections, _ := db.X.ListCollections()
	return ctx.JSON(util.Map(collections, func(c string) string {
		return strings.ReplaceAll(c, "./", "")
	}))
}

func handleGetCollection(ctx *fiber.Ctx) error {
	collectionName := ctx.Params("name")
	offset := ctx.QueryInt("offset", 0)
	id := util.UriToId(ctx.Query("id", ""))

	collection, _ := db.X.FindAll(query.NewQuery("./" + collectionName).Skip(offset).Limit(5))

	if id != "" {
		docs, err := db.X.FindAll(query.NewQuery("./" + collectionName).Where(query.Field("id").Eq(id)))
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		collection = docs
	}

	return ctx.JSON(util.Map(collection, func(doc *document.Document) map[string]interface{} {
		return doc.AsMap()
	}))
}

func handleDeleteCollection(ctx *fiber.Ctx) error {
	collectionName := ctx.Params("name")
	err := db.X.Delete(query.NewQuery("./" + collectionName))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(200)
}

func handleUpdateCollection(ctx *fiber.Ctx) error {
	collectionName := ctx.Params("name")
	jsonBody := make([]map[string]interface{}, 0)
	err := ctx.BodyParser(&jsonBody)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	for _, jsonDoc := range jsonBody {
		err := db.X.UpdateFunc(query.NewQuery("./"+collectionName).Where(query.Field("_id").Eq(jsonDoc["_id"].(string))), func(doc *document.Document) *document.Document {
			return util.NewDocumentOf(jsonDoc)
		})
		if err != nil {
			if errors.Is(err, clover.ErrDocumentNotExist) {
				_, err := db.X.InsertOne("./"+collectionName, document.NewDocumentOf(jsonDoc))
				if err == nil {
					continue
				}
			}

			return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.SendStatus(200)
}
