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
    await pb.admins.authWithPassword('grey@fhirfli.dev', '01 password');
    createOrUpdateRecord(pb, account);
  } catch (e) {
    print(e);
  }
}

// Function to create or update a record
Future<Map<String, dynamic>> createOrUpdateRecord(
    PocketBase pb, Map<String, dynamic> resourceMap) async {
  try {
    final RecordService recordService =
        pb.collection(resourceMap['resourceType'].toString().toLowerCase());
    final RecordModel recordModel = await (resourceMap['id'] != null
        ? recordService.update(resourceMap['id'].toString(),
            body: resourceMapToBody(resourceMap))
        : recordService.create(body: resourceMapToBody(resourceMap)));
    if (recordModel.data['resource'] != null) {
      print('resource returned');
      return recordModel.data['resource'];
    } else {
      return _operationOutcomes('Error creating record');
    }
  } catch (e) {
    print('Error creating or updating record: $e');
    if (e is ClientException && e.statusCode == 404) {
      // final RecordService recordService =
      //     pb.collection(resourceMap['resourceType'].toString().toLowerCase());
      // final RecordModel recordModel =
      //     await recordService.create(body: resourceMapToBody(resourceMap));
      // if (recordModel.data['resource'] != null) {
      //   print('resource returned');
      //   return recordModel.data['resource'];
      // } else {
      return _operationOutcomes(e.toString());
      // }
    } else {
      return _operationOutcomes(e.toString());
    }
  }
}

Map<String, dynamic> account = {
  "resourceType": "Account",
  "id": "1rkf6gc5me9np9r",
  "status": "active",
  "name": "Patient billing account"
};

Map<String, dynamic> resourceMapToBody(Map<String, dynamic> resourceMap) {
  final lastUpdated = DateTime.now().toUtc().toIso8601String();
  // resourceMap['meta'] = {
  //   "versionId": lastUpdated.replaceAll(':', '.'),
  //   "lastUpdated": lastUpdated,
  // };
  return {
    "lastUpdated": lastUpdated,
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
