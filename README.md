# tf-provider-demo

```bash
$ curl -s -H "Authorization: Bearer your_api_token" https://30cqdcgkq7.execute-api.eu-west-1.amazonaws.com/v1/tag?team=team_abc | jq .
{
  "msg": "👍Yep! it is the right api key",
  "tags": [
    {
      "hostname": "📛some_cool_hostname",
      "meta": "Blablablablabla, blabla bla blablablabla blabla blablablablablabla",
      "owner": "🐻some_cool_owner",
      "tag": "team_abc",
      "region": "📍some_cool_region",
      "id": "1",
      "name": "some_cool_name"
    }
  ]
}
```