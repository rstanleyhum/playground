import 'package:flutter/material.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux/redux.dart';

import '../../actions/counter_actions.dart';
import '../../models/app_state.dart';

class IncreaseCountButton extends StatelessWidget {
    IncreaseCountButton({Key key}) : super(key: key);

    @override
    Widget build(BuildContext context) =>
        StoreConnector<AppState, VoidCallback>(
            converter: (Store<AppState> store) {
                return () {
                  store.dispatch(IncrementCountAction());
                };
            },
            builder: (BuildContext context, VoidCallback increase) =>
                FloatingActionButton(
                    onPressed: increase,
                    child: Icon(Icons.add),
                ),
        );
}
