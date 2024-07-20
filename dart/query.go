package main

import (
  "fmt"
  "github.com/pocketbase/pocketbase"
  "github.com/pocketbase/pocketbase/models"
)

func queryResources(app *pocketbase.PocketBase, resourceType string, searchParams map[string]string) ([]*models.Record, error) {
  collection, err := app.Dao().FindCollectionByNameOrId(resourceType)
  if err != nil {
    return nil, fmt.Errorf("collection not found: %w", err)
  }

  query := app.Dao().RecordQuery(resourceType)

  // Add search conditions based on search parameters
  for param, value := range searchParams {
    switch param {
    case "_text":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._text') = ?", param), value)
    case "_content":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._content') = ?", param), value)
    case "_id":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._id') = ?", param), value)
    case "_lastUpdated":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._lastUpdated') = ?", param), value)
    case "_profile":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._profile') = ?", param), value)
    case "_query":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._query') = ?", param), value)
    case "_security":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._security') = ?", param), value)
    case "_source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._source') = ?", param), value)
    case "_tag":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$._tag') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "owner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.owner') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.period') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "actuality":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.actuality') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "event":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.event') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "recorder":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recorder') = ?", param), value)
    case "resultingcondition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.resultingcondition') = ?", param), value)
    case "seriousness":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.seriousness') = ?", param), value)
    case "severity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.severity') = ?", param), value)
    case "study":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.study') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "substance":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.substance') = ?", param), value)
    case "asserter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.asserter') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "clinical-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.clinical-status') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "criticality":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.criticality') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "last-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.last-date') = ?", param), value)
    case "manifestation":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.manifestation') = ?", param), value)
    case "onset":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.onset') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "recorder":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recorder') = ?", param), value)
    case "route":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.route') = ?", param), value)
    case "severity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.severity') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "verification-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.verification-status') = ?", param), value)
    case "actor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.actor') = ?", param), value)
    case "appointment-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.appointment-type') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "part-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-status') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "practitioner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.practitioner') = ?", param), value)
    case "reason-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-code') = ?", param), value)
    case "reason-reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-reference') = ?", param), value)
    case "service-category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-category') = ?", param), value)
    case "service-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-type') = ?", param), value)
    case "slot":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.slot') = ?", param), value)
    case "specialty":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specialty') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "supporting-info":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.supporting-info') = ?", param), value)
    case "actor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.actor') = ?", param), value)
    case "appointment":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.appointment') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "part-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-status') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "practitioner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.practitioner') = ?", param), value)
    case "action":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.action') = ?", param), value)
    case "address":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address') = ?", param), value)
    case "agent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.agent') = ?", param), value)
    case "agent-name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.agent-name') = ?", param), value)
    case "agent-role":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.agent-role') = ?", param), value)
    case "altid":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.altid') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "entity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.entity') = ?", param), value)
    case "entity-name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.entity-name') = ?", param), value)
    case "entity-role":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.entity-role') = ?", param), value)
    case "entity-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.entity-type') = ?", param), value)
    case "outcome":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.outcome') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "policy":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.policy') = ?", param), value)
    case "site":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.site') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "subtype":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subtype') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "morphology":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.morphology') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "composition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composition') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "message":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.message') = ?", param), value)
    case "timestamp":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.timestamp') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "fhirversion":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.fhirversion') = ?", param), value)
    case "format":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.format') = ?", param), value)
    case "guide":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.guide') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "mode":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.mode') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "resource":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.resource') = ?", param), value)
    case "resource-profile":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.resource-profile') = ?", param), value)
    case "security-service":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.security-service') = ?", param), value)
    case "software":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.software') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "supported-profile":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.supported-profile') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "activity-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.activity-code') = ?", param), value)
    case "activity-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.activity-date') = ?", param), value)
    case "activity-reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.activity-reference') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "care-team":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.care-team') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "condition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.condition') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "goal":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.goal') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "intent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intent') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "replaces":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.replaces') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "participant":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participant') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "account":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.account') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "entered-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.entered-date') = ?", param), value)
    case "enterer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.enterer') = ?", param), value)
    case "factor-override":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.factor-override') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "occurrence":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.occurrence') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "performer-actor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer-actor') = ?", param), value)
    case "performer-function":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer-function') = ?", param), value)
    case "performing-organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performing-organization') = ?", param), value)
    case "price-override":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.price-override') = ?", param), value)
    case "quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.quantity') = ?", param), value)
    case "requesting-organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requesting-organization') = ?", param), value)
    case "service":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "care-team":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.care-team') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "detail-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.detail-udi') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "enterer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.enterer') = ?", param), value)
    case "facility":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.facility') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "insurer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.insurer') = ?", param), value)
    case "item-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.item-udi') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "payee":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payee') = ?", param), value)
    case "priority":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.priority') = ?", param), value)
    case "procedure-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.procedure-udi') = ?", param), value)
    case "provider":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.provider') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subdetail-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subdetail-udi') = ?", param), value)
    case "use":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.use') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "disposition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.disposition') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "insurer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.insurer') = ?", param), value)
    case "outcome":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.outcome') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "payment-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payment-date') = ?", param), value)
    case "request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.request') = ?", param), value)
    case "requestor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requestor') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "use":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.use') = ?", param), value)
    case "assessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.assessor') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "finding-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.finding-code') = ?", param), value)
    case "finding-ref":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.finding-ref') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "investigation":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.investigation') = ?", param), value)
    case "previous":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.previous') = ?", param), value)
    case "problem":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.problem') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "supporting-info":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.supporting-info') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "content-mode":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.content-mode') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "language":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.language') = ?", param), value)
    case "supplements":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.supplements') = ?", param), value)
    case "system":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.system') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "medium":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.medium') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "received":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.received') = ?", param), value)
    case "recipient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recipient') = ?", param), value)
    case "sender":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.sender') = ?", param), value)
    case "sent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.sent') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "authored":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authored') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "group-identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.group-identifier') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "medium":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.medium') = ?", param), value)
    case "occurrence":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.occurrence') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "priority":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.priority') = ?", param), value)
    case "recipient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recipient') = ?", param), value)
    case "replaces":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.replaces') = ?", param), value)
    case "requester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requester') = ?", param), value)
    case "sender":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.sender') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "resource":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.resource') = ?", param), value)
    case "attester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.attester') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "confidentiality":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.confidentiality') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "entry":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.entry') = ?", param), value)
    case "period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.period') = ?", param), value)
    case "related-id":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.related-id') = ?", param), value)
    case "related-ref":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.related-ref') = ?", param), value)
    case "section":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.section') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "dependson":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.dependson') = ?", param), value)
    case "other":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.other') = ?", param), value)
    case "product":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.product') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "source-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source-code') = ?", param), value)
    case "source-system":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source-system') = ?", param), value)
    case "source-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source-uri') = ?", param), value)
    case "target":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target') = ?", param), value)
    case "target-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-code') = ?", param), value)
    case "target-system":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-system') = ?", param), value)
    case "target-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-uri') = ?", param), value)
    case "abatement-age":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.abatement-age') = ?", param), value)
    case "abatement-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.abatement-date') = ?", param), value)
    case "abatement-string":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.abatement-string') = ?", param), value)
    case "asserter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.asserter') = ?", param), value)
    case "body-site":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.body-site') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "clinical-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.clinical-status') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "evidence":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.evidence') = ?", param), value)
    case "evidence-detail":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.evidence-detail') = ?", param), value)
    case "onset-age":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.onset-age') = ?", param), value)
    case "onset-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.onset-date') = ?", param), value)
    case "onset-info":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.onset-info') = ?", param), value)
    case "recorded-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recorded-date') = ?", param), value)
    case "severity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.severity') = ?", param), value)
    case "stage":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.stage') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "verification-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.verification-status') = ?", param), value)
    case "action":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.action') = ?", param), value)
    case "actor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.actor') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "consentor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.consentor') = ?", param), value)
    case "data":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.data') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.period') = ?", param), value)
    case "purpose":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.purpose') = ?", param), value)
    case "scope":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.scope') = ?", param), value)
    case "security-label":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.security-label') = ?", param), value)
    case "source-reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source-reference') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "authority":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authority') = ?", param), value)
    case "domain":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.domain') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "instantiates":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates') = ?", param), value)
    case "issued":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.issued') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "signer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.signer') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "beneficiary":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.beneficiary') = ?", param), value)
    case "class-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.class-type') = ?", param), value)
    case "class-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.class-value') = ?", param), value)
    case "dependent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.dependent') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "payor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payor') = ?", param), value)
    case "policy-holder":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.policy-holder') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subscriber":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subscriber') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "enterer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.enterer') = ?", param), value)
    case "facility":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.facility') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "provider":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.provider') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "disposition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.disposition') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "insurer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.insurer') = ?", param), value)
    case "outcome":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.outcome') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.request') = ?", param), value)
    case "requestor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requestor') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "identified":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identified') = ?", param), value)
    case "implicated":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.implicated') = ?", param), value)
    case "device-name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.device-name') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "manufacturer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.manufacturer') = ?", param), value)
    case "model":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.model') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "udi-carrier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.udi-carrier') = ?", param), value)
    case "udi-di":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.udi-di') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "parent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.parent') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "parent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.parent') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "authored-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authored-on') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "device":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.device') = ?", param), value)
    case "event-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.event-date') = ?", param), value)
    case "group-identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.group-identifier') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "insurance":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.insurance') = ?", param), value)
    case "intent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intent') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "prior-request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.prior-request') = ?", param), value)
    case "requester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requester') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "device":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.device') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "conclusion":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.conclusion') = ?", param), value)
    case "issued":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.issued') = ?", param), value)
    case "media":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.media') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "result":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.result') = ?", param), value)
    case "results-interpreter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.results-interpreter') = ?", param), value)
    case "specimen":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specimen') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "item":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.item') = ?", param), value)
    case "recipient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recipient') = ?", param), value)
    case "related-id":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.related-id') = ?", param), value)
    case "related-ref":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.related-ref') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "authenticator":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authenticator') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "contenttype":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.contenttype') = ?", param), value)
    case "custodian":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.custodian') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "event":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.event') = ?", param), value)
    case "facility":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.facility') = ?", param), value)
    case "format":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.format') = ?", param), value)
    case "language":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.language') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.period') = ?", param), value)
    case "related":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.related') = ?", param), value)
    case "relatesto":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.relatesto') = ?", param), value)
    case "relation":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.relation') = ?", param), value)
    case "security-label":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.security-label') = ?", param), value)
    case "setting":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.setting') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "relationship":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.relationship') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "account":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.account') = ?", param), value)
    case "appointment":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.appointment') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "class":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.class') = ?", param), value)
    case "diagnosis":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.diagnosis') = ?", param), value)
    case "episode-of-care":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.episode-of-care') = ?", param), value)
    case "length":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.length') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "location-period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location-period') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "participant":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participant') = ?", param), value)
    case "participant-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participant-type') = ?", param), value)
    case "practitioner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.practitioner') = ?", param), value)
    case "reason-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-code') = ?", param), value)
    case "reason-reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-reference') = ?", param), value)
    case "service-provider":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-provider') = ?", param), value)
    case "special-arrangement":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.special-arrangement') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "connection-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.connection-type') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "payload-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payload-type') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.request') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "care-manager":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.care-manager') = ?", param), value)
    case "condition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.condition') = ?", param), value)
    case "incoming-referral":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.incoming-referral') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "care-team":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.care-team') = ?", param), value)
    case "claim":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.claim') = ?", param), value)
    case "coverage":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.coverage') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "detail-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.detail-udi') = ?", param), value)
    case "disposition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.disposition') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "enterer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.enterer') = ?", param), value)
    case "facility":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.facility') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "item-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.item-udi') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "payee":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payee') = ?", param), value)
    case "procedure-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.procedure-udi') = ?", param), value)
    case "provider":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.provider') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subdetail-udi":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subdetail-udi') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "relationship":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.relationship') = ?", param), value)
    case "sex":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.sex') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "achievement-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.achievement-status') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "lifecycle-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.lifecycle-status') = ?", param), value)
    case "start-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.start-date') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "target-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-date') = ?", param), value)
    case "start":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.start') = ?", param), value)
    case "actual":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.actual') = ?", param), value)
    case "characteristic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.characteristic') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "exclude":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.exclude') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "managing-entity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.managing-entity') = ?", param), value)
    case "member":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.member') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.value') = ?", param), value)
    case "characteristic-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.characteristic-value') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.request') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "characteristic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.characteristic') = ?", param), value)
    case "coverage-area":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.coverage-area') = ?", param), value)
    case "endpoint":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.endpoint') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "program":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.program') = ?", param), value)
    case "service-category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-category') = ?", param), value)
    case "service-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-type') = ?", param), value)
    case "specialty":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specialty') = ?", param), value)
    case "basedon":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.basedon') = ?", param), value)
    case "bodysite":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.bodysite') = ?", param), value)
    case "dicom-class":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.dicom-class') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "endpoint":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.endpoint') = ?", param), value)
    case "instance":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instance') = ?", param), value)
    case "interpreter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.interpreter') = ?", param), value)
    case "modality":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.modality') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "reason":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason') = ?", param), value)
    case "referrer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.referrer') = ?", param), value)
    case "series":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.series') = ?", param), value)
    case "started":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.started') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "lot-number":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.lot-number') = ?", param), value)
    case "manufacturer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.manufacturer') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "reaction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reaction') = ?", param), value)
    case "reaction-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reaction-date') = ?", param), value)
    case "reason-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-code') = ?", param), value)
    case "reason-reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-reference') = ?", param), value)
    case "series":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.series') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "status-reason":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status-reason') = ?", param), value)
    case "target-disease":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-disease') = ?", param), value)
    case "vaccine-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.vaccine-code') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "dose-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.dose-status') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "immunization-event":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.immunization-event') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "target-disease":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-disease') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "information":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.information') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "support":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.support') = ?", param), value)
    case "target-disease":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-disease') = ?", param), value)
    case "vaccine-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.vaccine-type') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "experimental":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.experimental') = ?", param), value)
    case "global":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.global') = ?", param), value)
    case "resource":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.resource') = ?", param), value)
    case "address":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address') = ?", param), value)
    case "address-city":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-city') = ?", param), value)
    case "address-country":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-country') = ?", param), value)
    case "address-postalcode":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-postalcode') = ?", param), value)
    case "address-state":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-state') = ?", param), value)
    case "address-use":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-use') = ?", param), value)
    case "administered-by":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.administered-by') = ?", param), value)
    case "endpoint":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.endpoint') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "owned-by":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.owned-by') = ?", param), value)
    case "phonetic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.phonetic') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "account":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.account') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "issuer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.issuer') = ?", param), value)
    case "participant":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participant') = ?", param), value)
    case "participant-role":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participant-role') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "recipient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recipient') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "totalgross":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.totalgross') = ?", param), value)
    case "totalnet":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.totalnet') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "content-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.content-type') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "item":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.item') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "empty-reason":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.empty-reason') = ?", param), value)
    case "item":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.item') = ?", param), value)
    case "notes":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.notes') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "address":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address') = ?", param), value)
    case "address-city":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-city') = ?", param), value)
    case "address-country":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-country') = ?", param), value)
    case "address-postalcode":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-postalcode') = ?", param), value)
    case "address-state":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-state') = ?", param), value)
    case "address-use":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-use') = ?", param), value)
    case "endpoint":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.endpoint') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "near":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.near') = ?", param), value)
    case "operational-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.operational-status') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "partof":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.partof') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "evaluated-resource":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.evaluated-resource') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "measure":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.measure') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.period') = ?", param), value)
    case "reporter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reporter') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "device":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.device') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "modality":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.modality') = ?", param), value)
    case "operator":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.operator') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "site":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.site') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "view":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.view') = ?", param), value)
    case "expiration-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.expiration-date') = ?", param), value)
    case "form":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.form') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "ingredient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.ingredient') = ?", param), value)
    case "ingredient-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.ingredient-code') = ?", param), value)
    case "lot-number":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.lot-number') = ?", param), value)
    case "manufacturer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.manufacturer') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "device":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.device') = ?", param), value)
    case "effective-time":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective-time') = ?", param), value)
    case "medication":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.medication') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "reason-given":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-given') = ?", param), value)
    case "reason-not-given":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-not-given') = ?", param), value)
    case "request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.request') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "destination":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.destination') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "prescription":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.prescription') = ?", param), value)
    case "receiver":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.receiver') = ?", param), value)
    case "responsibleparty":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.responsibleparty') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "whenhandedover":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.whenhandedover') = ?", param), value)
    case "whenprepared":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.whenprepared') = ?", param), value)
    case "classification":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.classification') = ?", param), value)
    case "classification-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.classification-type') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "doseform":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.doseform') = ?", param), value)
    case "ingredient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.ingredient') = ?", param), value)
    case "ingredient-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.ingredient-code') = ?", param), value)
    case "manufacturer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.manufacturer') = ?", param), value)
    case "monitoring-program-name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.monitoring-program-name') = ?", param), value)
    case "monitoring-program-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.monitoring-program-type') = ?", param), value)
    case "monograph":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.monograph') = ?", param), value)
    case "monograph-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.monograph-type') = ?", param), value)
    case "source-cost":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source-cost') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "authoredon":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authoredon') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "intended-dispenser":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intended-dispenser') = ?", param), value)
    case "intended-performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intended-performer') = ?", param), value)
    case "intended-performertype":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intended-performertype') = ?", param), value)
    case "intent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intent') = ?", param), value)
    case "priority":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.priority') = ?", param), value)
    case "requester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requester') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "name-language":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name-language') = ?", param), value)
    case "country":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.country') = ?", param), value)
    case "holder":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.holder') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "route":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.route') = ?", param), value)
    case "target-species":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target-species') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "event":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.event') = ?", param), value)
    case "focus":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.focus') = ?", param), value)
    case "parent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.parent') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "destination":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.destination') = ?", param), value)
    case "destination-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.destination-uri') = ?", param), value)
    case "enterer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.enterer') = ?", param), value)
    case "event":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.event') = ?", param), value)
    case "focus":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.focus') = ?", param), value)
    case "receiver":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.receiver') = ?", param), value)
    case "response-id":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.response-id') = ?", param), value)
    case "responsible":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.responsible') = ?", param), value)
    case "sender":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.sender') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "source-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source-uri') = ?", param), value)
    case "target":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target') = ?", param), value)
    case "chromosome":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.chromosome') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "referenceseqid":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.referenceseqid') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "variant-end":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.variant-end') = ?", param), value)
    case "variant-start":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.variant-start') = ?", param), value)
    case "window-end":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.window-end') = ?", param), value)
    case "window-start":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.window-start') = ?", param), value)
    case "chromosome-variant-coordinate":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.chromosome-variant-coordinate') = ?", param), value)
    case "chromosome-window-coordinate":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.chromosome-window-coordinate') = ?", param), value)
    case "referenceseqid-variant-coordinate":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.referenceseqid-variant-coordinate') = ?", param), value)
    case "referenceseqid-window-coordinate":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.referenceseqid-window-coordinate') = ?", param), value)
    case "contact":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.contact') = ?", param), value)
    case "id-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.id-type') = ?", param), value)
    case "kind":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.kind') = ?", param), value)
    case "period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.period') = ?", param), value)
    case "responsible":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.responsible') = ?", param), value)
    case "telecom":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.telecom') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.value') = ?", param), value)
    case "additive":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.additive') = ?", param), value)
    case "datetime":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.datetime') = ?", param), value)
    case "formula":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.formula') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "oraldiet":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.oraldiet') = ?", param), value)
    case "provider":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.provider') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "supplement":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.supplement') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "combo-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.combo-code') = ?", param), value)
    case "combo-data-absent-reason":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.combo-data-absent-reason') = ?", param), value)
    case "combo-value-concept":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.combo-value-concept') = ?", param), value)
    case "combo-value-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.combo-value-quantity') = ?", param), value)
    case "component-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.component-code') = ?", param), value)
    case "component-data-absent-reason":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.component-data-absent-reason') = ?", param), value)
    case "component-value-concept":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.component-value-concept') = ?", param), value)
    case "component-value-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.component-value-quantity') = ?", param), value)
    case "data-absent-reason":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.data-absent-reason') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "device":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.device') = ?", param), value)
    case "focus":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.focus') = ?", param), value)
    case "has-member":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.has-member') = ?", param), value)
    case "method":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.method') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "specimen":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specimen') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "value-concept":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.value-concept') = ?", param), value)
    case "value-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.value-date') = ?", param), value)
    case "value-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.value-quantity') = ?", param), value)
    case "value-string":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.value-string') = ?", param), value)
    case "code-value-concept":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code-value-concept') = ?", param), value)
    case "code-value-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code-value-date') = ?", param), value)
    case "code-value-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code-value-quantity') = ?", param), value)
    case "code-value-string":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code-value-string') = ?", param), value)
    case "combo-code-value-concept":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.combo-code-value-concept') = ?", param), value)
    case "combo-code-value-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.combo-code-value-quantity') = ?", param), value)
    case "component-code-value-concept":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.component-code-value-concept') = ?", param), value)
    case "component-code-value-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.component-code-value-quantity') = ?", param), value)
    case "base":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.base') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "input-profile":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.input-profile') = ?", param), value)
    case "instance":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instance') = ?", param), value)
    case "kind":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.kind') = ?", param), value)
    case "output-profile":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.output-profile') = ?", param), value)
    case "system":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.system') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "address":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address') = ?", param), value)
    case "address-city":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-city') = ?", param), value)
    case "address-country":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-country') = ?", param), value)
    case "address-postalcode":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-postalcode') = ?", param), value)
    case "address-state":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-state') = ?", param), value)
    case "address-use":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-use') = ?", param), value)
    case "endpoint":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.endpoint') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "partof":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.partof') = ?", param), value)
    case "phonetic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.phonetic') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "email":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.email') = ?", param), value)
    case "endpoint":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.endpoint') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "network":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.network') = ?", param), value)
    case "participating-organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participating-organization') = ?", param), value)
    case "phone":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.phone') = ?", param), value)
    case "primary-organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.primary-organization') = ?", param), value)
    case "role":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.role') = ?", param), value)
    case "service":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service') = ?", param), value)
    case "specialty":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specialty') = ?", param), value)
    case "telecom":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.telecom') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "address":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address') = ?", param), value)
    case "address-city":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-city') = ?", param), value)
    case "address-country":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-country') = ?", param), value)
    case "address-postalcode":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-postalcode') = ?", param), value)
    case "address-state":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-state') = ?", param), value)
    case "address-use":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.address-use') = ?", param), value)
    case "birthdate":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.birthdate') = ?", param), value)
    case "death-date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.death-date') = ?", param), value)
    case "deceased":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.deceased') = ?", param), value)
    case "email":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.email') = ?", param), value)
    case "family":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.family') = ?", param), value)
    case "gender":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.gender') = ?", param), value)
    case "general-practitioner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.general-practitioner') = ?", param), value)
    case "given":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.given') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "language":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.language') = ?", param), value)
    case "link":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.link') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "phone":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.phone') = ?", param), value)
    case "phonetic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.phonetic') = ?", param), value)
    case "telecom":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.telecom') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "payment-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payment-status') = ?", param), value)
    case "provider":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.provider') = ?", param), value)
    case "request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.request') = ?", param), value)
    case "response":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.response') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "created":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.created') = ?", param), value)
    case "disposition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.disposition') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "outcome":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.outcome') = ?", param), value)
    case "payment-issuer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payment-issuer') = ?", param), value)
    case "request":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.request') = ?", param), value)
    case "requestor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requestor') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "link":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.link') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "practitioner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.practitioner') = ?", param), value)
    case "relatedperson":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.relatedperson') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "definition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.definition') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "communication":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.communication') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "endpoint":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.endpoint') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "organization":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.organization') = ?", param), value)
    case "practitioner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.practitioner') = ?", param), value)
    case "role":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.role') = ?", param), value)
    case "service":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service') = ?", param), value)
    case "specialty":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specialty') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "reason-code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-code') = ?", param), value)
    case "reason-reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reason-reference') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "agent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.agent') = ?", param), value)
    case "agent-role":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.agent-role') = ?", param), value)
    case "agent-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.agent-type') = ?", param), value)
    case "entity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.entity') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "recorded":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.recorded') = ?", param), value)
    case "signature-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.signature-type') = ?", param), value)
    case "target":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target') = ?", param), value)
    case "when":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.when') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "definition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.definition') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject-type') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "authored":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authored') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "questionnaire":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.questionnaire') = ?", param), value)
    case "source":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.source') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "relationship":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.relationship') = ?", param), value)
    case "author":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.author') = ?", param), value)
    case "authored":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authored') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "group-identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.group-identifier') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "intent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intent') = ?", param), value)
    case "participant":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participant') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "priority":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.priority') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "composed-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.composed-of') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "depends-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.depends-on') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "predecessor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.predecessor') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "successor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.successor') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "topic":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.topic') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "focus":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.focus') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "keyword":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.keyword') = ?", param), value)
    case "location":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.location') = ?", param), value)
    case "partof":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.partof') = ?", param), value)
    case "principalinvestigator":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.principalinvestigator') = ?", param), value)
    case "protocol":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.protocol') = ?", param), value)
    case "site":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.site') = ?", param), value)
    case "sponsor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.sponsor') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "individual":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.individual') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "study":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.study') = ?", param), value)
    case "condition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.condition') = ?", param), value)
    case "method":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.method') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "probability":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.probability') = ?", param), value)
    case "risk":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.risk') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "effective":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.effective') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "active":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.active') = ?", param), value)
    case "actor":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.actor') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "service-category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-category') = ?", param), value)
    case "service-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-type') = ?", param), value)
    case "specialty":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specialty') = ?", param), value)
    case "base":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.base') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "component":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.component') = ?", param), value)
    case "derived-from":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derived-from') = ?", param), value)
    case "target":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "authored":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authored') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "body-site":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.body-site') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "instantiates-canonical":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-canonical') = ?", param), value)
    case "instantiates-uri":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.instantiates-uri') = ?", param), value)
    case "intent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intent') = ?", param), value)
    case "occurrence":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.occurrence') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "performer-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer-type') = ?", param), value)
    case "priority":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.priority') = ?", param), value)
    case "replaces":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.replaces') = ?", param), value)
    case "requester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requester') = ?", param), value)
    case "requisition":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requisition') = ?", param), value)
    case "specimen":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specimen') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "appointment-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.appointment-type') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "schedule":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.schedule') = ?", param), value)
    case "service-category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-category') = ?", param), value)
    case "service-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.service-type') = ?", param), value)
    case "specialty":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.specialty') = ?", param), value)
    case "start":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.start') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "accession":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.accession') = ?", param), value)
    case "bodysite":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.bodysite') = ?", param), value)
    case "collected":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.collected') = ?", param), value)
    case "collector":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.collector') = ?", param), value)
    case "container":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.container') = ?", param), value)
    case "container-id":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.container-id') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "parent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.parent') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "container":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.container') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "abstract":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.abstract') = ?", param), value)
    case "base":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.base') = ?", param), value)
    case "base-path":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.base-path') = ?", param), value)
    case "derivation":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.derivation') = ?", param), value)
    case "experimental":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.experimental') = ?", param), value)
    case "ext-context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.ext-context') = ?", param), value)
    case "keyword":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.keyword') = ?", param), value)
    case "kind":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.kind') = ?", param), value)
    case "path":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.path') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "valueset":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.valueset') = ?", param), value)
    case "contact":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.contact') = ?", param), value)
    case "criteria":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.criteria') = ?", param), value)
    case "payload":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.payload') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.type') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "container-identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.container-identifier') = ?", param), value)
    case "expiry":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.expiry') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.quantity') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "substance-reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.substance-reference') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "receiver":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.receiver') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "supplier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.supplier') = ?", param), value)
    case "category":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.category') = ?", param), value)
    case "requester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requester') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "supplier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.supplier') = ?", param), value)
    case "authored-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.authored-on') = ?", param), value)
    case "based-on":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.based-on') = ?", param), value)
    case "business-status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.business-status') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "encounter":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.encounter') = ?", param), value)
    case "focus":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.focus') = ?", param), value)
    case "group-identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.group-identifier') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "intent":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.intent') = ?", param), value)
    case "modified":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.modified') = ?", param), value)
    case "owner":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.owner') = ?", param), value)
    case "part-of":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.part-of') = ?", param), value)
    case "patient":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.patient') = ?", param), value)
    case "performer":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.performer') = ?", param), value)
    case "period":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.period') = ?", param), value)
    case "priority":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.priority') = ?", param), value)
    case "requester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.requester') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "subject":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.subject') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "issued":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.issued') = ?", param), value)
    case "participant":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.participant') = ?", param), value)
    case "result":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.result') = ?", param), value)
    case "tester":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.tester') = ?", param), value)
    case "testscript":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.testscript') = ?", param), value)
    case "context":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context') = ?", param), value)
    case "context-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-quantity') = ?", param), value)
    case "context-type":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type') = ?", param), value)
    case "date":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.date') = ?", param), value)
    case "description":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.description') = ?", param), value)
    case "identifier":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.identifier') = ?", param), value)
    case "jurisdiction":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.jurisdiction') = ?", param), value)
    case "name":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.name') = ?", param), value)
    case "publisher":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.publisher') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    case "testscript-capability":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.testscript-capability') = ?", param), value)
    case "title":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.title') = ?", param), value)
    case "url":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.url') = ?", param), value)
    case "version":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.version') = ?", param), value)
    case "context-type-quantity":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-quantity') = ?", param), value)
    case "context-type-value":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.context-type-value') = ?", param), value)
    case "code":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.code') = ?", param), value)
    case "expansion":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.expansion') = ?", param), value)
    case "reference":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.reference') = ?", param), value)
    case "target":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.target') = ?", param), value)
    case "datewritten":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.datewritten') = ?", param), value)
    case "prescriber":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.prescriber') = ?", param), value)
    case "status":
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.status') = ?", param), value)
    default:
      query.AndWhere(fmt.Sprintf("json_extract(resource, '$.%s') LIKE ?", param), fmt.Sprintf("%%%s%%", value))
    }
  }

  records, err := query.All()
  if err != nil {
    return nil, fmt.Errorf("failed to query resources: %w", err)
  }

  return records, nil
}
