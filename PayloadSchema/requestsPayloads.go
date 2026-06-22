package payloadschema

type StartTripReq struct {
	Rut   string `json:"rut"`
	Plate string `json:"plate"`
}
