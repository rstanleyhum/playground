import '../actions/counter_actions.dart';

int counterReducer(int state, action) {
    if (action is IncrementCountAction) {
        return state + 1;
    }

    if (action is DecrementCountAction) {
        return state - 1;
    }

    return state;
}


