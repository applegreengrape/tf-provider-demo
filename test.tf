provider "metadata" {
    api_tok = ""
}

data "metadata_tags" "test" {
    path = "/v1/tag"
	query_string = "abc"
}

output "test" {
    value = data.metadata_tags.test
}