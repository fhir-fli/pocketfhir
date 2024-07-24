import 'dart:convert';
import 'dart:io';

import 'package:fhir_r5/fhir_r5.dart';

Future<void> main() async {
  final searchMap = <String, List<Map<String, String>>>{};
  final Bundle spBundle = Bundle.fromJsonString(
      await File('search-parameters.json').readAsString());

  for (final BundleEntry entry in spBundle.entry ?? <BundleEntry>[]) {
    final SearchParameter sp = entry.resource as SearchParameter;
    final name = sp.name;
    final code = sp.code?.toString();
    final type = sp.type?.toString();
    final expression = sp.expression;

    if (name != null && code != null && type != null && expression != null) {
      for (final base in sp.base ?? <FhirCode>[]) {
        if (searchMap[base.toString()] == null) {
          searchMap[base.toString()] = <Map<String, String>>[];
        }
        if (expression.contains('|')) {
          searchMap[base.toString()]!
              .add({'name': name, 'type': type, 'expression': ''});
        } else {
          searchMap[base.toString()]!
              .add({'name': name, 'type': type, 'expression': expression});
        }
      }
    }
  }

  for (final key in searchMap.keys) {
    final List<Map<String, String>> resourceEntries =
        searchMap['Resource']!.map((e) {
      if (e['expression']!.startsWith('Resource.')) {
        return {
          'name': e['name']!,
          'type': e['type']!,
          'expression': e['expression']!.replaceFirst('Resource.', '$key.')
        };
      }
      return e;
    }).toList();
    searchMap[key] = resourceEntries
      // ..addAll(searchMap['DomainResource'] ?? [])
      ..addAll(searchMap[key]!);
  }

  searchMap.remove('Resource');
  searchMap.remove('DomainResource');

  // Change the output file name to collections.go
  String goString = '''// collections.go
package main

var collections = []map[string]interface{}{''';

  for (final key in searchMap.keys) {
    goString += mainClass(key);
    searchMap[key]!.forEach((entry) {
      final type = entry['type']!;
      final name = entry['name']!;
      if (name != '_id' && name != '_content') {
        goString += entries(name, type);
      }
    });
    goString += '''    },
  },''';
    goString += historyClass(key, searchMap[key]!);
  }

  goString += '\n}';

  // Write to collections.go instead of search_parameters_map.go
  await File('collections.go').writeAsString(goString);

  // Generate search_parameters.go file
  await generateSearchParametersFile(searchMap);
}

String entries(String key, String value) {
  switch (value) {
    case 'number':
      return '      {"name": "$key", "type": "REAL"},\n';
    case 'date':
      return '      {"name": "$key", "type": "TEXT"},\n';
    case 'string':
      return '      {"name": "$key", "type": "TEXT"},\n';
    case 'token':
      return '      {"name": "$key", "type": "JSON"},\n';
    case 'reference':
      return '      {"name": "$key", "type": "JSON"},\n';
    case 'composite':
      return '      {"name": "$key", "type": "TEXT"},\n';
    case 'quantity':
      return '      {"name": "$key", "type": "JSON"},\n';
    case 'uri':
      return '      {"name": "$key", "type": "TEXT"},\n';
    case 'special':
      return '      {"name": "$key", "type": "TEXT"},\n';
    default:
      return '';
  }
}

String mainClass(String className) => '''\n
  {
    "name": "$className",
    "schema": []map[string]interface{}{
      {"name": "id", "type": "TEXT"},
      {"name": "versionId", "type": "REAL"},
      {"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
''';

String historyClass(String className, List<Map<String, String>> columns) {
  String historySchema = '''\n
  {
    "name": "${className}History",
    "schema": []map[string]interface{}{
      {"name": "fhirId", "type": "TEXT"},
      {"name": "versionId", "type": "REAL"},
      {"name": "resource", "type": "json", "options": map[string]interface{}{"maxSize": 5242880}},
''';

  columns.forEach((entry) {
    final name = entry['name']!;
    final type = entry['type']!;
    if (name != '_id' && name != '_content') {
      historySchema += entries(name, type);
    }
  });

  historySchema += '''    },
  },''';

  return historySchema;
}

Future<void> generateSearchParametersFile(
    Map<String, List<Map<String, String>>> searchMap) async {
  String searchParamsGoString = '''// search_parameters.go
package main

type SearchParameter struct {
	Code       string
	Expression string
}

var searchParamsByResourceType = map[string][]SearchParameter{''';

  searchMap.forEach((resourceType, params) {
    searchParamsGoString += '''
    "$resourceType": {''';
    params.forEach((entry) {
      final name = entry['name']!;
      final expression = entry['expression']!;
      searchParamsGoString += '''
      {
        Code: "$name",
        Expression: "$expression",
      },''';
    });
    searchParamsGoString += '''
    },''';
  });

  searchParamsGoString += '''
}
''';

  // Write to search_parameters.go
  await File('search_parameters.go').writeAsString(searchParamsGoString);
}

JsonEncoder encoder = JsonEncoder.withIndent('  ');

String prettyPrint(Map<String, Map<String, String>> searchMap) {
  return encoder.convert(searchMap);
}
