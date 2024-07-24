import 'package:fhir_r5/fhir_r5.dart';
import 'package:pocketbase/pocketbase.dart';
import 'package:riverpod/riverpod.dart';

class PbRepository {
  PbRepository(String baseUrl) : pb = PocketBase(baseUrl);
  final PocketBase pb;

  // Authenticate
  Future<void> authWithPassword(String email, String password) =>
      pb.admins.authWithPassword('grey@fhirfli.dev', '01 password');

  Future<Resource> createOrUpdateFhirResource(
          PocketBase pb, Resource resource) async =>
      Resource.fromJson(await createOrUpdateRecord(pb, resource.toJson()));

// Function to create or update a record
  Future<Map<String, dynamic>> createOrUpdateRecord(
      PocketBase pb, Map<String, dynamic> resource) async {
    try {
      if (resource['id'] == null) {
        final RecordModel recordModel = await pb
            .collection(
                resource['resource']['resourceType'].toString().toLowerCase())
            .create(body: resource);
        if (recordModel.data['resource'] != null) {
          return recordModel.data['resource'];
        } else {
          return _operationOutcomes('Error creating record');
        }
      } else {
        final RecordService recordService = pb.collection(
            resource['resource']['resourceType'].toString().toLowerCase());
        final RecordModel recordModel = await recordService.update(
          resource['id'] as String,
          body: resource,
        );
        if (recordModel.data['resource'] != null) {
          return recordModel.data['resource'];
        } else {
          return _operationOutcomes('Error updating record');
        }
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
}

final pbRepositoryProvider = Provider<PbRepository>((ref) {
  return PbRepository('http://127.0.0.1:8090');
});
