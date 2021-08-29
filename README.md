# udemy-api-data

This repository contains code for fetching Courses and Reviews data from the Udemy API.

Refer to [their documentation](https://www.udemy.com/developers/affiliate/) for more details.

## Usage

Put your Udemy's client ID and secret into a .env file in the root of the repository with similar content to the following:

```yaml
CLIENT_ID=<Your Client ID>
CLIENT_SECRET=<Your Client Secret>
```

Then simply execute the following command to run the data fetching:

```bash
make run
```
