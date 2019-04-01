import axios from "axios";

export const types = {
    SET_STATISTICS: "SET_STATISTICS",
};

const DEFAULT_STATE = {
    voteCount: null,

};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_STATISTICS:
        default:
            return state;
    }
};

export const actions = {
    setStatistics(color) {
        return (dispatch) => {
            // TODO: setup data from GET /votes

        };
    },
};
