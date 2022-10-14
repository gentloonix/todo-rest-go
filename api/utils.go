package api

import (
	"encoding/json"
	"main/models"

	"github.com/gin-gonic/gin"
)

// parseApiQuery RESTful API query parser
func parseApiQuery(c *gin.Context) (where map[string]interface{}, order string, updates map[string]interface{}) {
	where_tmp := c.Request.URL.Query()

	if order_tmp, ok := where_tmp[SORT_BY]; ok {
		delete(where_tmp, SORT_BY)
		if len(order_tmp) != 0 {
			order = order_tmp[0]
		}
	}

	if updates_tmp_slice, ok := where_tmp[UPDATE_WITH]; ok {
		delete(where_tmp, UPDATE_WITH)
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

// ApiGet CRUD Read middleware
func ApiGet[T models.IModels](c *gin.Context) {
	where, order, _ := parseApiQuery(c)
	if objs, err := models.OrmQuery[T](where, order); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg":    "ok",
			"result": objs,
		})
	}
}

// ApiPost CRUD Create middleware
func ApiPost[T models.IModels](c *gin.Context) {
	var obj T
	var objs []T
	if err := c.BindJSON(&obj); err != nil {
		if err := c.BindJSON(&objs); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		objs = append(objs, obj)
	}
	if err := models.OrmCreate(objs); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg":    "ok",
			"result": objs,
		})
	}
}

// ApiPut CRUD Update middleware
func ApiPut[T models.IModels](c *gin.Context) {
	where, _, values := parseApiQuery(c)
	if err := models.OrmUpdate[T](where, values); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	}
}

// ApiDelete CRUD Delete middleware
func ApiDelete[T models.IModels](c *gin.Context) {
	where, _, _ := parseApiQuery(c)
	if err := models.OrmDelete[T](where); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	}
}
