import 'package:firebase_auth/firebase_auth.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:redux/redux.dart';

import '../actions/auth_actions.dart';
import '../models/app_state.dart';


List<Middleware<AppState>> createAuthMiddleware() {
    final logIn = _createLogInMiddleware();
    final logOut = _createLogOutMiddleware();

    return [
        TypedMiddleware<AppState, LogIn>(logIn),
        TypedMiddleware<AppState, LogOut>(logOut),
    ];
}

Middleware<AppState> _createLogInMiddleware() {
    return (Store store, action, NextDispatcher next) async {
        FirebaseUser user;

        final FirebaseAuth _auth = FirebaseAuth.instance;

        final GoogleSignIn _googleSignIn = GoogleSignIn();

        if (action is LogIn) {
            try {
                GoogleSignInAccount googleUser = await _googleSignIn.signIn();
                GoogleSignInAuthentication googleAuth = await googleUser.authentication;

                user = await _auth.signInWithGoogle(
                    accessToken: googleAuth.accessToken,
                    idToken: googleAuth.idToken,
                );

                print('Logged in ' + user.displayName);

                store.dispatch(LogInSuccessful(user: user));
            } catch (error) {
                store.dispatch(LogInFail(error));
            }
        }
        next(action);
    };
}


Middleware<AppState> _createLogOutMiddleware() {
    return (Store store, action, NextDispatcher next) async {
        final FirebaseAuth _auth = FirebaseAuth.instance;

        try {
            await _auth.signOut();
            print('logging out...');
            store.dispatch(LogOutSuccessful());
        } catch (error) {
            print(error);
        }
        next(action);
    };
}

