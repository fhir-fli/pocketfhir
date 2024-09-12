import 'dart:convert';
import 'dart:io';

Future<void> main() async {
  final Directory dir = Directory('.');
  final maps = <String, List<Map<String, dynamic>>>{
    'StructureDefinition': <Map<String, dynamic>>[],
    'CapabilityStatement': <Map<String, dynamic>>[],
    'ConceptMap': <Map<String, dynamic>>[],
    'SearchParameter': <Map<String, dynamic>>[],
    'CodeSystem': <Map<String, dynamic>>[],
    'ValueSet': <Map<String, dynamic>>[],
    'OperationDefinition': <Map<String, dynamic>>[],
    'CompartmentDefinition': <Map<String, dynamic>>[],
  };
  for (final file in dir.listSync()) {
    if (file.path.endsWith('.json')) {
      final fileString = jsonDecode(File(file.path).readAsStringSync());
      for (final entry in fileString['entry']) {
        if (entry['resource'] != null) {
          final resource = entry['resource'];
          if (maps.keys.contains(resource['resourceType'])) {
            String id = resource['id'].replaceAll(RegExp(r'[^a-zA-Z0-9]'), '');
            if (id.length > 15) {
              id = id.substring(0, 15);
            }
            resource['id'] = id;
            maps[resource['resourceType']]!.add(resource);
          } else {
            print(resource['resourceType']);
          }
        }
      }
    }
  }
  for (final key in maps.keys) {
    String fileString = '';
    maps[key]!.forEach((e) => fileString += '${jsonEncode(e)}\n');
    await File('${key.toLowerCase()}.ndjson').writeAsString(fileString);
  }
}
