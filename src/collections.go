package main

// Common schema applied to multiple collections
var commonSchema = []map[string]interface{}{
	{"name": "lastUpdated", "type": "date", "required": true},
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
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ActivityDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ActivityDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "AdverseEvent",
		"schema": commonSchema,
	},
	{
		"name":         "AdverseEventHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "AllergyIntolerance",
		"schema": commonSchema,
	},
	{
		"name":         "AllergyIntoleranceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Condition",
		"schema": commonSchema,
	},
	{
		"name":         "ConditionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DeviceRequest",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DiagnosticReport",
		"schema": commonSchema,
	},
	{
		"name":         "DiagnosticReportHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "List",
		"schema": commonSchema,
	},
	{
		"name":         "ListHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Medication",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicationAdministration",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationAdministrationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicationDispense",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationDispenseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicationRequest",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicationStatement",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationStatementHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Observation",
		"schema": commonSchema,
	},
	{
		"name":         "ObservationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Procedure",
		"schema": commonSchema,
	},
	{
		"name":         "ProcedureHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ServiceRequest",
		"schema": commonSchema,
	},
	{
		"name":         "ServiceRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "CarePlan",
		"schema": commonSchema,
	},
	{
		"name":         "CarePlanHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "CareTeam",
		"schema": commonSchema,
	},
	{
		"name":         "CareTeamHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ClinicalImpression",
		"schema": commonSchema,
	},
	{
		"name":         "ClinicalImpressionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Consent",
		"schema": commonSchema,
	},
	{
		"name":         "ConsentHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Encounter",
		"schema": commonSchema,
	},
	{
		"name":         "EncounterHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "EpisodeOfCare",
		"schema": commonSchema,
	},
	{
		"name":         "EpisodeOfCareHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Flag",
		"schema": commonSchema,
	},
	{
		"name":         "FlagHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Immunization",
		"schema": commonSchema,
	},
	{
		"name":         "ImmunizationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "RiskAssessment",
		"schema": commonSchema,
	},
	{
		"name":         "RiskAssessmentHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "SupplyRequest",
		"schema": commonSchema,
	},
	{
		"name":         "SupplyRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DetectedIssue",
		"schema": commonSchema,
	},
	{
		"name":         "DetectedIssueHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DocumentManifest",
		"schema": commonSchema,
	},
	{
		"name":         "DocumentManifestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DocumentReference",
		"schema": commonSchema,
	},
	{
		"name":         "DocumentReferenceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Goal",
		"schema": commonSchema,
	},
	{
		"name":         "GoalHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ImagingStudy",
		"schema": commonSchema,
	},
	{
		"name":         "ImagingStudyHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "NutritionOrder",
		"schema": commonSchema,
	},
	{
		"name":         "NutritionOrderHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "SupplyDelivery",
		"schema": commonSchema,
	},
	{
		"name":         "SupplyDeliveryHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "VisionPrescription",
		"schema": commonSchema,
	},
	{
		"name":         "VisionPrescriptionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DeviceUseStatement",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceUseStatementHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Appointment",
		"schema": commonSchema,
	},
	{
		"name":         "AppointmentHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "AppointmentResponse",
		"schema": commonSchema,
	},
	{
		"name":         "AppointmentResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "AuditEvent",
		"schema": commonSchema,
	},
	{
		"name":         "AuditEventHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Basic",
		"schema": commonSchema,
	},
	{
		"name":         "BasicHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "BodyStructure",
		"schema": commonSchema,
	},
	{
		"name":         "BodyStructureHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Bundle",
		"schema": commonSchema,
	},
	{
		"name":         "BundleHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "CapabilityStatement",
		"schema": commonSchema,
	},
	{
		"name":         "CapabilityStatementHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "CodeSystem",
		"schema": commonSchema,
	},
	{
		"name":         "CodeSystemHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ConceptMap",
		"schema": commonSchema,
	},
	{
		"name":         "ConceptMapHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "GraphDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "GraphDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ImplementationGuide",
		"schema": commonSchema,
	},
	{
		"name":         "ImplementationGuideHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MessageDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "MessageDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "NamingSystem",
		"schema": commonSchema,
	},
	{
		"name":         "NamingSystemHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "OperationDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "OperationDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "SearchParameter",
		"schema": commonSchema,
	},
	{
		"name":         "SearchParameterHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "StructureDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "StructureDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "StructureMap",
		"schema": commonSchema,
	},
	{
		"name":         "StructureMapHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "TerminologyCapabilities",
		"schema": commonSchema,
	},
	{
		"name":         "TerminologyCapabilitiesHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ValueSet",
		"schema": commonSchema,
	},
	{
		"name":         "ValueSetHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ChargeItem",
		"schema": commonSchema,
	},
	{
		"name":         "ChargeItemHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ChargeItemDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ChargeItemDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Claim",
		"schema": commonSchema,
	},
	{
		"name":         "ClaimHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ClaimResponse",
		"schema": commonSchema,
	},
	{
		"name":         "ClaimResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Communication",
		"schema": commonSchema,
	},
	{
		"name":         "CommunicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "CommunicationRequest",
		"schema": commonSchema,
	},
	{
		"name":         "CommunicationRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Contract",
		"schema": commonSchema,
	},
	{
		"name":         "ContractHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Coverage",
		"schema": commonSchema,
	},
	{
		"name":         "CoverageHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "CoverageEligibilityRequest",
		"schema": commonSchema,
	},
	{
		"name":         "CoverageEligibilityRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "CoverageEligibilityResponse",
		"schema": commonSchema,
	},
	{
		"name":         "CoverageEligibilityResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Device",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DeviceDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "DeviceMetric",
		"schema": commonSchema,
	},
	{
		"name":         "DeviceMetricHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "EffectEvidenceSynthesis",
		"schema": commonSchema,
	},
	{
		"name":         "EffectEvidenceSynthesisHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Endpoint",
		"schema": commonSchema,
	},
	{
		"name":         "EndpointHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "EnrollmentRequest",
		"schema": commonSchema,
	},
	{
		"name":         "EnrollmentRequestHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "EnrollmentResponse",
		"schema": commonSchema,
	},
	{
		"name":         "EnrollmentResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "EventDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "EventDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Evidence",
		"schema": commonSchema,
	},
	{
		"name":         "EvidenceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "EvidenceVariable",
		"schema": commonSchema,
	},
	{
		"name":         "EvidenceVariableHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ExampleScenario",
		"schema": commonSchema,
	},
	{
		"name":         "ExampleScenarioHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ExplanationOfBenefit",
		"schema": commonSchema,
	},
	{
		"name":         "ExplanationOfBenefitHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Group",
		"schema": commonSchema,
	},
	{
		"name":         "GroupHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "GuidanceResponse",
		"schema": commonSchema,
	},
	{
		"name":         "GuidanceResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "HealthcareService",
		"schema": commonSchema,
	},
	{
		"name":         "HealthcareServiceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ImmunizationEvaluation",
		"schema": commonSchema,
	},
	{
		"name":         "ImmunizationEvaluationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ImmunizationRecommendation",
		"schema": commonSchema,
	},
	{
		"name":         "ImmunizationRecommendationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "InsurancePlan",
		"schema": commonSchema,
	},
	{
		"name":         "InsurancePlanHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Invoice",
		"schema": commonSchema,
	},
	{
		"name":         "InvoiceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Library",
		"schema": commonSchema,
	},
	{
		"name":         "LibraryHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Linkage",
		"schema": commonSchema,
	},
	{
		"name":         "LinkageHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Location",
		"schema": commonSchema,
	},
	{
		"name":         "LocationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Measure",
		"schema": commonSchema,
	},
	{
		"name":         "MeasureHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MeasureReport",
		"schema": commonSchema,
	},
	{
		"name":         "MeasureReportHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Media",
		"schema": commonSchema,
	},
	{
		"name":         "MediaHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicationKnowledge",
		"schema": commonSchema,
	},
	{
		"name":         "MedicationKnowledgeHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProduct",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProductAuthorization",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductAuthorizationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProductContraindication",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductContraindicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProductIndication",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductIndicationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProductInteraction",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductInteractionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProductPackaged",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductPackagedHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProductPharmaceutical",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductPharmaceuticalHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MedicinalProductUndesirableEffect",
		"schema": commonSchema,
	},
	{
		"name":         "MedicinalProductUndesirableEffectHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MessageHeader",
		"schema": commonSchema,
	},
	{
		"name":         "MessageHeaderHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "MolecularSequence",
		"schema": commonSchema,
	},
	{
		"name":         "MolecularSequenceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Organization",
		"schema": commonSchema,
	},
	{
		"name":         "OrganizationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "OrganizationAffiliation",
		"schema": commonSchema,
	},
	{
		"name":         "OrganizationAffiliationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Patient",
		"schema": commonSchema,
	},
	{
		"name":         "PatientHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Person",
		"schema": commonSchema,
	},
	{
		"name":         "PersonHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Practitioner",
		"schema": commonSchema,
	},
	{
		"name":         "PractitionerHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "RelatedPerson",
		"schema": commonSchema,
	},
	{
		"name":         "RelatedPersonHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "PractitionerRole",
		"schema": commonSchema,
	},
	{
		"name":         "PractitionerRoleHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "PaymentNotice",
		"schema": commonSchema,
	},
	{
		"name":         "PaymentNoticeHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "PaymentReconciliation",
		"schema": commonSchema,
	},
	{
		"name":         "PaymentReconciliationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "PlanDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "PlanDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Provenance",
		"schema": commonSchema,
	},
	{
		"name":         "ProvenanceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Questionnaire",
		"schema": commonSchema,
	},
	{
		"name":         "QuestionnaireHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "QuestionnaireResponse",
		"schema": commonSchema,
	},
	{
		"name":         "QuestionnaireResponseHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "RequestGroup",
		"schema": commonSchema,
	},
	{
		"name":         "RequestGroupHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ResearchDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ResearchElementDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchElementDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ResearchStudy",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchStudyHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "ResearchSubject",
		"schema": commonSchema,
	},
	{
		"name":         "ResearchSubjectHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "RiskEvidenceSynthesis",
		"schema": commonSchema,
	},
	{
		"name":         "RiskEvidenceSynthesisHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Schedule",
		"schema": commonSchema,
	},
	{
		"name":         "ScheduleHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Slot",
		"schema": commonSchema,
	},
	{
		"name":         "SlotHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Specimen",
		"schema": commonSchema,
	},
	{
		"name":         "SpecimenHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "SpecimenDefinition",
		"schema": commonSchema,
	},
	{
		"name":         "SpecimenDefinitionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Subscription",
		"schema": commonSchema,
	},
	{
		"name":         "SubscriptionHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Substance",
		"schema": commonSchema,
	},
	{
		"name":         "SubstanceHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "SubstanceSpecification",
		"schema": commonSchema,
	},
	{
		"name":         "SubstanceSpecificationHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "Task",
		"schema": commonSchema,
	},
	{
		"name":         "TaskHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "TestReport",
		"schema": commonSchema,
	},
	{
		"name":         "TestReportHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "TestScript",
		"schema": commonSchema,
	},
	{
		"name":         "TestScriptHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
	{
		"name":   "VerificationResult",
		"schema": commonSchema,
	},
	{
		"name":         "VerificationResultHistory",
		"schema":       commonSchema,
		"compositeKey": []string{"id", "lastUpdated"},
	},
}
