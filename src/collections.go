// collections.go
package main

var collections = []map[string]interface{}{

	{
		"name": "Account",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AccountHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ActivityDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ActivityDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AdverseEvent",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AdverseEventHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AllergyIntolerance",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AllergyIntoleranceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Condition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ConditionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DiagnosticReport",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DiagnosticReportHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "FamilyMemberHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "FamilyMemberHistoryHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "List",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ListHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Medication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationAdministration",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationAdministrationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationDispense",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationDispenseHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationStatement",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationStatementHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Observation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ObservationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Procedure",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ProcedureHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ServiceRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ServiceRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CarePlan",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CarePlanHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CareTeam",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CareTeamHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ClinicalImpression",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ClinicalImpressionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Composition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CompositionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Consent",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ConsentHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Encounter",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EncounterHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EpisodeOfCare",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EpisodeOfCareHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Flag",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "FlagHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Immunization",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImmunizationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RiskAssessment",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RiskAssessmentHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SupplyRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SupplyRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DetectedIssue",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DetectedIssueHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DocumentManifest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DocumentManifestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DocumentReference",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DocumentReferenceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Goal",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "GoalHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImagingStudy",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImagingStudyHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "NutritionOrder",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "NutritionOrderHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SupplyDelivery",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SupplyDeliveryHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "VisionPrescription",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "VisionPrescriptionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceUseStatement",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceUseStatementHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Appointment",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AppointmentHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AppointmentResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AppointmentResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AuditEvent",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "AuditEventHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Basic",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "BasicHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "BodyStructure",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "BodyStructureHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Bundle",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "BundleHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CapabilityStatement",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
			{"name": "resourceSearch", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": false},
		},
	},

	{
		"name": "CapabilityStatementHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
			{"name": "resourceSearch", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": false},
		},
	},

	{
		"name": "CodeSystem",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CodeSystemHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CompartmentDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
			{"name": "resourceSearch", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": false},
		},
	},

	{
		"name": "CompartmentDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
			{"name": "resourceSearch", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": false},
		},
	},

	{
		"name": "ConceptMap",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ConceptMapHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "GraphDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "GraphDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImplementationGuide",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
			{"name": "resourceSearch", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": false},
		},
	},

	{
		"name": "ImplementationGuideHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
			{"name": "resourceSearch", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": false},
		},
	},

	{
		"name": "MessageDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MessageDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "NamingSystem",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "NamingSystemHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "OperationDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "OperationDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SearchParameter",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SearchParameterHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "StructureDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "StructureDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "StructureMap",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "StructureMapHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "TerminologyCapabilities",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "TerminologyCapabilitiesHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ValueSet",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ValueSetHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ChargeItem",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ChargeItemHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ChargeItemDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ChargeItemDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Claim",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ClaimHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ClaimResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ClaimResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Communication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CommunicationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CommunicationRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CommunicationRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Contract",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ContractHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Coverage",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CoverageHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CoverageEligibilityRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CoverageEligibilityRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CoverageEligibilityResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "CoverageEligibilityResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Device",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceMetric",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "DeviceMetricHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EffectEvidenceSynthesis",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EffectEvidenceSynthesisHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Endpoint",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EndpointHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EnrollmentRequest",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EnrollmentRequestHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EnrollmentResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EnrollmentResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EventDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EventDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Evidence",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EvidenceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EvidenceVariable",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "EvidenceVariableHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ExampleScenario",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ExampleScenarioHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ExplanationOfBenefit",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ExplanationOfBenefitHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Group",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "GroupHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "GuidanceResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "GuidanceResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "HealthcareService",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "HealthcareServiceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImmunizationEvaluation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImmunizationEvaluationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImmunizationRecommendation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ImmunizationRecommendationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "InsurancePlan",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "InsurancePlanHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Invoice",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "InvoiceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Library",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "LibraryHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Linkage",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "LinkageHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Location",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "LocationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Measure",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MeasureHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MeasureReport",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MeasureReportHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Media",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MediaHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationKnowledge",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicationKnowledgeHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProduct",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductAuthorization",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductAuthorizationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductContraindication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductContraindicationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductIndication",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductIndicationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductInteraction",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductInteractionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductPackaged",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductPackagedHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductPharmaceutical",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductPharmaceuticalHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductUndesirableEffect",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MedicinalProductUndesirableEffectHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MessageHeader",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MessageHeaderHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MolecularSequence",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "MolecularSequenceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Organization",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "OrganizationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "OrganizationAffiliation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "OrganizationAffiliationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Patient",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PatientHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Person",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PersonHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Practitioner",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PractitionerHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RelatedPerson",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RelatedPersonHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PractitionerRole",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PractitionerRoleHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PaymentNotice",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PaymentNoticeHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PaymentReconciliation",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PaymentReconciliationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PlanDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "PlanDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Provenance",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ProvenanceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Questionnaire",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "QuestionnaireHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "QuestionnaireResponse",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "QuestionnaireResponseHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RequestGroup",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RequestGroupHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchElementDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchElementDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchStudy",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchStudyHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchSubject",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ResearchSubjectHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RiskEvidenceSynthesis",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "RiskEvidenceSynthesisHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Schedule",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "ScheduleHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Slot",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SlotHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Specimen",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SpecimenHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SpecimenDefinition",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SpecimenDefinitionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Subscription",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SubscriptionHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Substance",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SubstanceHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SubstanceSpecification",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "SubstanceSpecificationHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "Task",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "TaskHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "TestReport",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "TestReportHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "TestScript",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "TestScriptHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "VerificationResult",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},

	{
		"name": "VerificationResultHistory",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number", "required": true},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
		},
	},
}
