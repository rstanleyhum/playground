import 'package:firebase_auth/firebase_auth.dart';

import 'package:flutter/foundation.dart';


class LogIn {}

class LogInSuccessful {
    final FirebaseUser user;

    LogInSuccessful({
        @required this.user,
    });

    @override
    String toString() => 'LogIn{user: $user}';
}


class LogInFail {
    final dynamic error;

    LogInFail(this.error);

    @override
    String toString() => 'LogIn{There was an error logging in: $error}';
}

class LogOut {}


class LogOutSuccessful {
    LogOutSuccessful();

    @override
    String toString() => 'LogOut{user: null}';
}

