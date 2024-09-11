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
    createOrUpdateRecord(pb, patient);

    // final String patientFile = File('assets/Patient.ndjson').readAsStringSync();
    // final List<String> patientStrings = patientFile.split('\n');
    // patientStrings.removeWhere((e) => e == '');
    // final List<dynamic> patients =
    //     patientStrings.map((e) => jsonDecode(e)).toList();
    // for (final dynamic patient in patients) {
    //   if (patient is Map<String, dynamic>) {
    //     createOrUpdateRecord(pb, patient);
    //   }
    // }
  } catch (e) {
    print(e);
  }
}

// Function to create or update a record
Future<Map<String, dynamic>> createOrUpdateRecord(
    PocketBase pb, Map<String, dynamic> resourceMap) async {
  final Map<String, dynamic> resource = resourceMapToBody(resourceMap);
  try {
    final RecordModel recordModel = await pb
        .collection(
            resource['resource']['resourceType'].toString().toLowerCase())
        .create(body: resource);
    print(recordModel);
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

Map<String, dynamic> patient = {
  'resourceType': 'Patient',
  'birthDate': '1981-09-18'
};

Map<String, dynamic> resourceMapToBody(Map<String, dynamic> resourceMap) => {
      "resourceType": resourceMap['resourceType'],
      "versionId": 1,
      "resource": resourceMap,
    };

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
