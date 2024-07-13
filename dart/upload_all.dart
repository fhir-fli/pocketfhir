import 'dart:io';

import 'package:fhir_r5/fhir_r5.dart';
import 'package:pocketbase/pocketbase.dart';

Future<void> main() async {
  final PocketBase pb = PocketBase('http://127.0.0.1:8090');

  // Authenticate
  await pb.admins.authWithPassword('grey@fhirfli.dev', '01 password');

  // await assets(pb);
  // await examples(pb);
  await mimic(pb);
}

final String pocketbaseId = '4260a7372075e64';
//PocketbaseId.generate();

Future<void> assets(PocketBase pb) async {
  final Directory dir1 = Directory('assets');
  final List<FileSystemEntity> files = dir1.listSync();
  for (final FileSystemEntity file in files) {
    if (file is File) {
      print(file.path);
      final Bundle bundle = Bundle.fromJsonString(await file.readAsString());
      for (final BundleEntry entry in bundle.entry ?? <BundleEntry>[]) {
        final resource = entry.resource;
        if (resource != null) {
          await createOrUpdateRecord(pb, resource.toJson());
        }
      }
    }
  }
}

Future<void> examples(PocketBase pb) async {
  final Directory dir2 = Directory('examples');
  final List<FileSystemEntity> files2 = dir2.listSync();
  for (final FileSystemEntity file in files2) {
    if (file is File) {
      print(file.path);
      final Resource resource =
          Resource.fromJsonString(await file.readAsString());
      await createOrUpdateRecord(pb, resource.toJson());
    }
  }
}

Future<void> mimic(PocketBase pb) async {
  final Directory dir3 = Directory('mimic-fhir');
  final List<FileSystemEntity> files3 = dir3.listSync();
  for (final FileSystemEntity file in files3) {
    if (file is File) {
      print(file.path);
      final String fileString = await file.readAsString();
      final List<String> lines = fileString.split('\n');
      for (final String line in lines) {
        try {
          final Resource resource = Resource.fromJsonString(line);
          await createOrUpdateRecord(pb, resource.toJson());
        } catch (e) {
          print(e);
        }
      }
    }
  }
}

// Function to create or update a record
Future<void> createOrUpdateRecord(
    PocketBase pb, Map<String, dynamic> resource) async {
  try {
    await pb
        .collection(resource['resourceType'].toString().toLowerCase())
        .create(body: createBody(resource));
  } catch (e) {
    print(e);
  }
}

Map<String, dynamic> createBody(Map<String, dynamic> resource) => {
      "resourceType": resource['resourceType'],
      "versionId": 1,
      "resource": resource,
    };
