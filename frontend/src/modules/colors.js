import axios from "axios";

const types = {
    FETCH_COLORS_SUCCESS: "FETCH_COLORS_SUCCESS",
    FETCH_COLORS_FAILURE: "FETCH_COLORS_FAILURE",
    SET_DISPLAYED_COLOR: "SET_DISPLAYED_COLOR",
    SET_DISPLAYED_COLOR_LIST: "SET_DISPLAYED_COLOR_LIST",
};

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
        case types.FETCH_COLORS_FAILURE:
            return {
                ...state,
                error: action.payload
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
        return (dispatch, getState) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors/keys`)
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
    setDisplayedColor(color) {
        return (dispatch, getState) => {
            const colors = getState().colors.colors;
            if (colors.length !== 0) {
                // TODO: better to have validation for the input color (should be contained in `colors`)
                const displayedColor = color === undefined ? colors[0] : color;
                dispatch({
                    type: types.SET_DISPLAYED_COLOR,
                    payload: displayedColor
                });

                const baseCode = displayedColor.code.substring(1); // remove "#"
                const size = Math.pow(getState().colors.boardSideLength, 2);
                const url = `${process.env.WEBAPI_HOST}/api/v1/colors/candidates/${baseCode}?size=${size}`;
                axios.get(url).then(({data}) => {
                    dispatch({
                        type: types.SET_DISPLAYED_COLOR_LIST,
                        payload: data
                    });
                });
            }
        };
    },
};
