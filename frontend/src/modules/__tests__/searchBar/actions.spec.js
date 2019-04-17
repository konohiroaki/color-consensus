import {actions, types} from "../../searchBar";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";

describe("fetchColors()", function () {
    it("should dispatch when fetch is success", () => {
        const fakeResponse = [{foo: "bar"}];
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors`)
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
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/colors`)
            .reply(400);

        const dispatch = jest.fn();
        actions.fetchColors()(dispatch).then(() => {
            expect(dispatch.mock.calls.length).toEqual(0);
        });
    });
});
