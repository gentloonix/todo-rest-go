package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func parseApiQuery(c *gin.Context) (where map[string]interface{}, order string, update map[string]interface{}) {
	where_tmp := c.Request.URL.Query()
	order_tmp, ok := where["sortBy"].([]string)
	delete(where, "sortBy")
	if ok {
		if len(order_tmp) != 0 {
			order = order_tmp[0]
		}
	}
	update_tmp_slice, ok := where["replaceWith"].([]string)
	delete(where, "replaceWith")
	if ok {
		if len(update_tmp_slice) != 0 {
			update_tmp := update_tmp_slice[0]
			update = make(map[string]interface{})
			if err := json.Unmarshal([]byte(update_tmp), &update); err != nil {
				update = nil
			}
		}
	}
	where = make(map[string]interface{})
	for k, v := range where_tmp {
		where[k] = v
	}
	return where, order, update
}

func ApiGet[T IModels](c *gin.Context) {
	where, order := parseApiQuery(c)
	objs, err := OrmQuery[T](where, order)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": objs})
	}
}

func ApiPost[T IModels](c *gin.Context) {
	// TODO
}

func ApiPut[T IModels](c *gin.Context) {
	where, _ := parseApiQuery(c)
	objs, err := OrmUpdate[T](where)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": objs})
	}
}

func ApiDelete[T IModels](c *gin.Context) {
	where, _ := parseApiQuery(c)
	objs, err := OrmDelete[T](where)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": objs})
	}
}
