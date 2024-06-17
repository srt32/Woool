# Woool

## Server Setup and Usage Instructions

This server accepts file uploads via HTTP and processes them to extract purchase details and amounts from receipts using OpenAI. 

### Setup

1. Clone the repository.
2. Ensure you have Go installed on your system.
3. Set up an OpenAI API key and add it to your environment variables as `OPENAI_API_KEY`.
4. Run `go build` to build the server.
5. Start the server with `./Woool`.

### Uploading a Receipt

To upload a receipt, send a POST request to the `/upload` endpoint with the receipt file as `multipart/form-data`. The file field should be named `receipt`.

Example using `curl`:

```bash
curl -F "receipt=@path_to_your_receipt_file" http://localhost:8080/upload
```

### Response Format

The server will return a JSON response containing the summary of the purchase and the dollar amount. Here is an example of the expected JSON response format:

```json
{
  "summary": "Grocery shopping at XYZ Store",
  "amount": "$150.00"
}
```

Please note that the actual summary and amount will vary based on the content of the uploaded receipt.
