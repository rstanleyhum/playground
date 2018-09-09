import 'package:flutter/material.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux/redux.dart';

import '../../actions/auth_actions.dart';
import '../../models/app_state.dart';

import 'google_auth_button.dart';

class GoogleAuthButtonContainer extends StatelessWidget {
    GoogleAuthButtonContainer({Key key}) : super(key: key);

    @override
    Widget build(BuildContext context) =>
        StoreConnector<AppState, _ViewModel>(
            converter: _ViewModel.fromStore,
            builder: (BuildContext context, _ViewModel vm) {
                return GoogleAuthButton(
                    buttonText: vm.buttonText,
                    onPressedCallback: vm.onPressedCallback,
                );
            },
        );
}

class _ViewModel {
    final String buttonText;
    final Function onPressedCallback;

    _ViewModel({
        this.onPressedCallback,
        this.buttonText,
    });

    static _ViewModel fromStore(Store<AppState> store) =>
        _ViewModel(
            buttonText:
                store.state.currentUser != null
                    ? 'Log Out' : 'Log in with Google',
            onPressedCallback: () {
                if (store.state.currentUser != null) {
                    store.dispatch(LogOut());
                } else {
                    store.dispatch(LogIn());
                }
            }
        );
}
