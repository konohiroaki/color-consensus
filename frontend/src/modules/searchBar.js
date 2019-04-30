import axios from "axios";
import {toast} from "react-toastify";

export const types = {
    FETCH_COLORS_SUCCESS: "FETCH_COLORS_SUCCESS",
};

const DEFAULT_STATE = {
    colors: [],
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.FETCH_COLORS_SUCCESS:
            return {
                ...state,
                colors: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    fetchColors() {
        return (dispatch) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors`)
                .then(({data}) => dispatch({type: types.FETCH_COLORS_SUCCESS, payload: data}))
                .catch(({response}) => toast.warn(response.data.error.message));
        };
    },
};
