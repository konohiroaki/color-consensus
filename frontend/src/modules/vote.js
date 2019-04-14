export const types = {
    SET_SELECTED_COLOR_CODE_LIST: "SET_SELECTED_COLOR_CODE_LIST",
};

const DEFAULT_STATE = {
    selectedColorCodeList: [],
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_SELECTED_COLOR_CODE_LIST:
            return {
                ...state,
                selectedColorCodeList: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    setSelectedColorCodeList: (colorList) => {
        return {
            type: types.SET_SELECTED_COLOR_CODE_LIST,
            payload: colorList,
        };
    },
    resetSelectedColorCodeList: () => {
        return {
            type: types.SET_SELECTED_COLOR_CODE_LIST,
            payload: [],
        };
    }
};
