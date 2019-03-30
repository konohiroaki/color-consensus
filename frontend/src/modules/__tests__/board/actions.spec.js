import {actions, types} from "../../board";
import MockAdapter from "axios-mock-adapter";
import axios from "axios";

describe("setBaseColor(color)", function () {
    const colorNonEmptyState = () => ({
        colors: {colors: [{code: "foo"}]},
        board: {sideLength: 3}
    });
    const colorEmptyState = () => ({
        colors: {colors: []},
        board: {sizeLength: 3}
    });

    it("should dispatch when colors are present", () => {
        const dispatch = jest.fn();
        actions.setBaseColor()(dispatch, colorNonEmptyState);
        expect(dispatch.mock.calls[0][0]).toEqual({
            type: types.SET_BASE_COLOR,
            payload: {code: "foo"},
        });
    });
    it("should not dispatch when colors are absent", () => {
        const dispatch = jest.fn();
        actions.setBaseColor()(dispatch, colorEmptyState);
        expect(dispatch.mock.calls.length).toEqual(0);
    });
    it("should dispatch for color list when get succeeds", () => {
        const fakeData = ["#ff0000", "#f00000"];
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors/candidates/oo?size=9`)
            .reply(200, fakeData);

        const dispatch = jest.fn();
        actions.setBaseColor()(dispatch, colorNonEmptyState).then(() => {
            expect(dispatch.mock.calls.length).toEqual(2);
            expect(dispatch.mock.calls[1][0]).toEqual({
                type: types.SET_COLOR_CODE_LIST,
                payload: fakeData,
            });
        });
    });
    it("should not dispatch for displayed color list when get fails", () => {
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors/candidates/oo?size=9`)
            .reply(400);

        const dispatch = jest.fn();
        actions.setBaseColor()(dispatch, colorNonEmptyState).then(() => {
            expect(dispatch.mock.calls.length).toEqual(1);
        });
    });
});

describe("setSelectedColorCodeList(colorList)", function () {
    it("should return colorList as payload", () => {
        const fakeData = ["#ff0000", "#f00000"];
        const action = actions.setSelectedColorCodeList(fakeData);
        expect(action.type).toEqual(types.SET_SELECTED_COLOR_CODE_LIST);
        expect(action.payload).toEqual(fakeData);
    });
});

describe("resetSelectedColorCodeList()", function () {
    it("should return empty list as payload", () => {
        const action = actions.resetSelectedColorCodeList();
        expect(action.type).toEqual(types.SET_SELECTED_COLOR_CODE_LIST);
        expect(action.payload).toEqual([]);
    });
});
