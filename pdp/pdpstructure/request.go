package pdpstructure

import (
	"github.com/integration-system/isp-mdb-lib/entity"
	"github.com/integration-system/isp-mdb-lib/pdp"
	"github.com/integration-system/isp-mdb-lib/structure"
)

type PdpUpsertRequest struct {
	structure.SudirUpdateRecordRequest
	UpdateMethod pdp.HandleMethod
}

type BatchConvertPdpRequest struct {
	*structure.ConvertRequestPayload `valid:"required~Required"`
	*structure.AbstractConvertBatchRequest
	Delta PdpDelta
}

/*
	DeletedObject e.q.
	{
		Id: "62D1FA909140BABC0C10DA90A13B041E30",
		DeleteFlag: "VehicleDelete",
		ObjectGroup: "VEHICLE",
	}
*/
type DeletedObject struct {
	Id          string
	DeleteFlag  string
	ObjectGroup string
}

/*
	DeltaInfo e.q.
	{
		AttributeDeltaData: {
			"FIO": {
				"MIDNAME": "IVANOVICH",
				"SURNAME": "",
			}
		}
		ObjectDeltaData: {
			"VEHICLES": [
				{
					"ITEM_ID": "1234567890_1",
					"VEHICLE_DESC": "TEST VEH 1"
				},
				{
					"ITEM_ID": "1234567890_2",
					"VEHICLE_DESC": ""
				}
			]
		}
	}
*/
type DeltaInfo struct {
	AttributeDeltaData map[string]map[string]interface{}
	ObjectDeltaData    map[string][]map[string]interface{}
}

type PdpDelta struct {
	DeltaInfo      DeltaInfo
	DeletedObjects []DeletedObject
}

type PdpNotification struct {
	entity.Notification

	PdpDelta PdpDelta
}

type PdpNotificationMessage struct {
	Notifications []PdpNotification
}