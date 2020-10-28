# custom terraform provider

This is an example about how to build a custom terraform provider for a cutom API built with aws api gateway, lambda functions and dynamodb

```bash
bash-3.2$ terraform init 

Initializing the backend...

Initializing provider plugins...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
bash-3.2$ terraform plan
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.

data.metadata_tags.test: Refreshing state...

------------------------------------------------------------------------

No changes. Infrastructure is up-to-date.

This means that Terraform did not detect any differences between your
configuration and real physical resources that exist. As a result, no
actions need to be performed.
bash-3.2$ terraform apply
data.metadata_tags.test: Refreshing state...

Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

test = {
  "id" = "tag-1603843782"
  "path" = "/v1/tag"
  "query_string" = "abc"
  "tags" = {
    "hostname" = "📛 hostname"
    "id" = "1"
    "meta" = "📝 Blablablablabla, blabla bla blablablabla blabla blablablablablabla"
    "name" = "🗿"
    "owner" = "🐻 owner"
    "region" = "📍 home"
    "tag" = "abc"
  }
}

```

# How to build it?
```code 
$ go build -o terraform-provider-metadata .
```