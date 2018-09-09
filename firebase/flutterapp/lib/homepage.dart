import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';
import 'package:firebase_storage/firebase_storage.dart';

import 'auth.dart';
import 'analytics.dart';


class HomePage extends StatelessWidget {
  HomePage({this.auth, this.onSignedOut, this.analytics, this.storage});
  final BaseAuth auth;
  final VoidCallback onSignedOut;
  final BaseAnalytics analytics;
  final FirebaseStorage storage;

  void _getFiles() async {
    var url = "https://us-central1-logindemo-5f7b4.cloudfunctions.net/assets";
    http.get(url).then( (response) {
      print("Response status: ${response.statusCode}");
      List<dynamic> parsed = jsonDecode(response.body);
      parsed.forEach((f) {
        print(f.toString());
      });
    })
    .catchError((e) {
      print(e);
    });
  }

  void _displayFiles() async {
    print('start display');
    print(await storage.ref().child('/assets/').getPath());
    print(await storage.ref().child('/assets/').getBucket());
    print(await storage.ref().child('/assets/').getName());
    this.storage.ref().child('/assets/pediatrics/chony_handbook/cardiology_ekg.md').getMetadata().then( (r) {
       print(r.name);
       print(r.path);
       print(r.bucket);
       print(r.contentType);
    }).catchError((e) {
      print(e);
    });
  }
  void _signOut() async {
    try {
      await auth.signOut();
      onSignedOut();
    } catch (e) {
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) => Scaffold(
        appBar: AppBar(title: Text('Welcome'), actions: <Widget>[
          FlatButton(
            child: Text('Logout',
                style: TextStyle(fontSize: 17.0, color: Colors.white)),
            onPressed: _signOut,
          ),
        ]),
        body: Container(
          child: Column(
            children: <Widget>[
              MaterialButton(
                child: Text('Test logEvent'),
                onPressed: analytics.sendAnalyticsEvent,
              ),
              MaterialButton(
                child: Text('Test standard event types'),
                onPressed: analytics.testAllEventTypes,
              ),
              MaterialButton(
                child: Text('Test setUserId'),
                onPressed: analytics.testSetUserId,
              ),
              MaterialButton(
                child: Text('Test setCurrentScreen'),
                onPressed: analytics.testSetCurrentScreen,
              ),
              MaterialButton(
                child: Text('Test setAnalyticsCollectionEnabled'),
                onPressed: analytics.testSetAnalyticsCollectionEnabled,
              ),
              MaterialButton(
                child: Text('Test setMinimumsessionDuration'),
                onPressed: analytics.testSetMinimumSessionDuration,
              ),
              MaterialButton(
                child: Text('Test setSessionTimeoutDuration'),
                onPressed: analytics.testSetSessionTimeoutDuration,
              ),
              MaterialButton(
                child: Text('Test setUserProperty'),
                onPressed: analytics.testSetUserProperty,
              ),
              MaterialButton(
                child: Text('Show Files'),
                onPressed: _displayFiles,
              ),
              MaterialButton(
                child: Text('Get Files'),
                onPressed: _getFiles,
              ),
            ],
          ),
        ),
      );
}
