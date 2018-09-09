import '../models/app_state.dart';

import 'counter_reducer.dart';
import 'auth_reducer.dart';


AppState appReducer(AppState state, action) => AppState(
    isLoading: false,
    count: counterReducer(state.count, action),
    currentUser: authReducer(state.currentUser, action),
);
