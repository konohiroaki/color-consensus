import axios from "axios";

export const types = {
    FETCH_COLORS_SUCCESS: "FETCH_COLORS_SUCCESS",
    FETCH_COLORS_FAILURE: "FETCH_COLORS_FAILURE",

};

const DEFAULT_STATE = {
    colors: [],
    error: null
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.FETCH_COLORS_SUCCESS:
            return {
                ...state,
                colors: action.payload
            };
        case types.FETCH_COLORS_FAILURE:
            return {
                ...state,
                error: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    fetchColors() {
        return function (dispatch, getState) {
            axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors/keys`)
                .then(({data}) => {
                    dispatch({
                        type: types.FETCH_COLORS_SUCCESS,
                        payload: data,
                    });
                })
                .catch(err => {
                    dispatch({
                        type: types.FETCH_COLORS_FAILURE,
                        payload: err.message,
                    });
                });
        };
    },
};
