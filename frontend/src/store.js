import {createStore, combineReducers, applyMiddleware} from "redux";
import thunk from "redux-thunk";
import {composeWithDevTools} from "redux-devtools-extension";

import {reducer as colors} from "./ducks/colors";
import {reducer as user} from "./ducks/user";

const rootReducer = combineReducers({
    colors,
    user
});

export default createStore(rootReducer, composeWithDevTools(
    applyMiddleware(thunk)
));
