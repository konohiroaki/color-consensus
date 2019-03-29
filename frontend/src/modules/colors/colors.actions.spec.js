import {actions, types} from "./colors";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";

describe("fetchColors()", function () {
    it("should dispatch when fetch is success", () => {
        const fakeResponse = [{foo: "bar"}];
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors/keys`)
            .reply(200, fakeResponse);

        const dispatch = jest.fn();
        actions.fetchColors()(dispatch).then(() => {
            expect(dispatch.mock.calls[0][0]).toEqual({
                type: types.FETCH_COLORS_SUCCESS,
                payload: fakeResponse,
            });
        });
    });
    it("should not dispatch when error response", () => {
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors/keys`)
            .reply(400);

        const dispatch = jest.fn();
        actions.fetchColors()(dispatch).then(() => {
            expect(dispatch.mock.calls.length).toEqual(0);
        });
    });
});

describe("setDisplayedColor(color)", function () {
    const getState = () => ({colors: {colors: [{code: "foo"}], boardSideLength: 3}});
    it("should dispatch when colors are present", () => {
        const dispatch = jest.fn();
        actions.setDisplayedColor()(dispatch, getState);
        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.SET_DISPLAYED_COLOR,
            payload: {code: "foo"},
        });
    });
    it("should not dispatch when colors are absent", () => {
        const dispatch = jest.fn();
        actions.fetchColors()(dispatch, getState);
        expect(dispatch.mock.calls.length).toEqual(0);
    });
    it("should dispatch for displayed color list when get succeeds", () => {
        const fakeData = ["#ff0000", "#f00000"];
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors/candidates/oo?size=9`)
            .reply(200, fakeData);

        const dispatch = jest.fn();
        actions.setDisplayedColor()(dispatch, getState).then(() => {
            expect(dispatch.mock.calls[1][0]).toEqual({
                type: types.SET_DISPLAYED_COLOR_LIST,
                payload: fakeData,
            });
        });
    });
    it("should not dispatch for displayed color list when get fails", () => {
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors/candidates/oo?size=9`)
            .reply(400);

        const dispatch = jest.fn();
        actions.setDisplayedColor()(dispatch, getState).then(() => {
            // toEqual(1) because prior dispatch is executed with this getState.
            expect(dispatch.mock.calls.length).toEqual(1);
        });
    });
});
