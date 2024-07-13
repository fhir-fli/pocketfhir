// collections.go
package main

var collections = []map[string]interface{}{

	{
		"name": "Account",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "owner", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "period", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "AccountHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "AccountToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ActivityDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "ActivityDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ActivityDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "AdverseEvent",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "recorder", "type": "text"},
			{"name": "resultingcondition", "type": "text"},
			{"name": "study", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "substance", "type": "text"},
	},
	  },

	{
		"name": "AdverseEventHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "AdverseEventToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "AllergyIntolerance",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "asserter", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "last-date", "type": "text"},
			{"name": "onset", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "recorder", "type": "text"},
	},
	  },

	{
		"name": "AllergyIntoleranceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "AllergyIntoleranceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Condition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "abatement-date", "type": "text"},
			{"name": "abatement-string", "type": "text"},
			{"name": "asserter", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "evidence-detail", "type": "text"},
			{"name": "onset-date", "type": "text"},
			{"name": "onset-info", "type": "text"},
			{"name": "recorded-date", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "ConditionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ConditionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DeviceRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "authored-on", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "device", "type": "text"},
			{"name": "event-date", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
			{"name": "insurance", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "prior-request", "type": "text"},
			{"name": "requester", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "DeviceRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DeviceRequestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DiagnosticReport",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "issued", "type": "text"},
			{"name": "media", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "result", "type": "text"},
			{"name": "results-interpreter", "type": "text"},
			{"name": "specimen", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "DiagnosticReportHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DiagnosticReportToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "FamilyMemberHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
	},
	  },

	{
		"name": "FamilyMemberHistoryHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "FamilyMemberHistoryToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "List",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "item", "type": "text"},
			{"name": "notes", "type": "text"},
			{"name": "source", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "title", "type": "text"},
	},
	  },

	{
		"name": "ListHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ListToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Medication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "expiration-date", "type": "text"},
			{"name": "ingredient", "type": "text"},
			{"name": "manufacturer", "type": "text"},
	},
	  },

	{
		"name": "MedicationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicationAdministration",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "context", "type": "text"},
			{"name": "device", "type": "text"},
			{"name": "effective-time", "type": "text"},
			{"name": "medication", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "request", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicationAdministrationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicationAdministrationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicationDispense",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "medication", "type": "text"},
			{"name": "context", "type": "text"},
			{"name": "destination", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "prescription", "type": "text"},
			{"name": "receiver", "type": "text"},
			{"name": "responsibleparty", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "whenhandedover", "type": "text"},
			{"name": "whenprepared", "type": "text"},
	},
	  },

	{
		"name": "MedicationDispenseHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicationDispenseToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicationRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "medication", "type": "text"},
			{"name": "authoredon", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "intended-dispenser", "type": "text"},
			{"name": "intended-performer", "type": "text"},
			{"name": "requester", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicationRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicationRequestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicationStatement",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "medication", "type": "text"},
			{"name": "context", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "source", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicationStatementHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicationStatementToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Observation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "device", "type": "text"},
			{"name": "focus", "type": "text"},
			{"name": "has-member", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "specimen", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "value-date", "type": "text"},
			{"name": "value-string", "type": "text"},
	},
	  },

	{
		"name": "ObservationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ObservationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Procedure",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "reason-reference", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "ProcedureHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ProcedureToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ServiceRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "authored", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
			{"name": "occurrence", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "replaces", "type": "text"},
			{"name": "requester", "type": "text"},
			{"name": "specimen", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "ServiceRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ServiceRequestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CarePlan",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "activity-date", "type": "text"},
			{"name": "activity-reference", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "care-team", "type": "text"},
			{"name": "condition", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "goal", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "replaces", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "CarePlanHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CarePlanToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CareTeam",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "participant", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "CareTeamHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CareTeamToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ClinicalImpression",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "assessor", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "finding-ref", "type": "text"},
			{"name": "investigation", "type": "text"},
			{"name": "previous", "type": "text"},
			{"name": "problem", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "supporting-info", "type": "text"},
	},
	  },

	{
		"name": "ClinicalImpressionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ClinicalImpressionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Composition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "attester", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "entry", "type": "text"},
			{"name": "period", "type": "text"},
			{"name": "related-ref", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "title", "type": "text"},
	},
	  },

	{
		"name": "CompositionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CompositionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Consent",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "actor", "type": "text"},
			{"name": "consentor", "type": "text"},
			{"name": "data", "type": "text"},
			{"name": "organization", "type": "text"},
			{"name": "period", "type": "text"},
			{"name": "source-reference", "type": "text"},
	},
	  },

	{
		"name": "ConsentHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ConsentToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Encounter",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "account", "type": "text"},
			{"name": "appointment", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "diagnosis", "type": "text"},
			{"name": "episode-of-care", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "location-period", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "participant", "type": "text"},
			{"name": "practitioner", "type": "text"},
			{"name": "reason-reference", "type": "text"},
			{"name": "service-provider", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "EncounterHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EncounterToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "EpisodeOfCare",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "care-manager", "type": "text"},
			{"name": "condition", "type": "text"},
			{"name": "incoming-referral", "type": "text"},
			{"name": "organization", "type": "text"},
	},
	  },

	{
		"name": "EpisodeOfCareHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EpisodeOfCareToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Flag",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "FlagHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "FlagToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Immunization",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "lot-number", "type": "text"},
			{"name": "manufacturer", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "reaction", "type": "text"},
			{"name": "reaction-date", "type": "text"},
			{"name": "reason-reference", "type": "text"},
			{"name": "series", "type": "text"},
	},
	  },

	{
		"name": "ImmunizationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ImmunizationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "RiskAssessment",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "condition", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "probability", "type": "number"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "RiskAssessmentHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "RiskAssessmentToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "SupplyRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "requester", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "supplier", "type": "text"},
	},
	  },

	{
		"name": "SupplyRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SupplyRequestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DetectedIssue",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "identified", "type": "text"},
			{"name": "implicated", "type": "text"},
	},
	  },

	{
		"name": "DetectedIssueHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DetectedIssueToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DocumentManifest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "item", "type": "text"},
			{"name": "recipient", "type": "text"},
			{"name": "related-ref", "type": "text"},
			{"name": "source", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "DocumentManifestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DocumentManifestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DocumentReference",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "authenticator", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "custodian", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "period", "type": "text"},
			{"name": "related", "type": "text"},
			{"name": "relatesto", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "DocumentReferenceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DocumentReferenceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Goal",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "start-date", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "target-date", "type": "text"},
	},
	  },

	{
		"name": "GoalHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "GoalToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ImagingStudy",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "basedon", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "endpoint", "type": "text"},
			{"name": "interpreter", "type": "text"},
			{"name": "performer", "type": "text"},
			{"name": "referrer", "type": "text"},
			{"name": "started", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "ImagingStudyHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ImagingStudyToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "NutritionOrder",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "datetime", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
			{"name": "provider", "type": "text"},
	},
	  },

	{
		"name": "NutritionOrderHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "NutritionOrderToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "SupplyDelivery",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "receiver", "type": "text"},
			{"name": "supplier", "type": "text"},
	},
	  },

	{
		"name": "SupplyDeliveryHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SupplyDeliveryToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "VisionPrescription",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "datewritten", "type": "text"},
			{"name": "prescriber", "type": "text"},
	},
	  },

	{
		"name": "VisionPrescriptionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "VisionPrescriptionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DeviceUseStatement",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "device", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "DeviceUseStatementHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DeviceUseStatementToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Appointment",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "actor", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "practitioner", "type": "text"},
			{"name": "reason-reference", "type": "text"},
			{"name": "slot", "type": "text"},
			{"name": "supporting-info", "type": "text"},
	},
	  },

	{
		"name": "AppointmentHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "AppointmentToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "AppointmentResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "actor", "type": "text"},
			{"name": "appointment", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "practitioner", "type": "text"},
	},
	  },

	{
		"name": "AppointmentResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "AppointmentResponseToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "AuditEvent",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "agent", "type": "text"},
			{"name": "agent-name", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "entity", "type": "text"},
			{"name": "entity-name", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "policy", "type": "text"},
			{"name": "source", "type": "text"},
	},
	  },

	{
		"name": "AuditEventHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "AuditEventToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Basic",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "BasicHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "BasicToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "BodyStructure",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
	},
	  },

	{
		"name": "BodyStructureHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "BodyStructureToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Bundle",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composition", "type": "text"},
			{"name": "message", "type": "text"},
			{"name": "timestamp", "type": "text"},
	},
	  },

	{
		"name": "BundleHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "BundleToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CapabilityStatement",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "guide", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "resource-profile", "type": "text"},
			{"name": "software", "type": "text"},
			{"name": "supported-profile", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "CapabilityStatementHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CapabilityStatementToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CodeSystem",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "supplements", "type": "text"},
			{"name": "system", "type": "text"},
	},
	  },

	{
		"name": "CodeSystemHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CodeSystemToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CompartmentDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "CompartmentDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CompartmentDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ConceptMap",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "dependson", "type": "text"},
			{"name": "other", "type": "text"},
			{"name": "product", "type": "text"},
			{"name": "source", "type": "text"},
			{"name": "source-system", "type": "text"},
			{"name": "source-uri", "type": "text"},
			{"name": "target", "type": "text"},
			{"name": "target-system", "type": "text"},
			{"name": "target-uri", "type": "text"},
	},
	  },

	{
		"name": "ConceptMapHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ConceptMapToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "GraphDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "GraphDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "GraphDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ImplementationGuide",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "global", "type": "text"},
			{"name": "resource", "type": "text"},
	},
	  },

	{
		"name": "ImplementationGuideHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ImplementationGuideToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MessageDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "parent", "type": "text"},
	},
	  },

	{
		"name": "MessageDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MessageDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "NamingSystem",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "contact", "type": "text"},
			{"name": "period", "type": "text"},
			{"name": "responsible", "type": "text"},
			{"name": "value", "type": "text"},
	},
	  },

	{
		"name": "NamingSystemHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "NamingSystemToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "OperationDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "base", "type": "text"},
			{"name": "input-profile", "type": "text"},
			{"name": "output-profile", "type": "text"},
	},
	  },

	{
		"name": "OperationDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "OperationDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "SearchParameter",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "component", "type": "text"},
			{"name": "derived-from", "type": "text"},
	},
	  },

	{
		"name": "SearchParameterHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SearchParameterToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "StructureDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "base", "type": "text"},
			{"name": "type", "type": "text"},
			{"name": "valueset", "type": "text"},
	},
	  },

	{
		"name": "StructureDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "StructureDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "StructureMap",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "StructureMapHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "StructureMapToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "TerminologyCapabilities",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "TerminologyCapabilitiesHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "TerminologyCapabilitiesToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ValueSet",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
			{"name": "expansion", "type": "text"},
			{"name": "reference", "type": "text"},
	},
	  },

	{
		"name": "ValueSetHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ValueSetToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ChargeItem",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "account", "type": "text"},
			{"name": "context", "type": "text"},
			{"name": "entered-date", "type": "text"},
			{"name": "enterer", "type": "text"},
			{"name": "factor-override", "type": "number"},
			{"name": "occurrence", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "performer-actor", "type": "text"},
			{"name": "performing-organization", "type": "text"},
			{"name": "requesting-organization", "type": "text"},
			{"name": "service", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "ChargeItemHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ChargeItemToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ChargeItemDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "ChargeItemDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ChargeItemDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Claim",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "care-team", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "detail-udi", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "enterer", "type": "text"},
			{"name": "facility", "type": "text"},
			{"name": "insurer", "type": "text"},
			{"name": "item-udi", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "payee", "type": "text"},
			{"name": "procedure-udi", "type": "text"},
			{"name": "provider", "type": "text"},
			{"name": "subdetail-udi", "type": "text"},
	},
	  },

	{
		"name": "ClaimHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ClaimToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ClaimResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "disposition", "type": "text"},
			{"name": "insurer", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "payment-date", "type": "text"},
			{"name": "request", "type": "text"},
			{"name": "requestor", "type": "text"},
	},
	  },

	{
		"name": "ClaimResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ClaimResponseToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Communication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "received", "type": "text"},
			{"name": "recipient", "type": "text"},
			{"name": "sender", "type": "text"},
			{"name": "sent", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "CommunicationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CommunicationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CommunicationRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "authored", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "occurrence", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "recipient", "type": "text"},
			{"name": "replaces", "type": "text"},
			{"name": "requester", "type": "text"},
			{"name": "sender", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "CommunicationRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CommunicationRequestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Contract",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "authority", "type": "text"},
			{"name": "domain", "type": "text"},
			{"name": "instantiates", "type": "text"},
			{"name": "issued", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "signer", "type": "text"},
			{"name": "subject", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "ContractHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ContractToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Coverage",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "beneficiary", "type": "text"},
			{"name": "class-value", "type": "text"},
			{"name": "dependent", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "payor", "type": "text"},
			{"name": "policy-holder", "type": "text"},
			{"name": "subscriber", "type": "text"},
	},
	  },

	{
		"name": "CoverageHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CoverageToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CoverageEligibilityRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "enterer", "type": "text"},
			{"name": "facility", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "provider", "type": "text"},
	},
	  },

	{
		"name": "CoverageEligibilityRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CoverageEligibilityRequestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "CoverageEligibilityResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "disposition", "type": "text"},
			{"name": "insurer", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "request", "type": "text"},
			{"name": "requestor", "type": "text"},
	},
	  },

	{
		"name": "CoverageEligibilityResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "CoverageEligibilityResponseToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Device",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "device-name", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "manufacturer", "type": "text"},
			{"name": "model", "type": "text"},
			{"name": "organization", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "udi-carrier", "type": "text"},
			{"name": "udi-di", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "DeviceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DeviceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DeviceDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "parent", "type": "text"},
	},
	  },

	{
		"name": "DeviceDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DeviceDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "DeviceMetric",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "parent", "type": "text"},
			{"name": "source", "type": "text"},
	},
	  },

	{
		"name": "DeviceMetricHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "DeviceMetricToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "EffectEvidenceSynthesis",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "EffectEvidenceSynthesisHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EffectEvidenceSynthesisToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Endpoint",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "organization", "type": "text"},
	},
	  },

	{
		"name": "EndpointHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EndpointToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "EnrollmentRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "EnrollmentRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EnrollmentRequestToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "EnrollmentResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "request", "type": "text"},
	},
	  },

	{
		"name": "EnrollmentResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EnrollmentResponseToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "EventDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "EventDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EventDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Evidence",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "EvidenceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EvidenceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "EvidenceVariable",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "EvidenceVariableHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "EvidenceVariableToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ExampleScenario",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "ExampleScenarioHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ExampleScenarioToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ExplanationOfBenefit",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "care-team", "type": "text"},
			{"name": "claim", "type": "text"},
			{"name": "coverage", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "detail-udi", "type": "text"},
			{"name": "disposition", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "enterer", "type": "text"},
			{"name": "facility", "type": "text"},
			{"name": "item-udi", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "payee", "type": "text"},
			{"name": "procedure-udi", "type": "text"},
			{"name": "provider", "type": "text"},
			{"name": "subdetail-udi", "type": "text"},
	},
	  },

	{
		"name": "ExplanationOfBenefitHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ExplanationOfBenefitToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Group",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "managing-entity", "type": "text"},
			{"name": "member", "type": "text"},
	},
	  },

	{
		"name": "GroupHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "GroupToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "GuidanceResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "GuidanceResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "GuidanceResponseToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "HealthcareService",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "coverage-area", "type": "text"},
			{"name": "endpoint", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "organization", "type": "text"},
	},
	  },

	{
		"name": "HealthcareServiceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "HealthcareServiceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ImmunizationEvaluation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "immunization-event", "type": "text"},
			{"name": "patient", "type": "text"},
	},
	  },

	{
		"name": "ImmunizationEvaluationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ImmunizationEvaluationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ImmunizationRecommendation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "information", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "support", "type": "text"},
	},
	  },

	{
		"name": "ImmunizationRecommendationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ImmunizationRecommendationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "InsurancePlan",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "address-city", "type": "text"},
			{"name": "address-country", "type": "text"},
			{"name": "address-postalcode", "type": "text"},
			{"name": "address-state", "type": "text"},
			{"name": "administered-by", "type": "text"},
			{"name": "endpoint", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "owned-by", "type": "text"},
			{"name": "phonetic", "type": "text"},
	},
	  },

	{
		"name": "InsurancePlanHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "InsurancePlanToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Invoice",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "account", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "issuer", "type": "text"},
			{"name": "participant", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "recipient", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "InvoiceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "InvoiceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Library",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "LibraryHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "LibraryToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Linkage",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "item", "type": "text"},
			{"name": "source", "type": "text"},
	},
	  },

	{
		"name": "LinkageHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "LinkageToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Location",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "address-city", "type": "text"},
			{"name": "address-country", "type": "text"},
			{"name": "address-postalcode", "type": "text"},
			{"name": "address-state", "type": "text"},
			{"name": "endpoint", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "organization", "type": "text"},
			{"name": "partof", "type": "text"},
	},
	  },

	{
		"name": "LocationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "LocationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Measure",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "MeasureHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MeasureToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MeasureReport",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "evaluated-resource", "type": "text"},
			{"name": "measure", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "period", "type": "text"},
			{"name": "reporter", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MeasureReportHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MeasureReportToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Media",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "device", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "operator", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MediaHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MediaToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicationKnowledge",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "ingredient", "type": "text"},
			{"name": "manufacturer", "type": "text"},
			{"name": "monograph", "type": "text"},
	},
	  },

	{
		"name": "MedicationKnowledgeHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicationKnowledgeToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProduct",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "name", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProductAuthorization",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "holder", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductAuthorizationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductAuthorizationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProductContraindication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductContraindicationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductContraindicationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProductIndication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductIndicationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductIndicationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProductInteraction",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductInteractionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductInteractionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProductPackaged",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductPackagedHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductPackagedToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProductPharmaceutical",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductPharmaceuticalHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductPharmaceuticalToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MedicinalProductUndesirableEffect",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "MedicinalProductUndesirableEffectHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MedicinalProductUndesirableEffectToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MessageHeader",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "destination", "type": "text"},
			{"name": "destination-uri", "type": "text"},
			{"name": "enterer", "type": "text"},
			{"name": "focus", "type": "text"},
			{"name": "receiver", "type": "text"},
			{"name": "responsible", "type": "text"},
			{"name": "sender", "type": "text"},
			{"name": "source", "type": "text"},
			{"name": "source-uri", "type": "text"},
			{"name": "target", "type": "text"},
	},
	  },

	{
		"name": "MessageHeaderHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MessageHeaderToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "MolecularSequence",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "variant-end", "type": "number"},
			{"name": "variant-start", "type": "number"},
			{"name": "window-end", "type": "number"},
			{"name": "window-start", "type": "number"},
	},
	  },

	{
		"name": "MolecularSequenceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "MolecularSequenceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Organization",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "address-city", "type": "text"},
			{"name": "address-country", "type": "text"},
			{"name": "address-postalcode", "type": "text"},
			{"name": "address-state", "type": "text"},
			{"name": "endpoint", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "partof", "type": "text"},
			{"name": "phonetic", "type": "text"},
	},
	  },

	{
		"name": "OrganizationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "OrganizationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "OrganizationAffiliation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "endpoint", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "network", "type": "text"},
			{"name": "participating-organization", "type": "text"},
			{"name": "primary-organization", "type": "text"},
			{"name": "service", "type": "text"},
	},
	  },

	{
		"name": "OrganizationAffiliationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "OrganizationAffiliationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Patient",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "address-city", "type": "text"},
			{"name": "address-country", "type": "text"},
			{"name": "address-postalcode", "type": "text"},
			{"name": "address-state", "type": "text"},
			{"name": "birthdate", "type": "text"},
			{"name": "death-date", "type": "text"},
			{"name": "family", "type": "text"},
			{"name": "general-practitioner", "type": "text"},
			{"name": "given", "type": "text"},
			{"name": "link", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "organization", "type": "text"},
			{"name": "phonetic", "type": "text"},
	},
	  },

	{
		"name": "PatientHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "PatientToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Person",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "address-city", "type": "text"},
			{"name": "address-country", "type": "text"},
			{"name": "address-postalcode", "type": "text"},
			{"name": "address-state", "type": "text"},
			{"name": "birthdate", "type": "text"},
			{"name": "phonetic", "type": "text"},
			{"name": "link", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "organization", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "practitioner", "type": "text"},
			{"name": "relatedperson", "type": "text"},
	},
	  },

	{
		"name": "PersonHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "PersonToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Practitioner",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "address-city", "type": "text"},
			{"name": "address-country", "type": "text"},
			{"name": "address-postalcode", "type": "text"},
			{"name": "address-state", "type": "text"},
			{"name": "family", "type": "text"},
			{"name": "given", "type": "text"},
			{"name": "phonetic", "type": "text"},
			{"name": "name", "type": "text"},
	},
	  },

	{
		"name": "PractitionerHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "PractitionerToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "RelatedPerson",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "address", "type": "text"},
			{"name": "address-city", "type": "text"},
			{"name": "address-country", "type": "text"},
			{"name": "address-postalcode", "type": "text"},
			{"name": "address-state", "type": "text"},
			{"name": "birthdate", "type": "text"},
			{"name": "phonetic", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "patient", "type": "text"},
	},
	  },

	{
		"name": "RelatedPersonHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "RelatedPersonToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "PractitionerRole",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "endpoint", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "organization", "type": "text"},
			{"name": "practitioner", "type": "text"},
			{"name": "service", "type": "text"},
	},
	  },

	{
		"name": "PractitionerRoleHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "PractitionerRoleToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "PaymentNotice",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "provider", "type": "text"},
			{"name": "request", "type": "text"},
			{"name": "response", "type": "text"},
	},
	  },

	{
		"name": "PaymentNoticeHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "PaymentNoticeToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "PaymentReconciliation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "created", "type": "text"},
			{"name": "disposition", "type": "text"},
			{"name": "payment-issuer", "type": "text"},
			{"name": "request", "type": "text"},
			{"name": "requestor", "type": "text"},
	},
	  },

	{
		"name": "PaymentReconciliationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "PaymentReconciliationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "PlanDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "definition", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "PlanDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "PlanDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Provenance",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "agent", "type": "text"},
			{"name": "entity", "type": "text"},
			{"name": "location", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "recorded", "type": "text"},
			{"name": "target", "type": "text"},
			{"name": "when", "type": "text"},
	},
	  },

	{
		"name": "ProvenanceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ProvenanceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Questionnaire",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "definition", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "QuestionnaireHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "QuestionnaireToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "QuestionnaireResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "authored", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "questionnaire", "type": "text"},
			{"name": "source", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "QuestionnaireResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "QuestionnaireResponseToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "RequestGroup",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "author", "type": "text"},
			{"name": "authored", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "instantiates-canonical", "type": "text"},
			{"name": "instantiates-uri", "type": "text"},
			{"name": "participant", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "RequestGroupHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "RequestGroupToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ResearchDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "ResearchDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ResearchDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ResearchElementDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "composed-of", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "depends-on", "type": "text"},
			{"name": "derived-from", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "predecessor", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "successor", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "ResearchElementDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ResearchElementDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ResearchStudy",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "partof", "type": "text"},
			{"name": "principalinvestigator", "type": "text"},
			{"name": "protocol", "type": "text"},
			{"name": "site", "type": "text"},
			{"name": "sponsor", "type": "text"},
			{"name": "title", "type": "text"},
	},
	  },

	{
		"name": "ResearchStudyHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ResearchStudyToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "ResearchSubject",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "individual", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "study", "type": "text"},
	},
	  },

	{
		"name": "ResearchSubjectHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ResearchSubjectToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "RiskEvidenceSynthesis",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "effective", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "RiskEvidenceSynthesisHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "RiskEvidenceSynthesisToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Schedule",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "actor", "type": "text"},
			{"name": "date", "type": "text"},
	},
	  },

	{
		"name": "ScheduleHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "ScheduleToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Slot",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "schedule", "type": "text"},
			{"name": "start", "type": "text"},
	},
	  },

	{
		"name": "SlotHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SlotToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Specimen",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "collected", "type": "text"},
			{"name": "collector", "type": "text"},
			{"name": "parent", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "SpecimenHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SpecimenToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "SpecimenDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
	},
	  },

	{
		"name": "SpecimenDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SpecimenDefinitionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Subscription",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "criteria", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "SubscriptionHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SubscriptionToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Substance",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "expiry", "type": "text"},
			{"name": "substance-reference", "type": "text"},
	},
	  },

	{
		"name": "SubstanceHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SubstanceToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "SubstanceSpecification",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
	},
	  },

	{
		"name": "SubstanceSpecificationHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "SubstanceSpecificationToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "Task",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "authored-on", "type": "text"},
			{"name": "based-on", "type": "text"},
			{"name": "encounter", "type": "text"},
			{"name": "focus", "type": "text"},
			{"name": "modified", "type": "text"},
			{"name": "owner", "type": "text"},
			{"name": "part-of", "type": "text"},
			{"name": "patient", "type": "text"},
			{"name": "period", "type": "text"},
			{"name": "requester", "type": "text"},
			{"name": "subject", "type": "text"},
	},
	  },

	{
		"name": "TaskHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "TaskToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "TestReport",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "issued", "type": "text"},
			{"name": "participant", "type": "text"},
			{"name": "tester", "type": "text"},
			{"name": "testscript", "type": "text"},
	},
	  },

	{
		"name": "TestReportHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "TestReportToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "TestScript",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "date", "type": "text"},
			{"name": "description", "type": "text"},
			{"name": "name", "type": "text"},
			{"name": "publisher", "type": "text"},
			{"name": "testscript-capability", "type": "text"},
			{"name": "title", "type": "text"},
			{"name": "url", "type": "text"},
	},
	  },

	{
		"name": "TestScriptHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "TestScriptToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},

	{
		"name": "VerificationResult",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
			{"name": "_content", "type": "text"},
			{"name": "_lastUpdated", "type": "text"},
			{"name": "_profile", "type": "text"},
			{"name": "_source", "type": "text"},
			{"name": "_text", "type": "text"},
			{"name": "target", "type": "text"},
	},
	  },

	{
		"name": "VerificationResultHistory",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},

	{
		"name": "VerificationResultToken",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},}