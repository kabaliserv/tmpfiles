package errcode

import (
	"encoding/json"
	"net/http"
)

// Err struct
type Err struct {
	error
	Map map[string]interface{}
}

// ToJSON :
func (state *Err) ToJSON() []byte {
	jsonToByte, _ := json.Marshal(state.Map)
	return jsonToByte
}

// JSONMap : Get map with error status code and message
func JSONMap(code int) *Err {
	data := make(map[string]interface{})

	switch code {
	case http.StatusNotFound:
		data["status"] = code
		data["message"] = "The resource you requested could not be found"
		break
	case http.StatusBadRequest:
		data["status"] = code
		data["message"] = "Wrong request syntax"
		break
	case http.StatusInternalServerError:
		data["status"] = code
		data["message"] = "Internal Error"
		break
	case http.StatusUnauthorized:
		data["status"] = code
		data["message"] = "The resource you requested requires authorization"
		break
	}
	return &Err{
		Map: data,
	}
}