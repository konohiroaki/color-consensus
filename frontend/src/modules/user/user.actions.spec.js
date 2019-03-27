import {actions, types} from "./user";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";

describe("verifyLoginState()", function () {
    it("should dispatch when user is present", () => {
        const fakeResponse = {userID: "foo"};
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
            .reply(200, fakeResponse);

        const spy = jest.fn();
        const thunk = actions.verifyLoginState();
        thunk(spy).then(() => {
            expect(spy.mock.calls[0][0]).toEqual({
                type: types.SET_ID,
                payload: fakeResponse.userID,
            });
        });
    });
    it("should not dispatch when user is absent", () => {
        const mockAxios = new MockAdapter(axios);
        mockAxios.onGet(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
            .reply(404);

        const spy = jest.fn();
        const thunk = actions.verifyLoginState();
        thunk(spy).then(() => {
            expect(spy.mock.calls.length).toEqual(0);
        });
    });
});

describe("login(id)", function () {
    it("should dispatch when user is present", () => {
        const fakeId = {id: "foo"};
        const mockAxios = new MockAdapter(axios);
        mockAxios.onPost(`${process.env.WEBAPI_HOST}/api/v1/users/presence`, fakeId)
            .reply(200);

        const spy = jest.fn();
        const thunk = actions.login(fakeId.id);
        thunk(spy).then(() => {
            expect(spy.mock.calls[0][0]).toEqual({
                type: types.SET_ID,
                payload: fakeId.id,
            });
        });
    });
    it("should not dispatch when user is absent", () => {
        const fakeId = {id: "foo"};
        const mockAxios = new MockAdapter(axios);
        mockAxios.onPost(`${process.env.WEBAPI_HOST}/api/v1/users/presence`, fakeId)
            .reply(404);

        const spy = jest.fn();
        const thunk = actions.login();
        thunk(spy).then(() => {
            expect(spy.mock.calls.length).toEqual(0);
        });
    });
});

describe("signUp(nationality, gender, birth)", function () {
    it("should dispatch when user is present", () => {
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

        const spy = jest.fn();
        const thunk = actions.signUp(fakeUser.nationality,
            fakeUser.gender, fakeUser.birth);
        thunk(spy).then(() => {
            expect(spy.mock.calls[0][0]).toEqual({
                type: types.SET_ID,
                payload: fakeResponse.id
            });
        });
    });
    it("should not dispatch when user is absent", () => {
        const fakeUser = {
            nationality: "Japan",
            gender: "Make",
            birth: 1990
        };
        const mockAxios = new MockAdapter(axios);
        mockAxios.onPost(`${process.env.WEBAPI_HOST}/api/v1/users`, fakeUser)
            .reply(400);

        const spy = jest.fn();
        const thunk = actions.signUp();
        thunk(spy).then(() => {
            expect(spy.mock.calls.length).toEqual(0);
        });
    });
});
