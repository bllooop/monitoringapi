package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bllooop/monitoringapi/backend/internal/domain"
	"github.com/gin-gonic/gin"
)

func JSONStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func (h *Handler) createData(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		newErrorResponse(c, http.StatusBadRequest, "Разрешен только метод POST")
		return
	}
	var input domain.PingResult
	err := c.BindJSON(&input)
	if err != nil || input.IP == "" {
		newErrorResponse(c, http.StatusBadRequest, "Некорретный ввод данных")
		return
	}
	input.Timestamp = time.Now()
	id, err := h.usecases.MonitoringService.CreateData(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	res, err := JSONStruct(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) getData(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		newErrorResponse(c, http.StatusBadRequest, "Требуется запрос GET")
		return
	}
	lists, err := h.usecases.GetData("1")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, lists)
}
