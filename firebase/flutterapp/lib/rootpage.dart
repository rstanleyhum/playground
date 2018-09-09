import 'package:flutter/material.dart';
import 'package:firebase_storage/firebase_storage.dart';

import 'auth.dart';
import 'analytics.dart';
import 'loginpage.dart';
import 'homepage.dart';

class RootPage extends StatefulWidget {
  RootPage({this.auth, this.analytics, this.storage});
  final BaseAuth auth;
  final BaseAnalytics analytics;
  final FirebaseStorage storage;

  @override
  State<StatefulWidget> createState() => _RootPageState();
}

enum AuthStatus { notSignedIn, signedIn }

class _RootPageState extends State<RootPage> {
  AuthStatus authStatus = AuthStatus.notSignedIn;

  @override
  void initState() {
    super.initState();
    widget.auth.currentUser().then((userId) {
      setState(() {
        authStatus = userId == null ? AuthStatus.notSignedIn : AuthStatus.signedIn;
      });
    });
  }

  void _signedIn() {
    setState(() {
      authStatus = AuthStatus.signedIn;
    });
  }

  void _signedOut() {
    setState( () {
      authStatus = AuthStatus.notSignedIn;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (authStatus == AuthStatus.signedIn) {
      return HomePage(
        auth: widget.auth,
        onSignedOut: _signedOut,
        analytics: widget.analytics,
        storage: widget.storage,
      );
    }

    return LoginPage(
      auth: widget.auth,
      onSignedIn: _signedIn,
    );
  }
  
}
