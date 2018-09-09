import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux/redux.dart';

import '../../models/app_state.dart';


class Counter extends StatelessWidget {
    @override
    Widget build(BuildContext context) =>
        StoreConnector<AppState, _ViewModel>(
            converter: _ViewModel.fromStore,
            builder: (BuildContext context, _ViewModel vm) =>
                Text(
                    vm.count.toString(),
                    style: Theme.of(context).textTheme.display1,
                ),
        );
}


class _ViewModel {
    final int count;

    _ViewModel({
        @required this.count,
    });

    static _ViewModel fromStore(Store<AppState> store) =>
        _ViewModel(
            count: store.state.count,
        );
}
