import { useAuthStore } from "@/stores/auth";
import { flushPromises } from "@vue/test-utils";
import { User, UserManager, WebStorageStateStore } from "oidc-client-ts";
import { setActivePinia, createPinia } from "pinia";
import { assert, expect, describe, it, vi, beforeEach } from "vitest";

describe("useAuthStore", () => {
  beforeEach(() => {
    setActivePinia(createPinia());
  });

  it("checks if auth. is enabled", () => {
    const authStore = useAuthStore();
    authStore.$patch((state) => (state.config.enabled = false));
    expect(authStore.isEnabled).toEqual(false);
    authStore.$patch((state) => (state.config.enabled = true));
    expect(authStore.isEnabled).toEqual(true);
  });

  it("checks if the user is valid", () => {
    const authStore = useAuthStore();
    authStore.$patch((state) => {
      state.config = {
        enabled: true,
        baseUrl: "",
        provider: "",
        clientId: "",
        extraScopes: "",
        extraQueryParams: "",
        abac: {
          enabled: false,
          claimPath: "",
          claimPathSeparator: "",
          claimValuePrefix: "",
          useRoles: false,
          rolesMapping: {},
        },
      };
    });

    // No user.
    expect(authStore.isUserValid).toEqual(false);

    // Expired user.
    var user = new User({
      access_token: "",
      token_type: "",
      profile: { aud: "", exp: 0, iat: 0, iss: "", sub: "" },
      expires_at: Date.now() / 1000 - 1,
    });
    authStore.setUser(user);
    expect(authStore.isUserValid).toEqual(false);

    // Valid user.
    user = new User({
      access_token: "",
      token_type: "",
      profile: { aud: "", exp: 0, iat: 0, iss: "", sub: "" },
      expires_at: Date.now() / 1000 + 1,
    });
    authStore.setUser(user);
    expect(authStore.isUserValid).toEqual(true);
  });

  it("gets the user display name", () => {
    const authStore = useAuthStore();
    authStore.$patch((state) => (state.config.enabled = false));

    // No user.
    expect(authStore.getUserDisplayName).toEqual(undefined);

    // User with preferred_username, name and email.
    var user = new User({
      access_token: "",
      token_type: "",
      profile: {
        aud: "",
        exp: 0,
        iat: 0,
        iss: "",
        sub: "",
        preferred_username: "preferred_username",
        name: "name",
        email: "name@example.com",
      },
    });
    authStore.setUser(user);
    expect(authStore.getUserDisplayName).toEqual("preferred_username");

    // User with name and email.
    user = new User({
      access_token: "",
      token_type: "",
      profile: {
        aud: "",
        exp: 0,
        iat: 0,
        iss: "",
        sub: "",
        name: "name",
        email: "name@example.com",
      },
    });
    authStore.setUser(user);
    expect(authStore.getUserDisplayName).toEqual("name");

    // User with email.
    user = new User({
      access_token: "",
      token_type: "",
      profile: {
        aud: "",
        exp: 0,
        iat: 0,
        iss: "",
        sub: "",
        email: "name@example.com",
      },
    });
    authStore.setUser(user);
    expect(authStore.getUserDisplayName).toEqual("name@example.com");
  });

  it("gets the user access token", () => {
    const authStore = useAuthStore();
    authStore.$patch((state) => (state.config.enabled = false));
    expect(authStore.getUserAccessToken).toEqual("");

    let user = new User({
      access_token: "access_token",
      token_type: "",
      profile: {
        aud: "",
        exp: 0,
        iat: 0,
        iss: "",
        sub: "",
      },
    });
    authStore.setUser(user);
    expect(authStore.getUserAccessToken).toEqual("access_token");
  });

  it.each([
    {
      title: "ABAC disabled",
      enabled: false,
      attributes: [],
      required: ["package:list"],
      expected: true,
    },
    {
      title: "all wildcard",
      enabled: true,
      attributes: ["*"],
      required: ["package:list", "storage:location:list"],
      expected: true,
    },
    {
      title: "exact matches",
      enabled: true,
      attributes: ["package:list", "storage:location:list"],
      required: ["package:list", "storage:location:list"],
      expected: true,
    },
    {
      title: "wildcard matches",
      enabled: true,
      attributes: ["package:*", "storage:*"],
      required: ["package:list", "storage:location:list"],
      expected: true,
    },
    {
      title: "no match",
      enabled: true,
      attributes: ["package:*"],
      required: ["package:list", "storage:location:list"],
      expected: false,
    },
    {
      title: "no match without wildcard",
      enabled: true,
      attributes: ["storage:location"],
      required: ["storage:location:list"],
      expected: false,
    },
  ])("checks attributes ($title)", (test) => {
    const authStore = useAuthStore();
    authStore.$patch((state) => {
      state.config = {
        enabled: true,
        baseUrl: "",
        provider: "",
        clientId: "",
        extraScopes: "",
        extraQueryParams: "",
        abac: {
          enabled: test.enabled,
          claimPath: "",
          claimPathSeparator: "",
          claimValuePrefix: "",
          useRoles: false,
          rolesMapping: {},
        },
      };
      state.attributes = test.attributes;
    });

    expect(authStore.checkAttributes(test.required)).toEqual(test.expected);
  });

  it("loads the configuration from the environment", () => {
    const authStore = useAuthStore();
    authStore.loadConfig();
    expect(authStore.config).toEqual({
      enabled: true,
      baseUrl: "http://localhost:8080",
      provider: "http://keycloak:7470/realms/artefactual",
      clientId: "enduro",
      extraScopes: "enduro",
      extraQueryParams: "audience=enduro-api, key = value",
      abac: {
        enabled: true,
        claimPath: "attributes.enduro",
        claimPathSeparator: ".",
        claimValuePrefix: "enduro:",
        useRoles: true,
        rolesMapping: {
          admin: ["*"],
          operator: [
            "package:list",
            "package:listActions",
            "package:move",
            "package:read",
            "package:upload",
          ],
          readonly: ["package:list", "package:listActions", "package:read"],
        },
      },
    });
  });

  it("loads the manager", () => {
    const authStore = useAuthStore();
    authStore.loadManager();
    expect(authStore.manager).toBeInstanceOf(UserManager);
    expect(authStore.manager?.settings.authority).toEqual(
      "http://keycloak:7470/realms/artefactual",
    );
    expect(authStore.manager?.settings.client_id).toEqual("enduro");
    expect(authStore.manager?.settings.redirect_uri).toEqual(
      "http://localhost:8080/user/signin-callback",
    );
    expect(authStore.manager?.settings.post_logout_redirect_uri).toEqual(
      "http://localhost:8080/user/signout-callback",
    );
    expect(authStore.manager?.settings.scope).toEqual(
      "openid email profile enduro",
    );
    expect(authStore.manager?.settings.extraQueryParams).toEqual({
      audience: "enduro-api",
      key: "value",
    });
    expect(authStore.manager?.settings.userStore).toBeInstanceOf(
      WebStorageStateStore,
    );
  });

  it("doesn't load the manager", () => {
    const authStore = useAuthStore();
    authStore.$patch((state) => (state.config.enabled = false));
    authStore.loadManager();
    expect(authStore.manager).toEqual(null);
  });

  it("redirects for signin", () => {
    const manager = new UserManager({
      authority: "",
      client_id: "",
      redirect_uri: "",
    });

    const redirectMock = vi.fn().mockImplementation(manager.signinRedirect);
    redirectMock.mockImplementation(async () => null);
    manager.signinRedirect = redirectMock;

    const authStore = useAuthStore();
    authStore.$patch((state) => (state.manager = manager));
    authStore.signinRedirect();
    expect(redirectMock).toHaveBeenCalledOnce();
  });

  it("receives a signin callback", async () => {
    const manager = new UserManager({
      authority: "",
      client_id: "",
      redirect_uri: "",
    });

    const callbackMock = vi.fn().mockImplementation(manager.signinCallback);
    callbackMock.mockImplementation(async () => undefined);
    manager.signinCallback = callbackMock;

    const authStore = useAuthStore();
    authStore.$patch((state) => {
      state.config.enabled = false;
      state.manager = manager;
    });

    authStore.signinCallback();
    await flushPromises();
    expect(authStore.user).toEqual(null);
  });

  it("redirects for signout", () => {
    const manager = new UserManager({
      authority: "",
      client_id: "",
      redirect_uri: "",
    });

    const redirectMock = vi.fn().mockImplementation(manager.signoutRedirect);
    redirectMock.mockImplementation(async () => null);
    manager.signoutRedirect = redirectMock;

    const authStore = useAuthStore();
    authStore.$patch((state) => (state.manager = manager));
    authStore.signoutRedirect();
    expect(redirectMock).toHaveBeenCalledOnce();
  });

  it("receives a signout callback", async () => {
    const manager = new UserManager({
      authority: "",
      client_id: "",
      redirect_uri: "",
    });

    const callbackMock = vi.fn().mockImplementation(manager.signoutCallback);
    callbackMock.mockImplementation(async () => null);
    manager.signoutCallback = callbackMock;

    const authStore = useAuthStore();
    authStore.$patch((state) => (state.manager = manager));

    const removeUserMock = vi.fn().mockImplementation(authStore.removeUser);
    removeUserMock.mockImplementation(async () => null);
    authStore.removeUser = removeUserMock;

    authStore.signoutCallback();
    await flushPromises();
    expect(callbackMock).toHaveBeenCalledOnce();
    expect(removeUserMock).toHaveBeenCalledOnce();
  });

  it("loads and removes the user and attributes", async () => {
    var user = new User({
      access_token: "",
      token_type: "",
      profile: { aud: "", exp: 0, iat: 0, iss: "", sub: "" },
    });

    const manager = new UserManager({
      authority: "",
      client_id: "",
      redirect_uri: "",
    });

    const getUserMock = vi.fn().mockImplementation(manager.getUser);
    getUserMock.mockImplementation(async () => user);
    manager.getUser = getUserMock;

    const removeUserMock = vi.fn().mockImplementation(manager.removeUser);
    removeUserMock.mockImplementation(async () => null);
    manager.removeUser = removeUserMock;

    const authStore = useAuthStore();
    authStore.$patch((state) => (state.manager = manager));

    const parseAttrMock = vi.fn().mockImplementation(authStore.parseAttributes);
    parseAttrMock.mockImplementation(() => {
      authStore.$patch((state) => (state.attributes = ["*"]));
    });
    authStore.parseAttributes = parseAttrMock;

    expect(authStore.user).toEqual(null);
    expect(authStore.attributes).toEqual([]);
    authStore.loadUser();
    await flushPromises();
    expect(authStore.user).toEqual(user);
    expect(authStore.attributes).toEqual(["*"]);
    authStore.removeUser();
    await flushPromises();
    expect(authStore.user).toEqual(null);
    expect(authStore.attributes).toEqual([]);
  });

  it("doesn't load an undefined user", async () => {
    const manager = new UserManager({
      authority: "",
      client_id: "",
      redirect_uri: "",
    });

    const getUserMock = vi.fn().mockImplementation(manager.getUser);
    getUserMock.mockImplementation(async () => undefined);
    manager.getUser = getUserMock;

    const authStore = useAuthStore();
    authStore.$patch((state) => (state.manager = manager));

    authStore.loadUser();
    await flushPromises();
    expect(authStore.user).toEqual(null);
  });

  it("parses the attributes when the user is set", () => {
    const authStore = useAuthStore();
    const spy = vi.spyOn(authStore, "parseAttributes");
    authStore.setUser(null);
    expect(spy).toHaveBeenCalledOnce();
  });

  it.each([
    {
      title: "ABAC disabled",
      enabled: false,
      claimPath: "",
      claimPathSeparator: "",
      claimValuePrefix: "",
      accessToken: "",
      expected: [],
    },
    {
      title: "top-level claim",
      enabled: true,
      claimPath: "enduro",
      claimPathSeparator: "",
      claimValuePrefix: "",
      accessToken:
        /*
        {
          "exp": 1485140984,
          "iat": 1485137384,
          "iss": "acme.com",
          "sub": "29ac0c18-0b4a-42cf-82fc-03d570318a1d",
          "enduro": [
            "package:*",
            "storage:*"
          ]
        }
        */
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODUxNDA5ODQsImlhdCI6MTQ4NTEzNzM4NCwiaXNzIjoiYWNtZS5jb20iLCJzdWIiOiIyOWFjMGMxOC0wYjRhLTQyY2YtODJmYy0wM2Q1NzAzMThhMWQiLCJlbmR1cm8iOlsicGFja2FnZToqIiwic3RvcmFnZToqIl19.Mp0Pcwsz5VECK11Kf2ZZNF_SMKu5CgBeLN9ZOP04kZo",
      expected: ["package:*", "storage:*"],
    },
    {
      title: "nested claim",
      enabled: true,
      claimPath: "attributes.enduro",
      claimPathSeparator: ".",
      claimValuePrefix: "",
      accessToken:
        /*
        {
          "exp": 1485140984,
          "iat": 1485137384,
          "iss": "acme.com",
          "sub": "29ac0c18-0b4a-42cf-82fc-03d570318a1d",
          "attributes": {
            "enduro": [
              "package:*",
              "storage:*"
            ]
          }
        }
        */
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODUxNDA5ODQsImlhdCI6MTQ4NTEzNzM4NCwiaXNzIjoiYWNtZS5jb20iLCJzdWIiOiIyOWFjMGMxOC0wYjRhLTQyY2YtODJmYy0wM2Q1NzAzMThhMWQiLCJhdHRyaWJ1dGVzIjp7ImVuZHVybyI6WyJwYWNrYWdlOioiLCJzdG9yYWdlOioiXX19.Mp0Pcwsz5VECK11Kf2ZZNF_SMKu5CgBeLN9ZOP04kZo",
      expected: ["package:*", "storage:*"],
    },
    {
      title: "filters values",
      enabled: true,
      claimPath: "attributes.enduro",
      claimPathSeparator: ".",
      claimValuePrefix: "enduro:",
      accessToken:
        /*
        {
          "exp": 1485140984,
          "iat": 1485137384,
          "iss": "acme.com",
          "sub": "29ac0c18-0b4a-42cf-82fc-03d570318a1d",
          "attributes": {
            "enduro": [
              "enduro:package:*",
              "enduro:storage:*",
              "ignore:this"
            ]
          }
        }
        */
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODUxNDA5ODQsImlhdCI6MTQ4NTEzNzM4NCwiaXNzIjoiYWNtZS5jb20iLCJzdWIiOiIyOWFjMGMxOC0wYjRhLTQyY2YtODJmYy0wM2Q1NzAzMThhMWQiLCJhdHRyaWJ1dGVzIjp7ImVuZHVybyI6WyJlbmR1cm86cGFja2FnZToqIiwiZW5kdXJvOnN0b3JhZ2U6KiIsImlnbm9yZTp0aGlzIl19fQ.Mp0Pcwsz5VECK11Kf2ZZNF_SMKu5CgBeLN9ZOP04kZo",
      expected: ["package:*", "storage:*"],
    },
    {
      title: "invalid token",
      enabled: true,
      claimPath: "enduro",
      claimPathSeparator: "",
      claimValuePrefix: "",
      accessToken: "invalid token",
      error:
        "Error decoding or parsing token: TypeError: The first argument must be one of type string, Buffer, ArrayBuffer, Array, or Array-like Object. Received type undefined",
    },
    {
      title: "not found top-level",
      enabled: true,
      claimPath: "enduro",
      claimPathSeparator: "",
      claimValuePrefix: "",
      accessToken:
        /*
        {
          "exp": 1485140984,
          "iat": 1485137384,
          "iss": "acme.com",
          "sub": "29ac0c18-0b4a-42cf-82fc-03d570318a1d",
          "enduro": [
            "package:*",
            "storage:*"
          ]
        }
        */
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODUxNDA5ODQsImlhdCI6MTQ4NTEzNzM4NCwiaXNzIjoiYWNtZS5jb20iLCJzdWIiOiIyOWFjMGMxOC0wYjRhLTQyY2YtODJmYy0wM2Q1NzAzMThhMWQiLCJvdGhlciI6WyJwYWNrYWdlOioiLCJzdG9yYWdlOioiXX0.Mp0Pcwsz5VECK11Kf2ZZNF_SMKu5CgBeLN9ZOP04kZo",
      error: "Attributes not found in token, claim path: enduro",
    },
    {
      title: "not found nested",
      enabled: true,
      claimPath: "enduro.attributes",
      claimPathSeparator: ".",
      claimValuePrefix: "",
      accessToken:
        /*
        {
          "exp": 1485140984,
          "iat": 1485137384,
          "iss": "acme.com",
          "sub": "29ac0c18-0b4a-42cf-82fc-03d570318a1d",
          "enduro": null
        }
        */
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODUxNDA5ODQsImlhdCI6MTQ4NTEzNzM4NCwiaXNzIjoiYWNtZS5jb20iLCJzdWIiOiIyOWFjMGMxOC0wYjRhLTQyY2YtODJmYy0wM2Q1NzAzMThhMWQiLCJlbmR1cm8iOm51bGx9.Mp0Pcwsz5VECK11Kf2ZZNF_SMKu5CgBeLN9ZOP04kZo",
      error: "Attributes not found in token, claim path: enduro.attributes",
    },
    {
      title: "not multi-value",
      enabled: true,
      claimPath: "enduro",
      claimPathSeparator: "",
      claimValuePrefix: "",
      accessToken:
        /*
        {
          "exp": 1485140984,
          "iat": 1485137384,
          "iss": "acme.com",
          "sub": "29ac0c18-0b4a-42cf-82fc-03d570318a1d",
          "enduro": "package:*"
        }
        */
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODUxNDA5ODQsImlhdCI6MTQ4NTEzNzM4NCwiaXNzIjoiYWNtZS5jb20iLCJzdWIiOiIyOWFjMGMxOC0wYjRhLTQyY2YtODJmYy0wM2Q1NzAzMThhMWQiLCJlbmR1cm8iOiJwYWNrYWdlOioifQ.Mp0Pcwsz5VECK11Kf2ZZNF_SMKu5CgBeLN9ZOP04kZo",
      error:
        "Attributes are not part of a multivalue claim, claim path: enduro",
    },
    {
      title: "uses roles mapping",
      enabled: true,
      claimPath: "roles",
      claimPathSeparator: "",
      claimValuePrefix: "",
      useRoles: true,
      rolesMapping: {
        admin: ["*"],
        operator: [
          "package:list",
          "package:listActions",
          "package:move",
          "package:read",
          "package:upload",
        ],
        readonly: ["package:list", "package:listActions", "package:read"],
      },
      accessToken:
        /*
        {
          "exp": 1485140984,
          "iat": 1485137384,
          "iss": "acme.com",
          "sub": "29ac0c18-0b4a-42cf-82fc-03d570318a1d",
          "roles": [
            "admin",
            "operator",
            "readonly"
          ]
        }
        */
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODUxNDA5ODQsImlhdCI6MTQ4NTEzNzM4NCwiaXNzIjoiYWNtZS5jb20iLCJzdWIiOiIyOWFjMGMxOC0wYjRhLTQyY2YtODJmYy0wM2Q1NzAzMThhMWQiLCJyb2xlcyI6WyJhZG1pbiIsIm9wZXJhdG9yIiwicmVhZG9ubHkiXX0.2PY8zO7vNcS-3RdLa0AIFLjmRFKrR55m3rlm3DI1cMM",
      expected: [
        "*",
        "package:list",
        "package:listActions",
        "package:move",
        "package:read",
        "package:upload",
      ],
    },
  ])("parses attributes ($title)", (test) => {
    const authStore = useAuthStore();
    authStore.$patch((state) => {
      state.config = {
        enabled: true,
        baseUrl: "",
        provider: "",
        clientId: "",
        extraScopes: "",
        extraQueryParams: "",
        abac: {
          enabled: test.enabled,
          claimPath: test.claimPath,
          claimPathSeparator: test.claimPathSeparator,
          claimValuePrefix: test.claimValuePrefix,
          useRoles: test.useRoles ? test.useRoles : false,
          rolesMapping: test.rolesMapping ? test.rolesMapping : {},
        },
      };
      state.user = new User({
        access_token: test.accessToken,
        token_type: "",
        profile: {
          aud: "",
          exp: 0,
          iat: 0,
          iss: "",
          sub: "",
        },
      });
      state.attributes = [];
    });

    if (test.error) {
      assert.throws(authStore.parseAttributes, test.error);
    } else {
      authStore.parseAttributes();
      expect(authStore.attributes).toEqual(test.expected);
    }
  });
});
