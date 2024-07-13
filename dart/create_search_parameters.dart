import 'dart:convert';
import 'dart:io';

import 'package:fhir_r5/fhir_r5.dart';

Future<void> main() async {
  final searchMap = <String, Map<String, String>>{};
  final Bundle spBundle = Bundle.fromJsonString(
      await File('search-parameters.json').readAsString());
  for (final BundleEntry entry in spBundle.entry ?? <BundleEntry>[]) {
    final SearchParameter sp = entry.resource as SearchParameter;
    final name = sp.name;
    final code = sp.code?.toString();
    final type = sp.type?.toString();
    if (name != null && code != null && type != null) {
      for (final base in sp.base ?? <FhirCode>[]) {
        if (searchMap[base.toString()] == null) {
          searchMap[base.toString()] = <String, String>{};
        }
        searchMap[base.toString()]![name] = type;
      }
    }
  }
  for (final key in searchMap.keys) {
    searchMap[key] = {}
      ..addEntries(searchMap['Resource']!.entries)
      ..addEntries(searchMap['DomainResource']!.entries)
      ..addEntries(searchMap[key]!.entries);
  }

  searchMap.remove('Resource');
  searchMap.remove('DomainResource');

  String goString = '''// collections.go
package main

var collections = []map[string]interface{}{''';

  for (final key in searchMap.keys) {
    goString += mainClass(key);
    final fieldMap = fhirFieldMap[key];
    searchMap[key]!.forEach((k, v) {
      final entry = fieldMap?[k];
      if (entry?.isList ?? false) {
        if (v != 'token' && v != 'reference') {
          print('$key $k $v');
        }
      }
      goString += entries(k, v);
    });
    goString += '''	},
	  },''';
    goString += historyClass(key);
    goString += tokenClass(key); // Add token table for each resource
  }

  goString += '}';

  await File('search_parameters_map.go').writeAsString(goString);
}

String entries(String key, String value) {
  switch (value) {
    case 'number':
      return '			{"name": "$key", "type": "number"},\n';
    case 'date':
      return '			{"name": "$key", "type": "text"},\n';
    case 'string':
      return '			{"name": "$key", "type": "text"},\n';
    case 'token':
      return ''; // Token is handled separately
    case 'reference':
      return '			{"name": "$key", "type": "text"},\n';
    case 'uri':
      return '			{"name": "$key", "type": "text"},\n';
    default:
      return '';
  }
}

String mainClass(String className) => '''\n
	{
		"name": "$className",
		"schema": []map[string]interface{}{
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
''';

String historyClass(String className) => '''\n
	{
		"name": "${className}History",
		"schema": []map[string]interface{}{
			{"name": "fhirId", "type": "text"},
			{"name": "versionId", "type": "number"},
			{"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
		},
	},''';

String tokenClass(String className) => '''\n
	{
		"name": "${className}Token",
		"schema": []map[string]interface{}{
			{"name": "resourceId", "type": "text"},
			{"name": "system", "type": "text"},
			{"name": "code", "type": "text"},
			{"name": "searchParameter", "type": "text"},
			{"name": "index", "type": "number"},
		},
	},''';

JsonEncoder encoder = JsonEncoder.withIndent('  ');

String prettyPrint(Map<String, Map<String, String>> searchMap) {
  return encoder.convert(searchMap);
}
