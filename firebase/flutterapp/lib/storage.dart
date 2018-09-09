import 'dart:io';

import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_storage/firebase_storage.dart';

class FBStorage {
  FirebaseApp app;
  FirebaseStorage storage;

  FBStorage() {
    FirebaseApp.configure(
      name: 'flutterapp',
      options: FirebaseOptions(
        googleAppID: Platform.isIOS ? 'placeholder' : '1:567254085450:android:16885149ecbb451c',
        gcmSenderID: '567254085450',
        apiKey: 'AIzaSyBtHKFm-y3VsDsNtnvjY7oNWk0XsPbjPMU ',
        projectID: 'logindemo-5f7b4',
      )
    ).then( (app) {
      this.app = app;

      storage = FirebaseStorage(
        app: app,
        storageBucket: 'gs://logindemo-5f7b4.appspot.com',
      );

    }).catchError((e) {
      print(e);
    });
  }
}