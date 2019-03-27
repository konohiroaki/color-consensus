import axios from "axios";

export const types = {
    FETCH_COLORS_SUCCESS: "FETCH_COLORS_SUCCESS",
    SET_DISPLAYED_COLOR: "SET_DISPLAYED_COLOR",
    SET_DISPLAYED_COLOR_LIST: "SET_DISPLAYED_COLOR_LIST",
};

// TODO: improve naming.
const DEFAULT_STATE = {
    colors: [],
    error: null,
    displayedColor: null,
    displayedColorList: [],
    boardSideLength: 31,
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.FETCH_COLORS_SUCCESS:
            return {
                ...state,
                colors: action.payload
            };
        case types.SET_DISPLAYED_COLOR:
            return {
                ...state,
                displayedColor: action.payload
            };
        case types.SET_DISPLAYED_COLOR_LIST:
            return {
                ...state,
                displayedColorList: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    fetchColors() {
        return (dispatch) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors/keys`)
                .then(({data}) => dispatch({type: types.FETCH_COLORS_SUCCESS, payload: data}))
                .catch(err => {});
        };
    },
    setDisplayedColor(color) {
        return (dispatch, getState) => {
            const colors = getState().colors.colors;
            if (colors.length !== 0) {
                // TODO: better to have validation for the input color (should be contained in `colors`)
                const displayedColor = color === undefined ? colors[0] : color;
                dispatch({type: types.SET_DISPLAYED_COLOR, payload: displayedColor});

                const baseCode = displayedColor.code.substring(1); // remove "#"
                const size = Math.pow(getState().colors.boardSideLength, 2);
                return axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors/candidates/${baseCode}?size=${size}`)
                    .then(({data}) => dispatch({type: types.SET_DISPLAYED_COLOR_LIST, payload: data}))
                    .catch(err => {});
            }
        };
    },
};
