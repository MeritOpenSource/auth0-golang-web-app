# Merit + Auth0 + Go Web App Sample

This sample demonstrates how to add authentication to a Go web app using Auth0 and interact with the Merit APIs.

Check the [Go Quickstart](https://auth0.com/docs/quickstart/webapp/golang) to better understand this sample for
authentication with Auth0. [Contact Merit](https://www.merits.com/resources/developers/contact) to get set up with an
Agent to perform API interactions and refer to the ReadMe documentation on the
[Merit API](https://merit-developer-portal.readme.io/docs/getting-started) for more information.

This sample web-app is a standard setup for Auth0 web-apps, not for single page applications or native binaries. For 
the recommended setup for those use cases, please refer to the [Auth0 Docs on PKCE].

This sample web-app does demonstrate standard interactions with the Merit API.

## Running the App

To run the app, make sure you have **go** installed.

Rename the `.env.example` file to `.env` and provide your Auth0 credentials. You will also be told of the API url as
well as the audience for the API.

```bash
# .env

AUTH0_CLIENT_ID={CLIENT_ID}
AUTH0_DOMAIN={DOMAIN}
AUTH0_CLIENT_SECRET={CLIENT_SECRET}
AUTH0_CALLBACK_URL=http://localhost:3000/callback
MERIT_API_URL={MERIT_API_URL}
MERIT_AUDIENCE={MERIT_AUDIENCE}

```

Once you've set your Auth0 credentials in the `.env` file, run `go mod download` to download the Go dependencies.

Run `go run main.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/).

## What is Auth0?

Auth0 helps you to:

* Add authentication with [multiple authentication sources](https://auth0.com/docs/authenticate/identity-providers), either social like **Google, Facebook, Microsoft Account, LinkedIn, GitHub, Twitter, Box, Salesforce, amont others**, or enterprise identity systems like **Windows Azure AD, Google Apps, Active Directory, ADFS or any SAML Identity Provider**.
* Add authentication through more traditional **[username/password databases](https://auth0.com/docs/authenticate/database-connections/custom-db/create-db-connection)**.
* Add support for **[linking different user accounts](https://auth0.com/docs/manage-users/user-accounts/user-account-linking/link-user-accounts)** with the same user.
* Support for generating signed [Json Web Tokens](https://auth0.com/docs/secure/tokens/json-web-tokens) to call your APIs and **flow the user identity** securely.
* Analytics of how, when and where users are logging in.
* Pull data from other sources and add it to the user profile, through [JavaScript rules](https://auth0.com/docs/customize/rules).

## Create a free Auth0 account

1. Go to [Auth0](https://auth0.com/signup) and click Sign Up.
2. Use Google, GitHub or Microsoft Account to login.

## Issue Reporting

If you have found a bug or if you have a feature request, please report them at this repository issues section. Please do not report security vulnerabilities on the public GitHub issue tracker. The [Responsible Disclosure Program](https://auth0.com/whitehat) details the procedure for disclosing security issues.

## Author

[Auth0](https://auth0.com)

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE.txt) file for more info.

[Auth0 Docs on PKCE]:https://auth0.com/docs/get-started/authentication-and-authorization-flow/authorization-code-flow-with-pkce