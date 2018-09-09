import 'package:flutter/material.dart';
import 'package:firebase_storage/firebase_storage.dart';

import 'rootpage.dart';
import 'auth.dart';
import 'analytics.dart';


void main() async {
  runApp(new MyApp(storage: FirebaseStorage.instance));
}

class MyApp extends StatelessWidget {
  MyApp({this.storage});

  final FirebaseStorage storage;

    @override
    Widget build(BuildContext context) => MaterialApp(
        title: 'flutter login demo',
        theme: ThemeData(
            primarySwatch: Colors.blue,
        ),
        home: RootPage(auth: Auth(), analytics: FbAnalytics(), storage: storage),
    );
}

