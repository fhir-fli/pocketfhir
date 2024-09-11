package main

// Common schema applied to multiple collections
var commonSchema = []map[string]interface{}{
	{"name": "versionId", "type": "datetime", "required": true}, // Changed from "number" to "datetime"
	{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}, "required": true},
}

// List of collections
var collections = []map[string]interface{}{

	{
		"name":   "Account",
		"schema": commonSchema,
	},
	{
		"name":         "AccountHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ActivityDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ActivityDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "AdverseEvent",
		"schema": commonSchema,
	},
	{
		"name":         "AdverseEventHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "AllergyIntolerance",
		"schema": commonSchema,
	},
	{
		"name":         "AllergyIntoleranceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Condition",
		"schema": commonSchema,
	},
	{
		"name":         "ConditionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DeviceRequest",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DiagnosticReport",
		"schema": commonSchema,
	},
	{
		"name":         "DiagnosticReportHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "List",
		"schema": commonSchema,
	},
	{
		"name":         "ListHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Medication",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicationAdministration",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationAdministrationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicationDispense",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationDispenseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicationRequest",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicationStatement",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationStatementHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Observation",
		"schema": commonSchema,
	},
	{
		"name":         "ObservationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Procedure",
		"schema": commonSchema,
	},
	{
		"name":         "ProcedureHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ServiceRequest",
		"schema": commonSchema,
	},
	{
		"name":         "ServiceRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "CarePlan",
		"schema": commonSchema,
	},
	{
		"name":         "CarePlanHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "CareTeam",
		"schema": commonSchema,
	},
	{
		"name":         "CareTeamHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ClinicalImpression",
		"schema": commonSchema,
	},
	{
		"name":         "ClinicalImpressionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Consent",
		"schema": commonSchema,
	},
	{
		"name":         "ConsentHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Encounter",
		"schema": commonSchema,
	},
	{
		"name":         "EncounterHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "EpisodeOfCare",
		"schema": commonSchema,
	},
	{
		"name":         "EpisodeOfCareHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Flag",
		"schema": commonSchema,
	},
	{
		"name":         "FlagHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Immunization",
		"schema": commonSchema,
	},
	{
		"name":         "ImmunizationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "RiskAssessment",
		"schema": commonSchema,
	},
	{
		"name":         "RiskAssessmentHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "SupplyRequest",
		"schema": commonSchema,
	},
	{
		"name":         "SupplyRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DetectedIssue",
		"schema": commonSchema,
	},
	{
		"name":         "DetectedIssueHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DocumentManifest",
		"schema": commonSchema,
	},
	{
		"name":         "DocumentManifestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DocumentReference",
		"schema": commonSchema,
	},
	{
		"name":         "DocumentReferenceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Goal",
		"schema": commonSchema,
	},
	{
		"name":         "GoalHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ImagingStudy",
		"schema": commonSchema,
	},
	{
		"name":         "ImagingStudyHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "NutritionOrder",
		"schema": commonSchema,
	},
	{
		"name":         "NutritionOrderHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "SupplyDelivery",
		"schema": commonSchema,
	},
	{
		"name":         "SupplyDeliveryHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "VisionPrescription",
		"schema": commonSchema,
	},
	{
		"name":         "VisionPrescriptionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DeviceUseStatement",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceUseStatementHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Appointment",
		"schema": commonSchema,
	},
	{
		"name":         "AppointmentHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "AppointmentResponse",
		"schema": commonSchema,
	},
	{
		"name":         "AppointmentResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "AuditEvent",
		"schema": commonSchema,
	},
	{
		"name":         "AuditEventHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Basic",
		"schema": commonSchema,
	},
	{
		"name":         "BasicHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "BodyStructure",
		"schema": commonSchema,
	},
	{
		"name":         "BodyStructureHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Bundle",
		"schema": commonSchema,
	},
	{
		"name":         "BundleHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "CapabilityStatement",
		"schema": commonSchema,
	},
	{
		"name":         "CapabilityStatementHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "CodeSystem",
		"schema": commonSchema,
	},
	{
		"name":         "CodeSystemHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ConceptMap",
		"schema": commonSchema,
	},
	{
		"name":         "ConceptMapHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "GraphDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "GraphDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ImplementationGuide",
		"schema": commonSchema,
	},
	{
		"name":         "ImplementationGuideHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MessageDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "MessageDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "NamingSystem",
		"schema": commonSchema,
	},
	{
		"name":         "NamingSystemHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "OperationDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "OperationDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "SearchParameter",
		"schema": commonSchema,
	},
	{
		"name":         "SearchParameterHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "StructureDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "StructureDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "StructureMap",
		"schema": commonSchema,
	},
	{
		"name":         "StructureMapHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "TerminologyCapabilities",
		"schema": commonSchema,
	},
	{
		"name":         "TerminologyCapabilitiesHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ValueSet",
		"schema": commonSchema,
	},
	{
		"name":         "ValueSetHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ChargeItem",
		"schema": commonSchema,
	},
	{
		"name":         "ChargeItemHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ChargeItemDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ChargeItemDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Claim",
		"schema": commonSchema,
	},
	{
		"name":         "ClaimHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ClaimResponse",
		"schema": commonSchema,
	},
	{
		"name":         "ClaimResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Communication",
		"schema": commonSchema,
	},
	{
		"name":         "CommunicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "CommunicationRequest",
		"schema": commonSchema,
	},
	{
		"name":         "CommunicationRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Contract",
		"schema": commonSchema,
	},
	{
		"name":         "ContractHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Coverage",
		"schema": commonSchema,
	},
	{
		"name":         "CoverageHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "CoverageEligibilityRequest",
		"schema": commonSchema,
	},
	{
		"name":         "CoverageEligibilityRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "CoverageEligibilityResponse",
		"schema": commonSchema,
	},
	{
		"name":         "CoverageEligibilityResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Device",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DeviceDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "DeviceMetric",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceMetricHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "EffectEvidenceSynthesis",
		"schema": commonSchema,
	},
	{
		"name":         "EffectEvidenceSynthesisHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Endpoint",
		"schema": commonSchema,
	},
	{
		"name":         "EndpointHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "EnrollmentRequest",
		"schema": commonSchema,
	},
	{
		"name":         "EnrollmentRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "EnrollmentResponse",
		"schema": commonSchema,
	},
	{
		"name":         "EnrollmentResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "EventDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "EventDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Evidence",
		"schema": commonSchema,
	},
	{
		"name":         "EvidenceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "EvidenceVariable",
		"schema": commonSchema,
	},
	{
		"name":         "EvidenceVariableHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ExampleScenario",
		"schema": commonSchema,
	},
	{
		"name":         "ExampleScenarioHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ExplanationOfBenefit",
		"schema": commonSchema,
	},
	{
		"name":         "ExplanationOfBenefitHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Group",
		"schema": commonSchema,
	},
	{
		"name":         "GroupHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "GuidanceResponse",
		"schema": commonSchema,
	},
	{
		"name":         "GuidanceResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "HealthcareService",
		"schema": commonSchema,
	},
	{
		"name":         "HealthcareServiceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ImmunizationEvaluation",
		"schema": commonSchema,
	},
	{
		"name":         "ImmunizationEvaluationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ImmunizationRecommendation",
		"schema": commonSchema,
	},
	{
		"name":         "ImmunizationRecommendationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "InsurancePlan",
		"schema": commonSchema,
	},
	{
		"name":         "InsurancePlanHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Invoice",
		"schema": commonSchema,
	},
	{
		"name":         "InvoiceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Library",
		"schema": commonSchema,
	},
	{
		"name":         "LibraryHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Linkage",
		"schema": commonSchema,
	},
	{
		"name":         "LinkageHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Location",
		"schema": commonSchema,
	},
	{
		"name":         "LocationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Measure",
		"schema": commonSchema,
	},
	{
		"name":         "MeasureHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MeasureReport",
		"schema": commonSchema,
	},
	{
		"name":         "MeasureReportHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Media",
		"schema": commonSchema,
	},
	{
		"name":         "MediaHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicationKnowledge",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationKnowledgeHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProduct",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProductAuthorization",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductAuthorizationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProductContraindication",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductContraindicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProductIndication",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductIndicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProductInteraction",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductInteractionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProductPackaged",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductPackagedHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProductPharmaceutical",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductPharmaceuticalHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MedicinalProductUndesirableEffect",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductUndesirableEffectHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MessageHeader",
		"schema": commonSchema,
	},
	{
		"name":         "MessageHeaderHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "MolecularSequence",
		"schema": commonSchema,
	},
	{
		"name":         "MolecularSequenceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Organization",
		"schema": commonSchema,
	},
	{
		"name":         "OrganizationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "OrganizationAffiliation",
		"schema": commonSchema,
	},
	{
		"name":         "OrganizationAffiliationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Patient",
		"schema": commonSchema,
	},
	{
		"name":         "PatientHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Person",
		"schema": commonSchema,
	},
	{
		"name":         "PersonHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Practitioner",
		"schema": commonSchema,
	},
	{
		"name":         "PractitionerHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "RelatedPerson",
		"schema": commonSchema,
	},
	{
		"name":         "RelatedPersonHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "PractitionerRole",
		"schema": commonSchema,
	},
	{
		"name":         "PractitionerRoleHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "PaymentNotice",
		"schema": commonSchema,
	},
	{
		"name":         "PaymentNoticeHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "PaymentReconciliation",
		"schema": commonSchema,
	},
	{
		"name":         "PaymentReconciliationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "PlanDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "PlanDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Provenance",
		"schema": commonSchema,
	},
	{
		"name":         "ProvenanceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Questionnaire",
		"schema": commonSchema,
	},
	{
		"name":         "QuestionnaireHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "QuestionnaireResponse",
		"schema": commonSchema,
	},
	{
		"name":         "QuestionnaireResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "RequestGroup",
		"schema": commonSchema,
	},
	{
		"name":         "RequestGroupHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ResearchDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ResearchElementDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchElementDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ResearchStudy",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchStudyHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "ResearchSubject",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchSubjectHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "RiskEvidenceSynthesis",
		"schema": commonSchema,
	},
	{
		"name":         "RiskEvidenceSynthesisHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Schedule",
		"schema": commonSchema,
	},
	{
		"name":         "ScheduleHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Slot",
		"schema": commonSchema,
	},
	{
		"name":         "SlotHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Specimen",
		"schema": commonSchema,
	},
	{
		"name":         "SpecimenHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "SpecimenDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "SpecimenDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Subscription",
		"schema": commonSchema,
	},
	{
		"name":         "SubscriptionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Substance",
		"schema": commonSchema,
	},
	{
		"name":         "SubstanceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "SubstanceSpecification",
		"schema": commonSchema,
	},
	{
		"name":         "SubstanceSpecificationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "Task",
		"schema": commonSchema,
	},
	{
		"name":         "TaskHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "TestReport",
		"schema": commonSchema,
	},
	{
		"name":         "TestReportHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "TestScript",
		"schema": commonSchema,
	},
	{
		"name":         "TestScriptHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
	{
		"name":   "VerificationResult",
		"schema": commonSchema,
	},
	{
		"name":         "VerificationResultHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "versionId"},
	},
}
