# azure_blob_example


export GOPATH=/path/to/your/desired/directory
export GOBIN=/path/to/your/desired/directory/bin

GOBIN=/usr/local/bin PATH=$PATH:/usr/local/go/bin make install


go get github.com/Azure/azure-sdk-for-go/sdk/azidentity@v1.1.0 

go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob@v0.4.1

export PATH=$PATH:$GOBIN




      rgn.sdk.create_role_assignment_storage_account_container(
        Config.archive_storage_account_rg_name, # <resource-group-name>
        rgn.wal_bucket_name, # <storage-account-name>
        eid, # <container-name>
        SecureRandom.uuid, # <role-assignment-id>
        {
          "properties": {
            "roleDefinitionId": "/providers/Microsoft.Authorization/roleDefinitions/ba92f5b4-2d11-453d-a403-e96b0029c9fe", # Storage Blob Data Contributor
            "principalId": principal_ids.first,
            "principalType": "ServicePrincipal",
          },
        },
      )
