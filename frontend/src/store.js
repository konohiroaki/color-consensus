import {createStore, combineReducers, applyMiddleware} from "redux";
import thunk from "redux-thunk";
import {composeWithDevTools} from "redux-devtools-extension";

import {reducer as colors} from "./modules/colors/colors";
import {reducer as user} from "./modules/user/user";
import {reducer as board} from "./modules/board/board";

const rootReducer = combineReducers({
    user,
    colors,
    board
});

export default createStore(rootReducer, composeWithDevTools(
    applyMiddleware(thunk)
));
