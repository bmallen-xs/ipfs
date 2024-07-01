# NFT Metadata Scraper

This application scrapes NFT metadata from IPFS using a CSV list of IPFS CIDs and stores the results in a PostgreSQL database. It also hosts an API that allows users to retrieve all stored metadata or a specific row based on a CID.

## Features

- Read a list of IPFS CIDs from a CSV file
- Fetch metadata for each CID from IPFS
- Store the `name` and `image` fields in a PostgreSQL database
- Provide an API to retrieve all data or a specific row based on a CID

## Prerequisites

- Go 1.19 or higher
- PostgreSQL
- Docker (for running PostgreSQL in a container, if desired)

## Setup

### Step 1: Clone the Repository

```
git clone https://github.com/shawnwollenberg/ipfs-metadata.git
cd nft_scraper
```

### Step 2: Set Up PostgreSQL
You can either set up PostgreSQL locally or use Docker to run it in a container.

Using Docker:

```
docker run --name postgres -e POSTGRES_USER=youruser -e POSTGRES_PASSWORD=yourpassword -e POSTGRES_DB=yourdb -p 5432:5432 -d postgres
```

### Step 3: Configure Environment Variables
Create a .env file in the root directory of the project with the following content:

env
```
POSTGRES_USER=youruser
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=yourdb
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
```

### Step 4: Install Dependencies
```
go mod tidy
```

### Step 5: Prepare the CSV File
A CSV file has been saved in the `data` directory, but if you would like to add additional CIDs feel free to adjust. Each row should contain one CID.

### Step 6: Run the Application
```
go run .
```
This will:

* Read the CSV file.
* Fetch metadata for each CID from IPFS.
* Store the name and image fields in the PostgreSQL database.
* Start the API server.

### API Endpoints

#### Get All Metadata
##### Request:

```
GET /metadata
```

##### Response:
```
[
  {
    "cid": "Qm...",
    "name": "Example Name",
    "image": "Example Image URL"
  },
  ...
]
```

#### Get Metadata by CID
##### Request:

```GET /metadata/:cid```

##### Response:

```
{
  "cid": "Qm...",
  "name": "Example Name",
  "image": "Example Image URL"
}
```

## Acknowledgements
* [Gin Gonic](https://github.com/gin-gonic/gin) for the web framework.
* [sqlx](https://github.com/jmoiron/sqlx) for SQL database interactions.
* [godotenv](https://github.com/joho/godotenv) for loading environment variables from a .env file.

## Contact
For any questions or suggestions, please open an issue or contact the repository owner.