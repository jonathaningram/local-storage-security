# Demonstration of `localStorage` security issues

When storing sensitive information in `localStorage` it is possible for any scripts running in your application to steal that information.

This demonstration shows how JWT tokens stored in local storage might be stolen.

Even if all of your scripts are served from your application's domain, you could still be at risk. If you're installing modules from a registry like NPM you could be vulnerable if you don't know what's actually in the module.

This demonstration simulates that an `analytics.js` file that you trusted is now compromised. The compromised script sends off a request to an attacker's server and bundles up your JWT token.

# Prerequisites

- Go
- Google App Engine dev tools

# Design

Each of the 3 applications are modelled as an App Engine service under the one Google Cloud project.

# Development

Start the App Engine development server which starts all services:

```
make start
```

# Deployment

Deploy all services:

```
make deploy
```
