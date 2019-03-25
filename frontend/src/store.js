import {createStore, combineReducers, applyMiddleware} from "redux";
import {reducer as colors} from "./ducks/colors";
import thunk from "redux-thunk";
import {composeWithDevTools} from "redux-devtools-extension";

const rootReducer = combineReducers({
    colors,
});

export default createStore(rootReducer, composeWithDevTools(
    applyMiddleware(thunk)
));
