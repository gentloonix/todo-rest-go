package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func parseApiQuery(c *gin.Context) (where map[string]interface{}, order string, updates map[string]interface{}) {
	where_tmp := c.Request.URL.Query()

	if order_tmp, ok := where_tmp["sortBy"]; ok {
		delete(where_tmp, "sortBy")
		if len(order_tmp) != 0 {
			order = order_tmp[0]
		}
	}

	if updates_tmp_slice, ok := where_tmp["replaceWith"]; ok {
		delete(where_tmp, "replaceWith")
		if len(updates_tmp_slice) != 0 {
			updates_tmp := updates_tmp_slice[0]
			updates = make(map[string]interface{})
			if err := json.Unmarshal([]byte(updates_tmp), &updates); err != nil {
				updates = nil
			}
		}
	}

	where = make(map[string]interface{})
	for k, v := range where_tmp {
		where[k] = v
	}

	return where, order, updates
}

func ApiGet[T IModels](c *gin.Context) {
	where, order, _ := parseApiQuery(c)
	objs, err := OrmQuery[T](where, order)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"result": objs,
			"msg":    "ok",
		})
	}
}

func ApiPost[T IModels](c *gin.Context) {
	// TODO
}

func ApiPut[T IModels](c *gin.Context) {
	where, _, values := parseApiQuery(c)
	if err := OrmUpdate[T](where, values); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"msg": "ok"})
	}
}

func ApiDelete[T IModels](c *gin.Context) {
	where, _, _ := parseApiQuery(c)
	if err := OrmDelete[T](where); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"msg": "ok"})
	}
}
