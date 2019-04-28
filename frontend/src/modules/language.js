import axios from "axios";

export const types = {
    SET_LANGUAGES: "SET_LANGUAGES",
};

const DEFAULT_STATE = {
    languages: []
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_LANGUAGES:
            return {
                ...state,
                languages: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    setLanguages() {
        return (dispatch) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/languages`)
                .then(({data}) => dispatch({type: types.SET_LANGUAGES, payload: data}))
                .catch(err => {});
        };
    },
};
