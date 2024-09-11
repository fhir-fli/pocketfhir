import 'dart:io';

import 'package:pocketbase/pocketbase.dart';

class HttpOverridesImpl extends HttpOverrides {
  @override
  HttpClient createHttpClient(SecurityContext? context) {
    return super.createHttpClient(context)
      ..badCertificateCallback =
          (X509Certificate cert, String host, int port) => true;
  }
}

Future<void> main() async {
  HttpOverrides.global = HttpOverridesImpl();

  final PocketBase pb = PocketBase('http://127.0.0.1:8090');

  try {
    final adminAuth =
        await pb.admins.authWithPassword('grey@fhirfli.dev', '01 password');
    print(adminAuth.admin?.email);
    createOrUpdateRecord(pb, account);
  } catch (e) {
    print(e);
  }
}

// Function to create or update a record
Future<Map<String, dynamic>> createOrUpdateRecord(
    PocketBase pb, Map<String, dynamic> resourceMap) async {
  final Map<String, dynamic> resource = resourceMapToBody(resourceMap);
  try {
    print('resource: $resource');
    final RecordService recordService =
        pb.collection(resource['resourceType'].toString().toLowerCase());
    print('RecordModel: ${recordService}');
    final RecordModel recordModel = await recordService.create(body: resource);
    if (recordModel.data['resource'] != null) {
      print('resource returned');
      return recordModel.data['resource'];
    } else {
      return _operationOutcomes('Error creating record');
    }
  } catch (e) {
    print('Error creating or updating record: $e');
    if (e is ClientException && e.statusCode == 404) {
      final RecordModel recordModel = await pb
          .collection(
              resource['resource']['resourceType'].toString().toLowerCase())
          .create(body: resource);
      if (recordModel.data['resource'] != null) {
        return recordModel.data['resource'];
      } else {
        return _operationOutcomes(e.toString());
      }
    } else {
      return _operationOutcomes(e.toString());
    }
  }
}

Map<String, dynamic> account = {
  "resourceType": "Account",
  "status": "active",
  "name": "Patient billing account"
};

Map<String, dynamic> resourceMapToBody(Map<String, dynamic> resourceMap) {
  final lastUpdated = DateTime.now().toUtc().toIso8601String();
  resourceMap['meta'] = {
    "versionId": lastUpdated.replaceAll(':', '.'),
    "lastUpdated": lastUpdated,
  };
  return {
    "versionId": lastUpdated,
    "resource": resourceMap,
  };
}

Map<String, dynamic> _operationOutcomes(String error) => <String, dynamic>{
      'resourceType': 'OperationOutcome',
      'issue': <Map<String, dynamic>>[
        <String, dynamic>{
          'severity': 'error',
          'code': 'processing',
          'diagnostics': error,
        }
      ]
    };
