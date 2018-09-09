import 'package:flutter/material.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux/redux.dart';
import 'package:redux_logging/redux_logging.dart';

import 'middleware/auth_middleware.dart';
import 'models/app_state.dart';
import 'pages/home_page.dart';
import 'reducers/app_reducer.dart';

void main() => runApp(MainApp());

class MainApp extends StatelessWidget {
  final String title = 'MeSuite';

  final store = Store<AppState>(
    appReducer,
    initialState: AppState(),
    middleware: []
        ..addAll(createAuthMiddleware())
        ..add(LoggingMiddleware.printer()),
  );

  @override
  Widget build(BuildContext context) =>
    StoreProvider(
        store:store,
        child: MaterialApp(
                title: title,
                home: HomePage(title),
        ),
    );
}

