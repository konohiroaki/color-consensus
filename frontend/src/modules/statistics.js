import axios from "axios";

export const types = {
    SET_VOTES: "SET_VOTES",
    SET_NATIONALITY_FILTER: "SET_NATIONALITY_FILTER",
    CALCULATE_BORDER: "CALCULATE_BORDER",
};

const DEFAULT_STATE = {
    votes: [],
    nationalityFilter: "",
    genderFilter: "",
    ageGroupFilter: "",

    cellBorder: [],
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_VOTES:
            return {
                ...state,
                votes: action.payload
            };
        case types.CALCULATE_BORDER:
            return {
                ...state,
                cellBorder: action.payload
            };
        case types.SET_NATIONALITY_FILTER:
            return {
                ...state,
                nationalityFilter: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    setVotes(color) {
        return (dispatch) => {
            return axios.get(getStatisticsUrl(color))
                .then(({data}) => {
                    dispatch({type: types.SET_VOTES, payload: data});
                    dispatch(this.calculateBorder());
                })
                .catch(err => {});
        };
    },
    calculateBorder() {
        return (dispatch, getState) => {
            const cellRatio = getCellRatio(getState);
            const cellBorder = getCellBorder(getState, cellRatio);
            dispatch({type: types.CALCULATE_BORDER, payload: cellBorder});
        };
    },
    setNationalityFilter(nationality) {
        return (dispatch) => {
            dispatch({type: types.SET_NATIONALITY_FILTER, payload: nationality});
            dispatch(this.calculateBorder());
        };
    },
};

const getStatisticsUrl = ({lang, name}) => {
    const fields = ["colors", "voter.nationality", "voter.gender", "voter.birth"];
    return `${process.env.WEBAPI_HOST}/api/v1/votes?lang=${lang}&name=${name}&fields=${fields}`;
};

const getCellRatio = (getState) => {
    const colorCodeList = getState().board.colorCodeList;
    const votes = getState().statistics.votes;
    const nationalityFilter = getState().statistics.nationalityFilter;
    const arraySize = getState().board.sideLength + 2;

    let ratio = Array(arraySize).fill(0).map(() => Array(arraySize).fill(0));

    const filteredVotes = votes
        .filter(vote => nationalityFilter === "" || nationalityFilter === vote.voter.nationality);
    filteredVotes.flatMap(vote => vote.colors)
        .forEach(color => {
            const idx = colorCodeList.indexOf(color);
            const ii = Math.floor(idx / arraySize) + 1, jj = idx % arraySize + 1;
            ratio[ii][jj] += 1 / filteredVotes.length;
        });
    return ratio;
};

const getCellBorder = (getState, cellRatio) => {
    const arraySize = getState().board.sideLength + 2;
    let border = Array(arraySize).fill(0)
        .map(() => Array(arraySize).fill({top: false, right: false, bottom: false, left: false}));
    const percentile = (100 - 60) / 100; // TODO: use user input for value subtracting from 100.
    for (let ii = 1; ii < border.length - 1; ii++) {
        for (let jj = 1; jj < border.length - 1; jj++) {
            border[ii][jj] = {
                // TODO: make the condition simpler if possible
                top: cellRatio[ii][jj] !== 0 && cellRatio[ii - 1][jj] > percentile && cellRatio[ii][jj] <= percentile
                     || cellRatio[ii][jj] !== 0 && cellRatio[ii - 1][jj] <= percentile && cellRatio[ii][jj] > percentile,
                right: cellRatio[ii][jj] !== 0 && cellRatio[ii][jj + 1] > percentile && cellRatio[ii][jj] <= percentile
                       || cellRatio[ii][jj] !== 0 && cellRatio[ii][jj + 1] <= percentile && cellRatio[ii][jj] > percentile,
                bottom: cellRatio[ii][jj] !== 0 && cellRatio[ii + 1][jj] > percentile && cellRatio[ii][jj] <= percentile
                        || cellRatio[ii][jj] !== 0 && cellRatio[ii + 1][jj] <= percentile && cellRatio[ii][jj] > percentile,
                left: cellRatio[ii][jj] !== 0 && cellRatio[ii][jj - 1] > percentile && cellRatio[ii][jj] <= percentile
                      || cellRatio[ii][jj] !== 0 && cellRatio[ii][jj - 1] <= percentile && cellRatio[ii][jj] > percentile,
            };
        }
    }
    return border;
};
