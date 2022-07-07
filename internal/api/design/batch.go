package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("batch", func() {
	Description("The batch service manages batches of packages.")
	HTTP(func() {
		Path("/batch")
	})
	Method("submit", func() {
		Description("Submit a new batch")
		Payload(func() {
			Attribute("batch_location", String)
			Attribute("manifest_location", String)
			Attribute("transfer_type", String)
			Attribute("workflow", String)
			Required("batch_location","manifest_location")
		})
		Result(BatchResult)
		Error("not_available")
		Error("not_valid")
		HTTP(func() {
			POST("/")
			Response(StatusAccepted)
			Response("not_available", StatusConflict)
			Response("not_valid", StatusBadRequest)
		})
	})
	Method("status", func() {
		Description("Retrieve status of current batch operation.")
		Result(BatchStatusResult)
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
	Method("hints", func() {
		Description("Retrieve form hints")
		Result(BatchHintsResult)
		HTTP(func() {
			GET("/hints")
			Response(StatusOK)
		})
	})
})

var BatchResult = Type("BatchResult", func() {
	Attribute("batch_uuid", String)
	Attribute("batch_name", String)
	Attribute("batch_run_id", String)
	Attribute("batch_status", String)
	Required("batch_uuid","batch_name","batch_run_id","batch_status")
})

var BatchStatusResult = Type("BatchStatusResult", func() {
	Attribute("running", Boolean)
	Attribute("status", String)
	Attribute("workflow_id", String)
	Attribute("run_id", String)
	Required("running")
})

var BatchHintsResult = Type("BatchHintsResult", func() {
	Attribute("completed_dirs", ArrayOf(String), "A list of known values of completedDir used by existing watchers.")
})
