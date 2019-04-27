import {actions, types} from "../../statistics";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";

// TODO: test other action called in dispatch.
describe("setVotes(color)", function () {
    const fakeResponse = [{
        colors: ["#ff0000"],
        voter: {nationality: "Japan", ageGroup: 20, gender: "Male"}
    }];
    const fakeArgument = {lang: "en", name: "red", code: "#ff0000"};
    const url = process.env.WEBAPI_HOST + "/api/v1/votes?lang=en&name=red"
                + "&fields=colors,voter.nationality,voter.ageGroup,voter.gender";

    it("should dispatch SET_VOTES when fetch is success", () => {
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(url).reply(200, fakeResponse);

        const dispatch = jest.fn();
        actions.setVotes(fakeArgument)(dispatch).then(() => {
            expect(dispatch.mock.calls[0][0]).toEqual({
                type: types.SET_VOTES,
                payload: fakeResponse,
            });
        });
    });
});

describe("calculateBorder()", function () {
    const filterEmptyState = () => ({
        board: {
            sideLength: 2,
            colorCodeList: ["#ff0000", "#f00000", "#f01000", "#f00010"],
        },
        statistics: {
            votes: [
                {colors: ["#ff0000"], voter: {nationality: "Japan", ageGroup: 20, gender: "Male"}},
                {colors: ["#ff0000"], voter: {nationality: "China", ageGroup: 30, gender: "Female"}},
                {colors: ["#f00000"], voter: {nationality: "Japan", ageGroup: 40, gender: "Male"}},
                {colors: ["#f01000"], voter: {nationality: "Korea", ageGroup: 50, gender: "Female"}},
            ],
            nationalityFilter: "",
            ageGroupFilter: "",
            genderFilter: "",
            percentile: 0
        }
    });

    const allFalsePayload = () => Array(4).fill(0).map(() => Array(4).fill(
        {top: false, right: false, bottom: false, left: false}
    ));

    it("should dispatch with empty filter", () => {
        const dispatch = jest.fn();

        actions.calculateBorder()(dispatch, filterEmptyState);

        let expectedPayload = allFalsePayload();
        expectedPayload[1][1] = {top: true, right: false, bottom: false, left: true};
        expectedPayload[1][2] = {top: true, right: true, bottom: true, left: false};
        expectedPayload[2][1] = {top: false, right: true, bottom: true, left: true};
        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.CALCULATE_BORDER,
            payload: expectedPayload
        });
    });

    it("should dispatch with nationality filter", () => {
        const dispatch = jest.fn();

        let state = filterEmptyState();
        state.statistics.nationalityFilter = "Japan";
        actions.calculateBorder()(dispatch, () => state);

        let expectedPayload = allFalsePayload();
        expectedPayload[1][1] = {top: true, right: false, bottom: true, left: true};
        expectedPayload[1][2] = {top: true, right: true, bottom: true, left: false};
        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.CALCULATE_BORDER,
            payload: expectedPayload
        });
    });

    it("should dispatch with ageGroup filter", () => {
        const dispatch = jest.fn();

        let state = filterEmptyState();
        state.statistics.ageGroupFilter = 50;
        actions.calculateBorder()(dispatch, () => state);

        let expectedPayload = allFalsePayload();
        expectedPayload[2][1] = {top: true, right: true, bottom: true, left: true};
        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.CALCULATE_BORDER,
            payload: expectedPayload
        });
    });

    it("should dispatch with gender filter", () => {
        const dispatch = jest.fn();

        let state = filterEmptyState();
        state.statistics.genderFilter = "Male";
        actions.calculateBorder()(dispatch, () => state);

        let expectedPayload = allFalsePayload();
        expectedPayload[1][1] = {top: true, right: false, bottom: true, left: true};
        expectedPayload[1][2] = {top: true, right: true, bottom: true, left: false};
        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.CALCULATE_BORDER,
            payload: expectedPayload
        });
    });

    it("should dispatch with percentile filter", () => {
        const dispatch = jest.fn();

        let state = filterEmptyState();
        state.statistics.percentile = 50;
        actions.calculateBorder()(dispatch, () => state);

        let expectedPayload = allFalsePayload();
        expectedPayload[1][1] = {top: true, right: true, bottom: true, left: true};
        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.CALCULATE_BORDER,
            payload: expectedPayload
        });
    });
});

describe("setNationalityFilter(nationality)", function () {
    it("should dispatch", () => {
        const fakeData = "Japan";
        const dispatch = jest.fn();

        actions.setNationalityFilter(fakeData)(dispatch);

        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.SET_NATIONALITY_FILTER,
            payload: fakeData
        });
    });
});

describe("setAgeGroupFilter(ageGroup)", function () {
    it("should dispatch as int 20", () => {
        const fakeData = "20";
        const dispatch = jest.fn();

        actions.setAgeGroupFilter(fakeData)(dispatch);

        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.SET_AGE_GROUP_FILTER,
            payload: 20
        });
    });
    it("should dispatch as \"\"", () => {
        const fakeData = "";
        const dispatch = jest.fn();

        actions.setAgeGroupFilter(fakeData)(dispatch);

        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.SET_AGE_GROUP_FILTER,
            payload: fakeData
        });
    });
});

describe("setGenderFilter(gender)", function () {
    it("should dispatch", () => {
        const fakeData = "Male";
        const dispatch = jest.fn();

        actions.setGenderFilter(fakeData)(dispatch);

        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.SET_GENDER_FILTER,
            payload: fakeData
        });
    });
});

describe("setPercentile(percentile)", function () {
    it("should dispatch", () => {
        const fakeData = 10;
        const dispatch = jest.fn();

        actions.setPercentile(fakeData)(dispatch);

        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.SET_PERCENTILE,
            payload: fakeData
        });
    });
});

describe("resetFilters()", function () {
    it("should dispatch", () => {
        const dispatch = jest.fn();

        actions.resetFilters()(dispatch);

        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.RESET_FILTERS,
        });
    });
});
