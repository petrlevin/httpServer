package srv

import (
	"encoding/json"
	"net/http"
)

type Handler struct
{

}

type JsonResponse struct{
	Name string  `json:"message"`
}

func(h *Handler) ServeHTTP(rw http.ResponseWriter,r *http.Request){

	var resp,err =json.Marshal( & JsonResponse{
		 Name: "hhh",
	} )

	if err!=nil{

	}

	rw.Write(resp);



}