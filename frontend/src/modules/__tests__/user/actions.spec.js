import {actions, types} from "../../user";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";

describe("verifyLoginState()", function () {
    it("should dispatch when user is present", () => {
        const fakeResponse = {userID: "foo"};
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
            .reply(200, fakeResponse);

        const dispatch = jest.fn();
        actions.verifyLoginState()(dispatch).then(() => {
            expect(dispatch.mock.calls[0][0]).toEqual({
                type: types.SET_ID,
                payload: fakeResponse.userID,
            });
        });
    });
    it("should not dispatch when user is absent", () => {
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
            .reply(404);

        const dispatch = jest.fn();
        actions.verifyLoginState()(dispatch).then(() => {
            expect(dispatch.mock.calls.length).toEqual(0);
        });
    });
});

describe("login(id)", function () {
    it("should dispatch when user is present", () => {
        const fakeId = {id: "foo"};
        const mockAxios = new MockAdapter(axios);
        mockAxios.onPost(`${process.env.WEBAPI_HOST}/api/v1/login`, fakeId)
            .reply(200);

        const dispatch = jest.fn();
        actions.login(fakeId.id)(dispatch).then(() => {
            expect(dispatch.mock.calls[0][0]).toEqual({
                type: types.SET_ID,
                payload: fakeId.id,
            });
        });
    });
    it("should not dispatch when user is absent", () => {
        const fakeId = {id: "foo"};
        const mockAxios = new MockAdapter(axios);
        mockAxios.onPost(`${process.env.WEBAPI_HOST}/api/v1/login`, fakeId)
            .reply(404);

        const dispatch = jest.fn();
        actions.login()(dispatch).then(() => {
            expect(dispatch.mock.calls.length).toEqual(0);
        });
    });
});

describe("signUp(nationality, gender, birth)", function () {
    it("should dispatch when user registration succeeds", () => {
        const fakeUser = {
            nationality: "Japan",
            gender: "Male",
            birth: 1990
        };
        const fakeResponse = {
            nationality: "Japan",
            gender: "Male",
            birth: 1990,
            date: "3000-01-01T00:00:00.0000000+09:00",
            id: "foo"
        };
        const mockAxios = new MockAdapter(axios);
        mockAxios.onPost(`${process.env.WEBAPI_HOST}/api/v1/users`, fakeUser)
            .reply(200, fakeResponse);

        const dispatch = jest.fn();
        actions.signUp(fakeUser.nationality, fakeUser.gender, fakeUser.birth)(dispatch).then(() => {
            expect(dispatch.mock.calls[0][0]).toEqual({
                type: types.SET_ID,
                payload: fakeResponse.id
            });
        });
    });
    it("should not dispatch when user registration fails", () => {
        const fakeUser = {
            nationality: "Japan",
            gender: "Male",
            birth: 1990
        };
        const mockAxios = new MockAdapter(axios);
        mockAxios.onPost(`${process.env.WEBAPI_HOST}/api/v1/users`, fakeUser)
            .reply(400);

        const dispatch = jest.fn();
        actions.signUp()(dispatch).then(() => {
            expect(dispatch.mock.calls.length).toEqual(0);
        });
    });
});
