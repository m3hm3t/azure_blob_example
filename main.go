package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func main() {
	accountName := "mitest3"
	containerName := "test1"

	// Construct the blob URL
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName)

	// Set up the DefaultAzureCredential with a client ID for the managed identity
	clientID := "ba2b86da-5bcb-41d4-b8cb-774affb9d15e"
	cred, err := azidentity.NewManagedIdentityCredential(&azidentity.ManagedIdentityCredentialOptions{
		ID: azidentity.ClientID(clientID),
	})
	if err != nil {
		fmt.Printf("failed to obtain a credential: %v\n", err)
		return
	}

	containerClient, err := azblob.NewContainerClient(containerURL, cred, nil)
	if err != nil {
		fmt.Printf("failed to create container client: %v\n", err)
		return
	}

	blobNameToDownload := "a.txt"
	localFileNameToDownload := "a.txt"
	// Create a new BlobClient
	blobClient, err := containerClient.NewBlobClient(blobNameToDownload)
	if err != nil {
		fmt.Printf("failed to create blob client: %v\n", err)
		return
	}

	// Create a local file to save the downloaded content
	file, err := os.Create(localFileNameToDownload)
	if err != nil {
		fmt.Printf("failed to create local file: %v\n", err)
		return
	}
	defer file.Close()

	// Download the blob to the local file
	ctx := context.Background()
	_, err = blobClient.Download(ctx, nil)
	if err != nil {
		fmt.Printf("failed to download blob: %v\n", err)
		return
	}

	fmt.Printf("Blob downloaded to '%s'\n", localFileNameToDownload)

	blobNameToUpload := "b.txt"
	localFileNameToUpload := "b.txt"

	// Create a new BlobClient
	blobClientToUpload, err := containerClient.NewBlockBlobClient(blobNameToUpload)
	if err != nil {
		fmt.Printf("failed to create blob client: %v\n", err)
		return
	}

	// Open the local file
	fileToUpload, err := os.Open(localFileNameToUpload)
	if err != nil {
		fmt.Printf("failed to open local file: %v\n", err)
		return
	}
	defer fileToUpload.Close()

	// Upload the file
	_, err = blobClientToUpload.Upload(ctx, fileToUpload, nil)
	if err != nil {
		fmt.Printf("failed to upload blob: %v\n", err)
		return
	}

	fmt.Printf("Blob uploaded to '%s'\n", blobNameToUpload)
}
