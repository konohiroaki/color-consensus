import axios from "axios";

export const types = {
    SET_BASE_COLOR: "SET_BASE_COLOR",
    SET_COLOR_CODE_LIST: "SET_COLOR_CODE_LIST",
};

const DEFAULT_STATE = {
    cellSize: "25px",
    sideLength: 19,

    baseColor: null,
    colorCodeList: [],
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_BASE_COLOR:
            return {
                ...state,
                baseColor: action.payload
            };
        case types.SET_COLOR_CODE_LIST:
            return {
                ...state,
                colorCodeList: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    setBaseColor(color) {
        return (dispatch, getState) => {
            const colors = getState().colors.colors;
            if (colors.length !== 0) {
                const baseColor = getBaseColor(colors, color);
                dispatch({type: types.SET_BASE_COLOR, payload: baseColor});

                return axios.get(getColorListUrl(baseColor, getState().board.sideLength))
                    .then(({data}) => dispatch({type: types.SET_COLOR_CODE_LIST, payload: data}))
                    .catch(err => {});
            }
        };
    },
};

const getBaseColor = (colors, color) => color !== undefined && colors.includes(color) ? color : colors[0];

const getColorListUrl = (baseColor, sideLength) => {
    const baseColorCode = baseColor.code.substring(1); // remove "#"
    const size = Math.pow(sideLength, 2);
    return `${process.env.WEBAPI_HOST}/api/v1/colors/${baseColorCode}/neighbors?size=${size}`;
};
